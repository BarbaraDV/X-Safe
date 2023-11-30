import { redirect } from "@sveltejs/kit";

export async function load({ url, locals }) {
  if (!locals.user?.admin) {
    throw redirect(301, "/app");
  }
  const page = +(url.searchParams.get("page") || "") || 1;
  const perPage = 10;
  let filters: string[] = ["deleted = true"];
  if (url.searchParams.get("type")) {
    filters.push(`severity = "${url.searchParams.get("type")}"`);
  }
  if (url.searchParams.get("keywords")) {
    filters.push(`short_description ~ "${url.searchParams.get("keywords")}"`);
  }
  const { items, totalItems, totalPages } = await locals.pb
    .collection("incidents")
    .getList(page, perPage, {
      filter: filters.length ? filters.join(" && ") : ``,
      expand: "author",
      requestKey: "incidents",
    });
  return {
    title: "Incidentes deschechados",
    perPage,
    page,
    items,
    totalItems,
    totalPages,
  };
}
