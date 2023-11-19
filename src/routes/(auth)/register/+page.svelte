<script lang="ts">
  import { applyAction, enhance } from "$app/forms";
  import { pb } from "$lib";
  import { expoOut } from "svelte/easing";
  import { fly, slide } from "svelte/transition";
  export let data;
  export let form;
</script>

<form
  method="POST"
  class="p-4 flex flex-col space-y-8 w-full p-2 items-center"
  in:fly={{ x: -10, duration: 400, easing: expoOut }}
  use:enhance={() => {
    return async ({ result }) => {
      pb.authStore.loadFromCookie(document.cookie);
      applyAction(result);
    };
  }}
>
  <div class="flex-col space-y-6 flex flex-centered">
    <div data-un-flex data-un-p-4 data-un-bg="dark-900" data-un-rounded-full>
      <un-i-carbon-user data-un-text-2xl />
    </div>
    <h1 class="text-xl text-center flex-inline space-x-2 flex-centered">
      <span>Registro de cuenta{data.admin ? " (administrador)" : ""}</span>
    </h1>
    {#if form?.error}
      <div
        class="text-sm text-red-400 text-center"
        in:slide={{ duration: 200 }}
      >
        {form.error}
      </div>
    {/if}
  </div>
  <div class="flex flex-col space-y-2 w-full">
    <label class="flex flex-col space-y-2">
      <span class="text-xs font-bold">Nombre</span>
      <input
        required
        type="text"
        name="name"
        placeholder="Ej. Pepe"
        class="bg-light-700 p-2 rounded text-gray-700 w-full flex"
      />
    </label>
    <label class="flex flex-col space-y-2">
      <span class="text-xs font-bold">Correo</span>
      <input
        required
        type="email"
        name="email"
        placeholder="Ej. pepe@gmail.com"
        class="bg-light-700 p-2 rounded text-gray-700 w-full flex"
      />
    </label>
    <label class="flex flex-col space-y-2">
      <span class="text-xs font-bold">Contraseña</span>
      <input
        required
        type="password"
        name="password"
        minlength="8"
        class="bg-light-700 p-2 rounded text-gray-700 w-full flex"
      />
    </label>
  </div>
  <button
    class="btn btn-white w-full flex-inline space-x-2 flex-centered group"
  >
    <span>Crear cuenta</span>
    <span
      data-i-carbon-arrow-right
      class="transform group-hover-translate-x-0.2rem"
    />
  </button>
</form>
