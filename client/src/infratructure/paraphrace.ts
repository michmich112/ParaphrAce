export type User = {
  id: number,
  session_token: string,
}

// could do with a bit more/better error handling
export async function getUser(): Promise<User> {
  const url = `${process.env.PARAPHRACE_API_URL}api/user/create`
  const res = await fetch(url, {
    method: 'POST',
    mode: 'cors',
  })

  return await res.json() as User;
}

// could do with better error handling but it will do for now
export async function getParaphrase(user: User, original: string): Promise<string> {
  const url = `${process.env.PARAPHRACE_API_URL}api/paraphrase/create`;
  console.log("session_token", user.session_token);
  const res = await fetch(url, {
    method: 'POST',
    mode: 'cors',
    headers: {
      'Content-Type': 'application/json',
    },
    body: JSON.stringify({
      session_token: user.session_token,
      original_text: original,
    })
  });
  const body: any = await res.json();
  if (body.hasOwnProperty('result')) {
    return body.result;
  }
  throw new Error("Invalid return type from API.");
}

