<script lang="ts">
 import { onMount, onDestroy } from "svelte";
 import { renderBanner, unmountReact } from "$lib/renderReact/renderReact.jsx";
 let bannerContainer: any;

 onMount(() => {
  if (bannerContainer) {
   renderBanner(bannerContainer);
  }
 });
 onDestroy(() => {
  if (bannerContainer) {
   unmountReact(bannerContainer);
  }
 });
</script>

<div bind:this={bannerContainer} />
