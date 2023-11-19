<script>
  import { fade, fly } from "svelte/transition";
  import { portal } from "./Portal.svelte";
  import { expoOut } from "svelte/easing";
  import { createEventDispatcher } from "svelte";

  export let open = false;
  const distpatcher = createEventDispatcher();
  function close() {
    open = false;
    distpatcher("close");
  }
</script>

<div
  class="w-full h-screen pointer-events-auto flex flex-col flex-centered relative"
  use:portal={"#modals"}
  hidden
>
  <div
    class="absolute w-full h-full bg-dark-900 bg-opacity-90"
    transition:fade={{ duration: 600, easing: expoOut }}
    role="none"
    on:click={close}
  />
  <div
    class="<lg:w-90% max-w-95% min-w-40% p4 rounded-lg bg-white dark:bg-dark-700 flex flex-col z-20 shadow-lg space-y-4"
    transition:fly={{ y: "1rem", duration: 400, easing: expoOut }}
    style="will-change: transform, opacity"
  >
    <slot />
  </div>
</div>
