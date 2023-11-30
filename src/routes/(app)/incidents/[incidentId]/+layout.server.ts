import { redirect } from "@sveltejs/kit";

export async function load({ locals, params }) {
  const incident = await locals.pb
    .collection("incidents")
    .getOne(params.incidentId, {
      expand: "author,comments.author",
    });
  if (incident.deleted && !locals.user?.admin) {
    throw redirect(301, "/app");
  }
  return {
    incident,
  };
}
