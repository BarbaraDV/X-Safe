<script lang="ts">
  import {
    onMount,
    afterUpdate,
    createEventDispatcher,
    onDestroy,
  } from "svelte";
  import type { Quill } from "quill";

  export let placeholder = "";
  export let value = "";
  export let wordCount = 0;
  let className = "";
  export { className as class };

  export let label: string = "";
  export let error: any = undefined;
  export let required = false;

  let quill: Quill | null = null;
  let selectionIdx = 0;
  let selectionLen = 0;

  let internalValue = value;
  let previousValue = value;
  let parentDrivenChange = false;

  let externalValueDebouncer: NodeJS.Timeout | undefined;
  function updateExternalValue() {
    if (parentDrivenChange) {
      previousValue = value;
      internalValue = value;
      quill?.clipboard.dangerouslyPasteHTML(value);
      parentDrivenChange = false;
    } else {
      clearTimeout(externalValueDebouncer);
      externalValueDebouncer = setTimeout(() => {
        previousValue = internalValue;
        value = internalValue;
      }, 500);
    }
  }

  $: if (value !== internalValue && quill) {
    let changedInParent = value !== previousValue;
    if (changedInParent) {
      parentDrivenChange = true;
      updateExternalValue();
    }
  }

  afterUpdate(() => {
    if (
      typeof document != "undefined" &&
      parentDrivenChange &&
      document.activeElement?.getRootNode() === (quill as unknown as Node)
    ) {
      const selection = quill?.getSelection();
      selectionIdx = selection?.index || 0;
      selectionLen = selection?.length || 0;
      quill?.focus();
      quill?.setSelection(selectionIdx, selectionLen);
    }
  });

  let dispatch = createEventDispatcher();
  let editorEl: HTMLDivElement;
  onMount(async () => {
    console.log(value);
    if (typeof document == "undefined") return;
    let katexScript = document.createElement("script");
    katexScript.src =
      "https://cdnjs.cloudflare.com/ajax/libs/KaTeX/0.9.0-alpha1/katex.min.js";
    katexScript.integrity =
      "sha384-GR8SEkOO1rBN/jnOcQDFcFmwXAevSLx7/Io9Ps1rkxWp983ZIuUGfxivlF/5f5eJ";
    katexScript.crossOrigin = "anonymous";
    const { default: Quill } = await import("quill");

    katexScript.addEventListener("load", async () => {
      quill = new Quill(editorEl, {
        modules: {
          toolbar: [
            [{ header: [1, 2, 3, false] }],
            ["bold", "italic", "underline", "strike"],
            [{ align: [] }],
            [{ list: "ordered" }, { list: "bullet" }],
            ["blockquote", "code-block"],
            ["formula", "link"],
            ["clean"],
          ],
        },
        theme: "snow",
        placeholder: placeholder,
      });

      let Link = Quill.import("formats/link");
      class CustomLink extends Link {
        static sanitize(url: string) {
          let value = url;
          if (!url.startsWith("http://") && !url.startsWith("https://")) {
            value = "http://" + url.replace(/^\/+/, "");
          }
          return super.sanitize(value);
        }
      }
      Quill.register(CustomLink);

      quill.clipboard.dangerouslyPasteHTML(value);
      // console.log(delta);
      // if (delta) {
      //   quill?.setContents(delta);
      // }

      wordCount = (quill?.getText() || "").split(/\s+/).length - 1;

      let textDebouncer: NodeJS.Timeout | undefined;
      quill.on("text-change", () => {
        clearTimeout(textDebouncer);
        textDebouncer = setTimeout(() => {
          if (quill?.getText() !== "") {
            wordCount = (quill?.getText() || "").split(/\s+/).length - 1;
            internalValue = quill?.root.innerHTML || "";
            updateExternalValue();
            dispatch("change");
          }
        }, 300);
      });
    });

    document.head.appendChild(katexScript);
  });
</script>

<svelte:head>
  <link
    rel="stylesheet"
    href="https://cdnjs.cloudflare.com/ajax/libs/KaTeX/0.9.0-alpha1/katex.min.css"
    integrity="sha384-8QOKbPtTFvh/lMY0qPVbXj9hDh+v8US0pD//FcoYFst2lCIf0BmT58+Heqj0IGyx"
    crossorigin="anonymous"
  />
</svelte:head>

<div class="h-fit {className}">
  {#if label}
    <label
      for="editor"
      class:required
      class="font-medium leading-[160%] text-black-60">{label}</label
    >
  {/if}

  <div class="editor-wrapper" class:editor-container-error={error}>
    <div id="editor" bind:this={editorEl} />
  </div>

  {#if error}
    <div class="mb-2 flex w-full items-center justify-between">
      <p class="small mt-1 text-left !text-red-100">
        {error ? error : ""} <span>&nbsp;</span>
      </p>
    </div>
  {/if}
</div>

<style>
  .editor-container-error {
    --uno: "border-2 border-red-100/40 text-red-100";
  }

  .editor-wrapper {
    position: relative;
  }

  .editor-wrapper::after {
    --uno: "pointer-events-none absolute h-6 w-6 bg-contain bg-no-repeat";
    content: "";
    z-index: 10;
    bottom: 1px;
    right: 1px;
    /* background-image: url("/icons/chevronDownRight-orange-100.svg"); */
  }

  :global(.ql-toolbar) {
    --uno: "rounded-t-lg";
  }

  :global(.ql-container) {
    --uno: "rounded-b-lg";
  }

  :global(.ql-toolbar, .ql-container) {
    --uno: "border-gray-300 !dark:border-dark-100";
  }

  :global(.ql-fill) {
    --uno: "!dark:fill-white ";
  }

  :global(.ql-stroke) {
    --uno: "!dark:stroke-white ";
  }
</style>
