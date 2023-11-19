import { redirect } from "@sveltejs/kit";
import type { ListResult, RecordModel } from "pocketbase";

export async function load({ locals }) {
  if (!locals.user) {
    throw redirect(303, "/login");
  }
  const categories = await locals.pb
    .collection("active_categories")
    .getFullList({
      sort: "name",
    });
  return {
    categories,
  };
}
