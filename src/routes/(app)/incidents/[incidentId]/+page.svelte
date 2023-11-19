<script lang="ts">
  import { invalidateAll } from "$app/navigation";
  import { base } from "$app/paths";
  import { pb, useUser } from "$lib";
  import Comments from "$lib/components/Comments.svelte";
  export let data;

  const user = useUser();

  async function createComment(id: string) {
    try {
      await pb.collection("incidents").update(data.incident.id, {
        "comments+": id,
      });
      invalidateAll();
    } catch (err) {
      console.log(err);
      await pb.collection("comments").delete(id);
    }
  }
</script>

<div class="flex flex-col space-y-8">
  <div class="flex w-full justify-between items-center space-x-2">
    <h1 class="text-xl font-bold">
      {new Date(data.incident.date).toLocaleDateString()} -
      {data.incident.short_description}
    </h1>
    {#if $user?.admin || data.incident.author == $user?.id}
      <a
        class="bg-blue-500 p2 rounded transition-200 hover:bg-blue-400 text-white flex items-center space-x-2"
        href="{base}/incidents/{data.incident.id}/update"
      >
        <un-i-carbon-pen />
        <span class="text-xs">Modificar incidente</span>
      </a>
    {/if}
  </div>
  <div
    class="grid gap-4 grid-cols-1 lg:grid-cols-2 text-sm border border-dashed p-4 dark:border-dark-50 rounded-lg"
  >
    <div class="flex flex-col space-y-2">
      <div class="font-bold">Severidad</div>
      <div
        class="px4 py1 rounded-full bg-gray-200 dark:bg-dark-200 <lg:hidden flex self-start items-center space-x-2"
      >
        {#if data.incident.severity == "critic"}
          <un-i-ph-fire-simple-fill class="text-red" />
          <span>Crítica</span>
        {:else if data.incident.severity == "high"}
          <un-i-material-symbols-priority-high class="text-orange" />
          <span>Alta</span>
        {:else if data.incident.severity == "medium"}
          <un-i-mdi-speedometer-medium class="text-purple" />
          <span>Media</span>
        {:else if data.incident.severity == "low"}
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
        {data.incident.expand?.author.email}
      </p>
    </div>
  </div>
  <div class="flex flex-col space-y-2">
    {#if data.incident.description}
      <div class="font-bold">Descripción</div>
      <div
        class="ql-editor bg-light-100 border rounded-lg dark:border-dark-50 dark:bg-dark-800"
      >
        {@html data.incident.description}
      </div>
    {:else}
      <div class="font-bold">Sin descripción</div>
    {/if}
  </div>
  <Comments
    comments={data.incident.expand?.comments}
    on:create={(id) => createComment(id.detail)}
  />
</div>
