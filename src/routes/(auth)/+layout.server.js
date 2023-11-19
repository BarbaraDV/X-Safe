import { redirect } from "@sveltejs/kit";

export async function load({ locals, url }) {
  if (locals.user != null && !url.pathname.includes("/logout")) {
    throw redirect(303, "/");
  }

  const { totalItems } = await locals.pb.collection("users").getList(1, 1);
  console.log(totalItems);
  return {
    admin: totalItems == 0,
  };
}
