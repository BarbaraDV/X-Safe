<script lang="ts">
  import { browser } from "$app/environment";
  import Viewport from "$lib/components/Viewport.svelte";
  import { onMount } from "svelte";

  let headerSize = 100;

  onMount(adjustHeaderSize);

  function adjustHeaderSize() {
    if (!browser) return;
    const el = document.getElementById("header")!;
    if (!el) return;
    const list = [
      "margin-top",
      "margin-bottom",
      "border-top",
      "border-bottom",
      "padding-top",
      "padding-bottom",
      "height",
    ];
    const style = window.getComputedStyle(el);
    const height = list
      .map((k) => parseInt(style.getPropertyValue(k), 10))
      .reduce((prev, cur) => prev + cur);
    headerSize = height;
  }
</script>

<svelte:window on:resize={adjustHeaderSize} />

<Viewport
  class="flex w-full w-full h-[calc(100vh-var(--x-header))] flex-centered p4 flex-col space-y-8 relative overflow-hidden"
  animate
  once
  --a-y="0.5rem"
  --a-d="200ms"
  --x-header="{headerSize}px"
  intersecting={false}
>
  <div class="absolute flex w-full h-full anim" style:--anim-d="600ms">
    <un-i-ri-twitter-x-fill class="flex w-full h-full opacity-10" />
  </div>
  <p
    class="font-bold text-centered text-3xl lg:text-6xl flex flex-col items-center relative"
  >
    <span class="anim"> Aquí en X-Safe</span>
    <span class="anim" style:--anim-d="400ms"> la seguridad digital es</span>
    <span class="anim" style:--anim-d="600ms">nuestra pasión</span>
  </p>
  <a
    href="/app"
    class="bg-blue-500 p4 rounded text-sm shadow-lg transition-200 hover:bg-blue-400 text-white anim"
    style:--anim-d="1000ms">Ir a la app</a
  >
</Viewport>
