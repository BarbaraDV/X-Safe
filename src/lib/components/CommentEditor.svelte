<script lang="ts">
  import type { RecordModel } from "pocketbase";
  import RichEditor from "./RichEditor.svelte";
  import UploadImages from "./UploadImages.svelte";
  import { createEventDispatcher, onMount } from "svelte";
  import { pb, useUser } from "$lib/pb";

  export let comment: Partial<RecordModel> = {
    body: "",
  };
  let wordCount: number;
  export let oldAttachments: string[] = [];
  export let newAttachments: File[] = [];

  let files: (File | string)[] = [];
  onMount(() => {
    files = [...(comment.attachments || [])];
  });

  let loading = false;

  const user = useUser();

  const distpatcher = createEventDispatcher<{
    create: string;
    update: string;
  }>();

  $: toDelete = oldAttachments?.length
    ? (comment.attachments || []).filter(
        (f: string) => !(oldAttachments || []).find((of) => of == f)
      )
    : comment.attachments || [];
  $: console.log(toDelete);
  $: console.log("Viejos", oldAttachments);

  async function createComment() {
    if (loading) return;
    if (wordCount <= 1) {
      alert("Debes llenar los campos obligatorios");
      return;
    }
    loading = true;
    try {
      const formData = new FormData();
      formData.append("author", $user!.id);
      formData.append("body", comment.body);
      for (const [_, file] of newAttachments.entries()) {
        formData.append("attachments", file);
      }
      if (comment.id) {
        const record = await pb
          .collection("comments")
          .update(comment.id!, formData);
        await pb.collection("comments").update(comment.id, {
          "attachments-": toDelete,
        });
        distpatcher("update", record.id);
      } else {
        const record = await pb.collection("comments").create(formData);
        distpatcher("create", record.id);
      }
    } catch (err) {
      console.log(err);
    } finally {
      loading = false;
      comment = {
        body: "",
      };
      files = [];
      oldAttachments = [];
      newAttachments = [];
    }
  }
</script>

<!-- svelte-ignore a11y-label-has-associated-control -->
<label class="flex flex-col space-y-2">
  <span class="text-xs font-bold">Comentario *</span>
  <RichEditor bind:value={comment.body} bind:wordCount />
</label>
<!-- svelte-ignore a11y-label-has-associated-control -->
<div class="flex flex-col space-y-2">
  <span class="text-xs font-bold">Imágenes adjuntas</span>
  <UploadImages
    record={comment}
    bind:files
    bind:newFiles={newAttachments}
    bind:oldFiles={oldAttachments}
  />
</div>
<button
  class="bg-blue-500 p2 rounded transition-200 hover:bg-blue-400 text-white flex-inline items-center space-x-2 self-end"
  on:click={createComment}
>
  <un-i-carbon-add />
  <span class="text-xs">Comentar</span>
</button>
