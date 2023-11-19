<script lang="ts">
  import { invalidateAll } from "$app/navigation";
  import { pb, useUser } from "$lib/pb";
  import { createEventDispatcher } from "svelte";
  import CommentEditor from "./CommentEditor.svelte";
  import type { RecordModel } from "pocketbase";
  import Gallery from "./Gallery.svelte";

  const dispatcher = createEventDispatcher<{
    create: string;
  }>();

  async function del(id: string) {
    await pb.collection("comments").delete(id);
    invalidateAll();
  }

  let user = useUser();

  let edit: string | null = null;

  export let comments: RecordModel[];

  $: commentsOrdered = (comments || []).sort(
    ({ created: a }: RecordModel, { created: b }: RecordModel) =>
      a > b ? -1 : a < b ? 1 : 0
  );
</script>

<div class="flex flex-col space-y-2 pt-4">
  <div class="font-bold pb-4">
    {(comments || []).length} comentarios
  </div>
  <div
    class="bg-light-200 border dark:border-dark-50 dark:bg-dark-600 flex p4 flex-col space-y-4 rounded-lg"
  >
    <CommentEditor on:create={(id) => dispatcher("create", id.detail)} />
  </div>
  <div class="flex flex-col space-y-2">
    {#each commentsOrdered || [] as comment}
      <div
        class="bg-light-200 border dark:border-dark-50 dark:bg-dark-600 flex p4 flex-col space-y-4 rounded-lg"
      >
        <div
          class="flex w-full justify-between lg:items-center <lg:(space-y-2 flex-col) lg:space-x-2"
        >
          <div class="grid gap-2 grid-cols-1 lg:grid-cols-2">
            <div class="flex flex-col space-y-2 text-xs">
              <div class="font-bold">Autor</div>
              <p
                class="px4 py1 rounded-full bg-gray-200 dark:bg-dark-200 flex self-start"
              >
                {comment.expand?.author.email}
              </p>
            </div>
            <div class="flex flex-col space-y-2 text-xs">
              <div class="font-bold">Fecha</div>
              <p
                class="px4 py1 rounded-full bg-gray-200 dark:bg-dark-200 flex self-start"
              >
                {new Date(comment.created).toLocaleString()}
              </p>
            </div>
          </div>
          {#if $user?.admin || comment.expand?.author.id == $user?.id}
            <div
              class="flex <lg:flex-col <lg:space-y-2 lg:space-x-2 items-center <lg:w-full <lg:pt-2"
            >
              {#if comment.expand?.author.id == $user?.id && edit != comment.id}
                <button
                  class="bg-blue-500 p2 rounded transition-200 hover:bg-blue-400 text-white flex items-center space-x-2 text-center justify-center <lg:w-full"
                  on:click={() => {
                    edit = comment.id;
                  }}
                >
                  <un-i-carbon-pen />
                  <span class="text-xs">Modificar comentario</span>
                </button>
              {/if}
              <button
                class="bg-red-500 p2 rounded transition-200 hover:bg-red-400 text-white flex items-center space-x-2 <lg:w-full items-center justify-center"
                type="button"
                on:click={() => del(comment.id)}
              >
                <un-i-carbon-trash-can />
                <span class="text-xs">Eliminar</span>
              </button>
            </div>
          {/if}
        </div>
        {#if edit == comment.id}
          <CommentEditor
            {comment}
            on:update={async () => {
              await invalidateAll();
              edit = null;
            }}
          />
        {:else}
          <div class="ql-editor !p0">
            {@html comment.body}
          </div>
          {#if comment.attachments?.length}
            <Gallery record={comment} images={comment.attachments} />
          {/if}
        {/if}
      </div>
    {/each}
  </div>
</div>
