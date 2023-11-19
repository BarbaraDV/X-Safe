import { redirect } from "@sveltejs/kit";

export const actions = {
  default: async ({ locals }) => {
    locals.pb.authStore.clear();
    locals.user = null;
    throw redirect(303, "/");
  },
};

export async function load({ locals }) {
  locals.pb.authStore.clear();
  locals.user = null;
  throw redirect(303, "/");
}
