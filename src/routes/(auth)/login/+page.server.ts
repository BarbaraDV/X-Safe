import { fail, redirect } from "@sveltejs/kit";
import { ClientResponseError } from "pocketbase";

export const actions = {
  default: async ({ locals, request }) => {
    const form = await request.formData();
    const data = Object.fromEntries(form) as NullablePartial<
      Record<"email" | "password", string>
    >;
    if (!data.email || !data.password) {
      return fail(400, {
        error: "El correo y la contraseña son campos obligatorios",
      });
    }

    try {
      await locals.pb
        .collection("users")
        .getFirstListItem(`blocked = false && email = "${data.email}"`);
      await locals.pb
        .collection("users")
        .authWithPassword(data.email, data.password);
    } catch (e) {
      console.error(e);
      if (e instanceof ClientResponseError) {
        return fail(400, { error: "Correo o contraseña incorrectos" });
      }
      throw e;
    }

    throw redirect(303, "/app");
  },
};

export async function load({ parent }) {
  if ((await parent()).admin) {
    throw redirect(303, "/register");
  }
  return {};
}
