export type User = {
  id: number,
  token: string,
}

// could do with a bit more/better error handling
export async function getUser(): Promise<User> {
  const url = `${process.env.PARAPHRACE_API_URL}/api/user/create`
  const res = await fetch(url)
  return await res.json() as User;
}

