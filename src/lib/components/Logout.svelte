<script>
  import { applyAction, enhance } from "$app/forms";
  import { pb } from "$lib";
</script>

<form
  method="POST"
  action="/logout"
  use:enhance={() => {
    return async ({ result }) => {
      await applyAction(result);
      pb.authStore.clear();
    };
  }}
>
  <slot />
</form>
