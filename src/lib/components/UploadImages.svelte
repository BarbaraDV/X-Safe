<script lang="ts">
  import { pb } from "$lib/pb";
  import type { RecordModel } from "pocketbase";
  import { onMount } from "svelte";

  export let files: (File | string)[] = [];
  export let record: Partial<RecordModel> = {};

  export let newFiles: File[];
  export let oldFiles: string[];

  $: newFiles = files.filter((f) => typeof f != "string") as File[];
  $: oldFiles = files.filter((f) => typeof f == "string") as string[];

  let id: string;
  onMount(() => {
    id = crypto.randomUUID();
  });

  async function previewFile(file: File | string) {
    if (typeof file == "string") return pb.files.getUrl(record, file);
    return new Promise<string>((resolve, reject) => {
      const reader = new FileReader();
      reader.readAsDataURL(file);
      reader.onload = () => resolve(reader.result as string);
      reader.onerror = reject;
    });
  }

  function fileLabel(file: File | string) {
    let name: string = "";
    if (typeof file != "string") {
      name = file.name;
    } else {
      name = file;
    }
    let label = name.substring(0, 15).trim();
    if (name.length > 15) {
      label += "...";
    }
    if (name.split(".").length > 1) {
      label += ` .${name.split(".").at(-1)}`;
    }
    return label;
  }
</script>

<div class="flex flex-col space-y-4">
  <div class="flex items-center justify-center w-full">
    <label
      for={id}
      class="flex flex-col items-center justify-center w-full h-64 border border-dashed rounded-lg cursor-pointer bg-gray-50 dark:hover:bg-bray-800 dark:bg-dark-800 hover:bg-gray-100 dark:border-dark-50 dark:hover:bg-dark-500"
    >
      <div class="flex flex-col items-center justify-center pt-5 pb-6">
        <un-i-carbon-document-add
          class="w-8 h-8 mb-4 text-gray-500 dark:text-gray-400"
        />
        <p class="mb-2 text-sm text-gray-500 dark:text-gray-400">
          <span class="font-semibold">Click para subir</span> o arrastra y suelta
        </p>
        <p class="text-xs text-gray-500 dark:text-gray-400">
          SVG, PNG, JPG o GIF (MAX. 800x400px)
        </p>
      </div>
      <input
        {id}
        type="file"
        class="hidden w-0 h-0"
        accept="image/png,image/jepg,image/jpg,image/gif,image/svg"
        multiple
        on:change={({ currentTarget }) => {
          if (!currentTarget.files) return;
          files = [...files, ...Array.from([...currentTarget.files])];
          currentTarget.value = "";
        }}
      />
    </label>
  </div>
  {#if files?.length}
    <div
      class="flex flex-col px2 overflow-auto max-h-80 border border-dashed dark:border-dark-50 rounded-md divide-y dark:divide-dark-50 divide-dashed"
    >
      {#each files as f, idx}
        <div
          class="flex w-full justify-between py4 items-center px2 pointer-events-auto"
        >
          <div class="flex items-center space-x-4">
            {#await previewFile(f) then data}
              <img
                class="flex w-10 h-10 rounded object-cover"
                src={data}
                alt=""
              />
            {/await}
            <p class="text-xs">
              {fileLabel(f)}
            </p>
          </div>
          <button
            class="flex items-center"
            type="button"
            on:click={() => {
              const copy = [...files];
              copy.splice(idx, 1);
              files = [...copy];
            }}
          >
            <un-i-carbon-trash-can data-un-text-2xl data-un-text-red />
          </button>
        </div>
      {/each}
    </div>
  {/if}
</div>
