<script lang="ts">
  import { base } from "$app/paths";
  import { queryParam } from "sveltekit-search-params";

  export let data;

  const type = queryParam("type");
  const page = queryParam("page");
  const keywords = queryParam("keywords");

  let timer: NodeJS.Timeout;
  const debounce = (v: string) => {
    clearTimeout(timer);
    timer = setTimeout(() => {
      $keywords = v;
    }, 400);
  };
</script>

<div class="flex flex-col space-y-8">
  <div
    class="flex w-full justify-between lg:items-center <lg:(space-y-2 flex-col) lg:space-x-2"
  >
    <div class="flex space-x-2 items-center text-xl">
      <un-i-octicon-issue-opened />
      <h1 class="font-bold">
        <span> Incidentes </span>
      </h1>
    </div>
  </div>
  <div
    class="grid gap-4 grid-cols-1 lg:grid-cols-2 text-sm border border-dashed p-4 dark:border-dark-50 rounded-lg"
  >
    <label class="flex flex-col space-y-2 w-full">
      <span class="text-xs font-bold">Palabras claves</span>
      <input
        type="search"
        on:keyup={({ currentTarget: { value } }) => debounce(value)}
        class="p-2 rounded w-full flex border dark:bg-dark-700 dark:border-dark-50"
        placeholder="Busca por título..."
      />
    </label>
    <label class="flex flex-col space-y-2 w-full">
      <span class="text-xs font-bold">Severidad</span>
      <select
        class="p-2 rounded w-full flex border bg-white dark:bg-dark-700 dark:border-dark-50"
        bind:value={$type}
      >
        <option value={null}>Todas</option>
        <option value="low">Baja</option>
        <option value="medium">Media</option>
        <option value="high">Alta</option>
        <option value="critic">Crítica</option>
      </select>
    </label>
  </div>

  <div class="flex flex-col space-y-8 w-full">
    <div class="flex space-x-2 items-center text-lg">
      <un-i-carbon-search />
      <h1 class="font-bold">
        <span
          >{data.totalItems} {data.totalItems == 1 ? "resultado" : "resultados"}
        </span>
      </h1>
    </div>
    {#each data.items as i (i.id)}
      <div
        class="flex flex-col space-y-8 p4 border dark:border-dark-50 rounded-lg"
      >
        <div class="flex w-full justify-between items-center space-x-2">
          <h1 class="text-xl font-bold">
            {new Date(i.date).toLocaleDateString()} -
            {i.short_description}
          </h1>
          <a
            class="bg-blue-500 p2 rounded transition-200 hover:bg-blue-400 text-white flex items-center space-x-2 text-center justify-center"
            href="{base}/incidents/{i.id}"
          >
            <span class="text-xs">Ver incidente</span>
            <un-i-carbon-arrow-up-right />
          </a>
        </div>
        <div
          class="grid gap-4 grid-cols-1 lg:grid-cols-2 text-sm border border-dashed p-4 dark:border-dark-50 rounded-lg"
        >
          <div class="flex flex-col space-y-2">
            <div class="font-bold">Severidad</div>
            <div
              class="px4 py1 rounded-full bg-gray-200 dark:bg-dark-200 <lg:hidden flex self-start items-center space-x-2"
            >
              {#if i.severity == "critic"}
                <un-i-ph-fire-simple-fill class="text-red" />
                <span>Crítica</span>
              {:else if i.severity == "high"}
                <un-i-material-symbols-priority-high class="text-orange" />
                <span>Alta</span>
              {:else if i.severity == "medium"}
                <un-i-mdi-speedometer-medium class="text-purple" />
                <span>Media</span>
              {:else if i.severity == "low"}
                <un-i-material-symbols-humidity-high class="text-blue" />
                <span>Baja</span>
              {/if}
            </div>
          </div>
          <div class="flex flex-col space-y-2">
            <div class="font-bold">Autor</div>
            <p
              class="px4 py1 rounded-full bg-gray-200 dark:bg-dark-200 <lg:hidden flex self-start"
            >
              {i.expand?.author.email}
            </p>
          </div>
        </div>
      </div>
    {/each}
  </div>
  <div class="flex justify-end w-full items-center">
    <div class="flex space-x-2 items-center">
      {#if (+($page || "") || 1) > 1}
        <button
          class="flex text-2xl"
          on:click={() => {
            $page = `${Math.max(1, (parseInt($page || "") || 1) - 1)}`;
          }}
        >
          <un-i-carbon-chevron-left />
        </button>
      {/if}
      <span>
        Página {$page || 1} de {data.totalPages}
      </span>
      {#if (+($page || "") || 1) < data.totalPages}
        <button
          class="flex text-2xl"
          on:click={() => {
            $page = `${Math.min(
              data.totalPages,
              (parseInt($page || "") || 1) + 1
            )}`;
          }}
        >
          <un-i-carbon-chevron-right />
        </button>
      {/if}
    </div>
  </div>
</div>
