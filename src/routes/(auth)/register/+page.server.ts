import { fail, redirect } from "@sveltejs/kit";
import { ClientResponseError, type RecordModel } from "pocketbase";

export const actions = {
  default: async ({ locals, request }) => {
    const form = await request.formData();
    const data = Object.fromEntries(form) as NullablePartial<
      Record<"name" | "email" | "password", string>
    >;
    if (!data.email || !data.password) {
      return fail(400, {
        error: "Debes llenar todos los campos",
      });
    }

    try {
      let existentUser: RecordModel | undefined;
      try {
        existentUser = await locals.pb
          .collection("users")
          .getFirstListItem(`email = "${data.email}"`);
      } catch {
        existentUser = undefined;
      }
      if (existentUser) {
        if (existentUser.blocked) {
          return fail(400, {
            error: "Usted ha sido bloqueado, contacte con un administrador",
          });
        }
        return fail(400, {
          error: "Correo tomado",
        });
      }
      const { totalItems } = await locals.pb.collection("users").getList(1, 1);
      const user = await locals.pb.collection("users").create({
        username: data.email.split("@")[0],
        email: data.email,
        emailVisibility: true,
        password: data.password,
        passwordConfirm: data.password,
        name: data.name,
        admin: totalItems == 0,
      });
      await locals.pb
        .collection("users")
        .authWithPassword(user.email, data.password);
    } catch (e) {
      console.error(e);
      if (e instanceof ClientResponseError) {
        return fail(400, { error: e.message });
      }
      throw e;
    } finally {
      locals.pb.autoCancellation(true);
    }

    throw redirect(303, "/app");
  },
};
