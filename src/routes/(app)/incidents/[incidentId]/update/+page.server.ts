import { base } from "$app/paths";
import { redirect } from "@sveltejs/kit";

export async function load({ parent, locals }) {
  const { incident } = await parent();
  if (locals.user?.admin || incident.author == locals.user?.id) {
    return {
      title: "Modificar incidente",
    };
  }
  throw redirect(301, base + "/incidents/" + incident.id);
}
