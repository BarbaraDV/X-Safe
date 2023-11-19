export async function load({ url, locals }) {
  const page = +(url.searchParams.get("page") || "") || 1;
  const perPage = 10;
  let filters: string[] = [];
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
    title: "Incidentes",
    perPage,
    page,
    items,
    totalItems,
    totalPages,
  };
}
