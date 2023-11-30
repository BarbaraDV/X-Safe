import { redirect } from "@sveltejs/kit";

export async function load({ url, locals }) {
  if (!locals.user?.admin) {
    throw redirect(301, "/app");
  }
  const page = +(url.searchParams.get("page") || "") || 1;
  const perPage = 10;
  let filters: string[] = ["deleted = true"];
  if (url.searchParams.get("category")) {
    filters.push(`category = "${url.searchParams.get("category")}"`);
  }
  if (url.searchParams.get("incident")) {
    filters.push(`incident = "${url.searchParams.get("incident")}"`);
  }
  if (url.searchParams.get("type")) {
    filters.push(`severity = "${url.searchParams.get("type")}"`);
  }
  if (url.searchParams.get("status")) {
    filters.push(`status = "${url.searchParams.get("status")}"`);
  }
  if (url.searchParams.get("keywords")) {
    filters.push(`title ~ "${url.searchParams.get("keywords")}"`);
  }
  const { items, totalItems, totalPages } = await locals.pb
    .collection("vulnerabilities")
    .getList(page, perPage, {
      filter: filters.length ? filters.join(" && ") : ``,
      expand: "author,category,incident",
    });
  return {
    title: "Vulnerabilidades deshechadas",
    perPage,
    page,
    items,
    totalItems,
    totalPages,
  };
}
