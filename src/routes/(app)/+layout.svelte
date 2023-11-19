<script>
  import Logout from "$lib/components/Logout.svelte";
  import { useUser } from "$lib";
  import { uploadModal } from "./stores";
  import UploadModal from "./UploadModal.svelte";
  import Logo from "$lib/components/Logo.svelte";
  import SidebarMenu from "$lib/components/SidebarMenu.svelte";
  import { onMount } from "svelte";
  import { browser } from "$app/environment";
  import Hamburger from "svelte-hamburger";
  import { fade } from "svelte/transition";
  import { afterNavigate } from "$app/navigation";
  import { page } from "$app/stores";
  import { expoOut } from "svelte/easing";
  const user = useUser();

  let open = false;
  afterNavigate(() => {
    open = false;
  });

  onMount(() => {
    if (!browser) return;
    let katexScript = document.createElement("script");
    katexScript.src =
      "https://cdnjs.cloudflare.com/ajax/libs/KaTeX/0.9.0-alpha1/katex.min.js";
    katexScript.integrity =
      "sha384-GR8SEkOO1rBN/jnOcQDFcFmwXAevSLx7/Io9Ps1rkxWp983ZIuUGfxivlF/5f5eJ";
    katexScript.crossOrigin = "anonymous";
    document.head.appendChild(katexScript);
  });
</script>

<svelte:head>
  <link
    rel="stylesheet"
    href="https://cdnjs.cloudflare.com/ajax/libs/KaTeX/0.9.0-alpha1/katex.min.css"
    integrity="sha384-8QOKbPtTFvh/lMY0qPVbXj9hDh+v8US0pD//FcoYFst2lCIf0BmT58+Heqj0IGyx"
    crossorigin="anonymous"
  />
</svelte:head>

{#if $uploadModal}
  <UploadModal bind:open={$uploadModal} />
{/if}

<div
  id="header"
  class="w-full flex justify-between p4 items-center border-b dark:border-dark-50 text-md relative"
>
  {#if open}
    <div
      class="div fixed w-full flex flex-col h-full bg-white dark:bg-dark-900 z-50 top-0 left-0 overflow-y-auto lg:hidden"
      transition:fade={{ duration: 200 }}
    >
      <div
        class="flex w-full h-200px bg-gradient-to-b fixed z-20 from-white to-transparent dark:from-dark-900"
      />
      <div class="pt28 flex flex-col p4 relative">
        <p
          class="p4 py1 rounded-full self-start bg-gray-200 dark:bg-dark-200 mb6"
        >
          {$user?.email}
        </p>
        <div class="flex items-center space-x-4 w-full mb2">
          <!-- <button -->
          <!--   on:click={() => { -->
          <!--     $uploadModal = true; -->
          <!--   }} -->
          <!--   class="bg-blue-500 p2 rounded transition-200 hover:bg-blue-400 text-white flex items-center space-x-2" -->
          <!-- > -->
          <!--   <un-i-carbon-upload /> -->
          <!--   <span class="text-xs">Subir</span> -->
          <!-- </button> -->
          <div class="flex items-center lg:space-x-6">
            <Logout>
              <button class="hover:underline">Cerrar sesión</button>
            </Logout>
          </div>
        </div>
        <SidebarMenu />
      </div>
    </div>
  {/if}
  <div class="z-51">
    <Logo />
  </div>
  <div class="z-52 pr4 lg:hidden">
    <Hamburger {open} on:click={() => (open = !open)} duoLine={false} />
  </div>
  <div class="flex items-center space-x-4 <lg:hidden">
    <!-- <button -->
    <!--   on:click={() => { -->
    <!--     $uploadModal = true; -->
    <!--   }} -->
    <!--   class="bg-blue-500 p2 rounded transition-200 hover:bg-blue-400 text-white flex items-center space-x-2" -->
    <!-- > -->
    <!--   <un-i-carbon-upload /> -->
    <!--   <span class="text-xs">Subir</span> -->
    <!-- </button> -->
    <div class="flex items-center lg:space-x-6">
      <p class="px4 py1 rounded-full bg-gray-200 dark:bg-dark-200 <lg:hidden">
        {$user?.email}
      </p>
      <Logout>
        <button class="hover:underline">Cerrar sesión</button>
      </Logout>
    </div>
  </div>
</div>

<div class="flex flex-grow h-0 w-full">
  <div
    class="w-15% border-r dark:border-dark-50 h-full flex flex-grow flex-col <lg:hidden overflow-y-auto"
  >
    <SidebarMenu />
  </div>
  {#key $page.url.pathname}
    <div
      class="p4 w-full overflow-y-auto"
      in:fade={{ duration: 600, easing: expoOut }}
    >
      <slot />
    </div>
  {/key}
</div>
