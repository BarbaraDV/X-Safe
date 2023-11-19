<script lang="ts">
  import { browser } from "$app/environment";
  import { clickOutside } from "$lib/actions";
  import { pb } from "$lib/pb";
  import type { RecordModel } from "pocketbase";
  import { createEventDispatcher, tick } from "svelte";
  import { expoOut } from "svelte/easing";
  import { fly } from "svelte/transition";

  export let input: string = "";
  export let placeholder = "Crea o busca una categoría";
  export let selected: Partial<RecordModel> | null = null;
  export let collection: string;
  export let sort = "name";
  export let key = "name";
  export let filter = (input: string): string => {
    return `name ~ "${input}"`;
  };
  export let displayNameFn = (record: Partial<RecordModel>): string => {
    return record.name;
  };
  export let crud = true;
  export let entityName = "categoría";
  const dispatcher = createEventDispatcher<{
    select: Partial<RecordModel> | RecordModel;
    clear: void;
  }>();
  let loading = false;
  let data: RecordModel[] = [];
  let showDropdown = false;

  $: if (input?.length > 25) {
    input = input.slice(0, 25);
  }

  async function search(input: string) {
    if (input) {
      showDropdown = true;
    }
    loading = true;
    await new Promise((resolve) => {
      setTimeout(() => {
        resolve(null);
      }, 200);
    });
    data = await pb.collection(collection).getFullList({
      filter: filter(input.trim()),
      sort,
    });
    await new Promise((resolve) => {
      setTimeout(() => {
        resolve(null);
      }, 200);
    });
    loading = false;
  }

  $: if (browser) {
    search(input);
  }

  let inputHeight = 0;

  function cleanSelected() {
    selected = null;
    loading = false;
    showDropdown = false;
    input = "";
    dispatcher("clear");
  }

  function select(item: Partial<RecordModel>) {
    return () => {
      selected = item;
      input = displayNameFn(item);
      tick().then(() => {
        loading = false;
        showDropdown = false;
      });
      dispatcher("select", item);
    };
  }

  function focusOut() {
    if (input) {
      select({ id: "", name: input });
    }
    showDropdown = false;
  }
</script>

<div
  class="flex relative flex-col pointer-events-auto"
  on:focusin={() => {
    if (input) {
      showDropdown = true;
    }
  }}
  use:clickOutside={focusOut}
  bind:clientHeight={inputHeight}
>
  <div class="flex w-full relative">
    {#if selected}
      <div class="absolute right-0 top-0 z-20 h-full items-center flex p4">
        <button class="flex" on:click={cleanSelected}>
          <un-i-carbon-trash-can data-un-text-red />
        </button>
      </div>
    {/if}
    {#if selected}
      <input
        type="search"
        value={displayNameFn(selected) + (!selected.id ? " *" : "")}
        class="p-2 rounded w-full flex border dark:bg-dark-700 dark:border-dark-50"
        disabled
      />
    {:else}
      <input
        type="search"
        bind:value={input}
        class="p-2 rounded w-full flex border dark:bg-dark-700 dark:border-dark-50"
        {placeholder}
      />
    {/if}
  </div>
  {#if showDropdown && ((crud && input) || data.length)}
    <div
      class="absolute bottom-0 w-full z-20"
      transition:fly={{ y: "0.5rem", duration: 400, easing: expoOut }}
      style:top="{inputHeight + 2}px"
    >
      <div
        class="w-full max-h-64 dark:border-dark-50 border-2 shadow-lg bg-white dark:bg-dark-800 rounded overflow-x-hidden overflow-y-auto flex flex-col"
      >
        <div
          class="flex flex-col w-full divide-y divide-dashed dark:divide-dark-50"
        >
          {#if input && crud}
            <button
              class="dropdown-item"
              type="button"
              on:click={select({ id: "", [key]: input })}
            >
              <div class="flex space-x-2 items-center">
                <un-i-carbon-edit />
                <span class="text-sm"
                  >Crear {entityName} <strong>{input}</strong></span
                >
              </div>
              {#if loading}
                <un-i-svg-spinners-180-ring-with-bg class="w-4 h-4" />
              {/if}
            </button>
          {/if}
          {#each data as item}
            <button class="dropdown-item" type="button" on:click={select(item)}>
              <div class="flex space-x-2 items-center">
                <un-i-carbon-chevron-right />
                <span class="text-sm">{displayNameFn(item)}</span>
              </div>
            </button>
          {/each}
        </div>
      </div>
    </div>
  {/if}
</div>

<style>
  .dropdown-item {
    --uno: "flex w-full justify-between items-center p4 hover:bg-gray-100 hover:dark:bg-dark-400";
  }
</style>
