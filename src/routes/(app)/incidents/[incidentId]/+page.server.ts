export async function load({ parent }) {
  const { incident } = await parent();
  return {
    title: incident.short_description,
  };
}
