import { browser } from "$app/environment";
import { PUBLIC_POCKETBASE_URL } from "$env/static/public";
import PocketBase from "pocketbase";
import { getContext, setContext } from "svelte";
import { writable } from "svelte/store";

export const pb = new PocketBase(PUBLIC_POCKETBASE_URL, undefined, "es-ES");

export function createUserStore() {
  const user = writable(pb.authStore.model);
  if (browser) {
    pb.authStore.loadFromCookie(document.cookie);
    pb.authStore.onChange(() => {
      user.set(pb.authStore.model);
      document.cookie = pb.authStore.exportToCookie({ httpOnly: false });
    });
  }
  setContext("currentUser", user);
  return user;
}

export function useUser(): ReturnType<typeof createUserStore> {
  return getContext("currentUser");
}
