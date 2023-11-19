<script lang="ts">
  import { pb } from "$lib/pb";
  import type { RecordModel } from "pocketbase";
  import { onMount } from "svelte";
  import Viewport from "./Viewport.svelte";
  import { fade, fly } from "svelte/transition";
  import { expoOut } from "svelte/easing";
  import { portal } from "./Portal.svelte";

  export let record: RecordModel;
  export let images: string[];
  let image: string | null = null;

  let token: string;

  onMount(async () => {
    token = await pb.files.getToken();
  });
</script>

{#if !images.length}
  <div class="w-full h-full flex flex-grow flex-centered space-y-4 flex-col">
    <un-i-arcticons-nothing-camera data-un-text-8xl />
    <p class="text-2xl font-bold">Sin imágenes</p>
  </div>
{:else}
  <div class="grid w-full grid-cols-2 md:grid-cols-4 lg:grid-cols-6 gap-4">
    {#each images as img (img)}
      <Viewport let:intersecting intersecting={false}>
        <!-- svelte-ignore a11y-click-events-have-key-events -->
        <!-- svelte-ignore a11y-no-static-element-interactions -->
        <div
          class="w-full flex rounded-lg h-full bg-dark-50 aspect-square overflow-hidden cursor-pointer"
          on:click={() => (image = img)}
        >
          {#if intersecting}
            <img
              src={pb.getFileUrl(record, img, {
                token,
                thumb: "200x200",
              })}
              transition:fade={{ duration: 200 }}
              alt=""
              class="flex w-full h-full object-cover transform hover:scale-105 duration-600"
            />
          {/if}
        </div>
      </Viewport>
    {/each}
  </div>
{/if}

{#if image}
  <div
    class="w-full h-screen pointer-events-auto flex flex-col flex-centered relative"
    use:portal={"#modals"}
    hidden
  >
    <div
      class="absolute w-full h-full bg-dark-900 bg-opacity-90"
      transition:fade={{ duration: 600, easing: expoOut }}
      role="none"
      on:click={() => {
        image = null;
      }}
    />
    <div
      class="max-w-95% max-h-95% rounded-xl bg-white dark:bg-dark-700 flex flex-col z-20 shadow-xl bg-gray-500 overflow-hidden relative"
      transition:fly={{ y: "1rem", duration: 400, easing: expoOut }}
      style="will-change: transform, opacity"
    >
      <div
        class="absolute bottom-0 flex items-end h-40 bg-gradient-to-b from-transparent to-dark-900 z-20 w-full transform justify-between p4"
      >
        <div class="flex items-center space-x-4">
          <button
            class="bg-green-500 p2 rounded transition-200 hover:bg-green-400 text-white"
          >
            <a
              href={pb.getFileUrl(record, image, {
                token,
                download: true,
              })}
              class="flex items-center space-x-2"
            >
              <un-i-carbon-download />
              <span class="text-xs">Descargar foto</span>
            </a>
          </button>
        </div>
      </div>
      <img
        src={pb.getFileUrl(record, image, {
          token,
        })}
        class="h-full flex"
        alt=""
      />
    </div>
  </div>
{/if}
