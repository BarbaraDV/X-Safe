<script lang="ts">
  import { goto } from "$app/navigation";
  import { base } from "$app/paths";
  import { pb, useUser } from "$lib";
  import IncidentEditor from "$lib/components/IncidentEditor.svelte";
  export let data;

  const user = useUser();

  async function update() {
    const required = ["date", "short_description", "severity"];
    if (required.some((r) => !String(data.incident[r] || "").trim())) {
      alert("Debes llenar los campos obligatorios");
      return;
    }
    const created = await pb.collection("incidents").update(data.incident.id, {
      ...data.incident,
      id: undefined,
      date: new Date(data.incident.date).toUTCString(),
      author: $user!.id,
    });
    goto(`${base}/incidents/${created.id}`);
  }

  async function del() {
    await pb.collection("incidents").delete(data.incident.id);
    goto(`${base}/incidents`);
  }
</script>

<form class="flex flex-col space-y-8" on:submit={update}>
  <div
    class="flex w-full justify-between lg:items-center <lg:(space-y-2 flex-col) lg:space-x-2"
  >
    <h1 class="text-xl font-bold">{data.title}</h1>
    <div class="flex space-x-2 items-center <lg:w-full">
      <button
        class="bg-blue-500 p2 rounded transition-200 hover:bg-blue-400 text-white flex items-center space-x-2 <lg:w-full items-center justify-center"
      >
        <un-i-carbon-upload />
        <span class="text-xs">Subir cambios</span>
      </button>
      <button
        class="bg-red-500 p2 rounded transition-200 hover:bg-red-400 text-white flex items-center space-x-2 <lg:w-full items-center justify-center"
        type="button"
        on:click={del}
      >
        <un-i-carbon-trash-can />
        <span class="text-xs">Eliminar</span>
      </button>
    </div>
  </div>
  <IncidentEditor bind:incident={data.incident} />
</form>
