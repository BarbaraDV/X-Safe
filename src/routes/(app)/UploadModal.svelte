<script lang="ts">
  import { portal } from "$lib/components/Portal.svelte";
  import { fade, fly } from "svelte/transition";
  import { expoOut } from "svelte/easing";
  import AlbumDropdown from "$lib/components/CategoryDropdown.svelte";
  import { pb, useUser } from "$lib";
  import type { RecordModel } from "pocketbase";
  import { page } from "$app/stores";
  import { onMount } from "svelte";
  import { goto } from "$app/navigation";
  import UploadImages from "$lib/components/UploadImages.svelte";

  const user = useUser();

  export let open = false;
  let album: Record<"id" | "name", string> | null;
  let loading = false;

  onMount(() => {
    album = $page.data.album
      ? {
          id: $page.data.album.id,
          name: $page.data.album.name,
        }
      : null;
  });

  function close() {
    if (loading) return;
    open = false;
  }

  let files: File[] = [];

  function clean() {
    files = [];
  }

  $: if (!open) {
    clean();
  }

  async function uploadImages() {
    loading = true;
    let createdAlbum = album;
    if (album != null && !album?.id.trim()) {
      createdAlbum = await pb.collection("albums").create({
        name: album.name,
        author: $user!.id,
      });
    }
    try {
      let promises: Promise<RecordModel>[] = [];
      for (const [idx, file] of files.entries()) {
        const formData = new FormData();
        formData.append("file", file);
        formData.append("author", $user!.id);
        if (createdAlbum != null) {
          formData.append("album", createdAlbum.id);
        }
        promises = [
          ...promises,
          pb.collection("images").create(formData, {
            requestKey: `upload-img-${idx}`,
          }),
        ];
      }
      await Promise.allSettled(promises);
      await new Promise((resolve) => {
        setTimeout(() => {
          resolve(null);
        }, 400);
      });
      if (createdAlbum != null) {
        await goto(`/album/${createdAlbum.id}`);
      }
      open = false;
    } catch (err) {
      console.log(err);
      if (createdAlbum != null && !album?.id.trim()) {
        await pb.collection("albums").delete(createdAlbum?.id);
      }
      loading = false;
    }
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
    <div class="flex w-full justify-between items-center">
      <div class="flex items-center space-x-4">
        <un-i-carbon-cloud data-un-text-2xl />
        <div class="text-xl font-bold">Subir imagen</div>
      </div>
      {#if !loading}
        <button class="flex items-center" on:click={close}>
          <un-i-carbon-close data-un-text-2xl />
        </button>
      {/if}
    </div>
    {#if loading}
      <div class="w-full flex flex-centered space-y-4 flex-col py16">
        <un-i-svg-spinners-180-ring-with-bg class="w-12 h-12" />
        <p class="text-xl font-bold">Subiendo imágenes...</p>
      </div>
    {:else}
      <!-- svelte-ignore a11y-label-has-associated-control -->
      <label class="flex flex-col space-y-2">
        <span class="text-xs font-bold">Álbum (opcional)</span>
        <AlbumDropdown bind:selected={album} />
      </label>
      <UploadImages bind:files />
      {#if files.length}
        <button
          class="bg-green-500 p2 rounded transition-200 hover:bg-green-400 text-white flex items-center space-x-2 self-end"
          on:click={uploadImages}
        >
          <un-i-carbon-upload data-un-text-lg />
          <span class="">Subir</span>
        </button>
      {/if}
    {/if}
  </div>
</div>
