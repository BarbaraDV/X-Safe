<script lang="ts">
  export let theme: "auto" | "light" | "dark" = "auto";
  export let href = "/";
</script>

<a
  {href}
  class="lg:text-4xl font-bold transform hover:scale-95 duration-200 will-change-transform flex items-center space-x-2"
>
  {#if theme == "auto"}
    <picture>
      <source srcset="/logo-dark.png" media="(prefers-color-scheme: dark)" />
      <source
        srcset="/logo.png"
        media="(prefers-color-scheme: light), (prefers-color-scheme: no-preference)"
      />
      <img decoding="async" src="/logo.png" alt="X Safe" class="h-14" />
    </picture>
  {:else}
    <img
      decoding="async"
      src="/logo{theme == 'dark' ? '-dark' : ''}.png"
      alt="X Safe"
      class="h-14"
    />
  {/if}
</a>
