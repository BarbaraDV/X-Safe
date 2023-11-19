<script lang="ts">
  import { goto, invalidateAll } from "$app/navigation";
  import { base } from "$app/paths";
  import { pb, useUser } from "$lib";
  import VulnerabilityEditor from "$lib/components/VulnerabilityEditor.svelte";
  import type { RecordModel } from "pocketbase";
  let vulnerability: Partial<RecordModel> = {};
  let category: Partial<RecordModel> | null = null;
  let incident: Partial<RecordModel> | null = null;
  let files: File[];
  let wordCount = 0;
  let loading = false;

  export let data;

  const user = useUser();

  async function createVulnerability() {
    if (loading) return;
    const required = ["title", "severity", "status"];
    if (
      required.some((r) => !String(vulnerability[r] || "").trim()) ||
      wordCount == 0
    ) {
      alert("Debes llenar los campos obligatorios");
      return;
    }
    loading = true;
    let createdCategory = category;
    if (category != null && !category?.id?.trim()) {
      createdCategory = await pb.collection("categories").create(
        {
          name: category.name,
        },
        {
          requestKey: "category",
        }
      );
    }
    try {
      let promises: Promise<RecordModel>[] = [];
      const formData = new FormData();
      formData.append("author", $user!.id);
      formData.append("title", vulnerability.title);
      formData.append("description", vulnerability.description);
      formData.append("severity", vulnerability.severity);
      formData.append("status", vulnerability.status);
      if (createdCategory != null) {
        formData.append("category", createdCategory?.id!);
      }
      if (incident != null) {
        formData.append("incident", incident?.id!);
      }
      for (const [_, file] of files.entries()) {
        formData.append("attachments", file);
        promises = [...promises];
      }
      const created = await pb.collection("vulnerabilities").create(formData);
      await invalidateAll();
      await goto(`${base}/vulnerabilities/${created.id}`);
    } catch (err) {
      console.log(err);
      if (createdCategory != null && !category?.id?.trim()) {
        await pb.collection("categories").delete(createdCategory.id!);
      }
      loading = false;
    }
  }
</script>

<form class="flex flex-col space-y-8" on:submit={createVulnerability}>
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
  <VulnerabilityEditor
    bind:vulnerability
    bind:wordCount
    bind:category
    bind:incident
    bind:newAttachments={files}
  />
</form>
