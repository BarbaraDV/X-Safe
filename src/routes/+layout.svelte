<script>
  import "$lib/styles";
  import { createUserStore, globals } from "$lib";
  import { onMount } from "svelte";
  import { browser } from "$app/environment";
  import { page } from "$app/stores";
  import { title } from "$lib/globals";

  const darkMode = globals.darkMode;

  createUserStore();
  onMount(() => {
    if (!browser) return;
    const matchDark = window.matchMedia("(prefers-color-scheme: dark)");
    $darkMode = matchDark.matches;
    matchDark.addEventListener("change", (e) => {
      $darkMode = e.matches;
    });
  });
</script>

<svelte:head>
  <title>
    {$page.data.title || title}
  </title>
</svelte:head>

<div class="min-h-screen flex flex-col flex-grow">
  <slot />
</div>

<div
  id="modals"
  class="fixed w-full h-screen z-999 top-0 left-0 pointer-events-none"
/>

<style>
  :global(body) {
    --uno: "dark:bg-dark-900 font-sans content dark:text-gray-100";
  }
</style>
