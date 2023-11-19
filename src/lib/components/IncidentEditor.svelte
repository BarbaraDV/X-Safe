<script lang="ts">
  import type { RecordModel } from "pocketbase";
  import RichEditor from "./RichEditor.svelte";

  export let incident: Partial<RecordModel> = {};

  $: if (incident.description === undefined) {
    incident.description = "";
  }

  $: if (incident.short_description === undefined) {
    incident.short_description = "";
  }

  $: if (incident.date === undefined) {
    incident.date = "";
  }

  $: if (incident.severity === undefined) {
    incident.severity = "";
  }

  let internal: string = "";

  const input = (x: string) => {
    const date = new Date(x);
    internal = `${date.getFullYear()}-${date.getMonth() + 1}-${date.getDate()}`;
  };
  const output = (x: string) => {
    try {
      const date = new Date(x);
      date.setDate(date.getDate() + 1);
      incident.date = date.toUTCString();
    } catch (e) {
      console.log(e);
    }
  };

  $: input(incident.date);
  $: output(internal);
</script>

<div class="flex flex-col space-y-4">
  <p class="text-xs">
    <strong>Nota:</strong> los campos marcados con asteriscos (*) son obligatorios
  </p>
  <div class="flex items-center <lg:(flex-col space-y-2) lg:space-x-2 w-full">
    <label class="flex flex-col space-y-2 w-full">
      <span class="text-xs font-bold">Descripción corta del incidente *</span>
      <input
        type="text"
        bind:value={incident.short_description}
        class="p-2 rounded w-full flex border dark:bg-dark-700 dark:border-dark-50"
      />
    </label>
    <label class="flex flex-col space-y-2 w-full">
      <span class="text-xs font-bold">Fecha del incidente *</span>
      <input
        type="date"
        bind:value={internal}
        class="p-2 rounded w-full flex border dark:bg-dark-700 dark:border-dark-50"
      />
    </label>
  </div>
  <label class="flex flex-col space-y-2 w-full">
    <span class="text-xs font-bold">Severidad del incidente *</span>
    <select
      class="p-2 rounded w-full flex border bg-white dark:bg-dark-700 dark:border-dark-50"
      bind:value={incident.severity}
    >
      <option hidden value="">Selecciona la severidad del incidente</option>
      <option value="low">Baja</option>
      <option value="medium">Media</option>
      <option value="high">Alta</option>
      <option value="critic">Crítica</option>
    </select>
  </label>
  <!-- svelte-ignore a11y-label-has-associated-control -->
  <label class="flex flex-col space-y-2">
    <span class="text-xs font-bold">Descripción del incidente</span>
    <RichEditor bind:value={incident.description} />
  </label>
</div>
