import { writable, Writable } from "svelte/store";

function persistentStore<T>(name: string, initial: T): Writable<T> {
  const previous = localStorage.getItem(name);
  if (previous !== null) console.debug(`Persisten store with key ${name} found on system. Hydrating.`);
  const store = writable<T>(previous ? JSON.parse(previous) : initial);
  store.subscribe(val => localStorage.setItem(name, JSON.stringify(val)));
  return store;
}

export type UserStore = {
  id: number,
  session_token: string,
}

export const userStore = persistentStore<UserStore>("userStore", {
  id: -1,
  session_token: "",
})

