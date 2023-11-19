// See https://kit.svelte.dev/docs/types#app
// for information about these interfaces
declare global {
  type NullablePartial<T> = {
    [P in keyof T]: T[P] | null;
  };
  namespace App {
    // interface Error {}
    interface Locals {
      pb: import("pocketbase").default;
      user: import("pocketbase").Record | import("pocketbase").Admin | null;
    }
    // interface PageData {}
    // interface Platform {}
  }
}

declare namespace svelte.JSX {
  interface DOMAttributes<T> {
    onclick_outside?: CompositionEventHandler<T>;
  }
}

export {};
