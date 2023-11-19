<script lang="ts">
  import { goto } from "$app/navigation";
  import { base } from "$app/paths";
  import { pb, useUser } from "$lib";
  import IncidentEditor from "$lib/components/IncidentEditor.svelte";
  import type { RecordModel } from "pocketbase";
  let incident: Partial<RecordModel> = {};
  export let data;

  const user = useUser();

  async function createIncident() {
    const required = ["date", "short_description", "severity"];
    if (required.some((r) => !String(incident[r] || "").trim())) {
      alert("Debes llenar los campos obligatorios");
      return;
    }
    const created = await pb.collection("incidents").create({
      ...incident,
      date: new Date(incident.date).toUTCString(),
      author: $user!.id,
    });
    goto(`${base}/incidents/${created.id}`);
  }
</script>

<form class="flex flex-col space-y-8" on:submit={createIncident}>
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
    </div>
  </div>
  <IncidentEditor bind:incident />
</form>
