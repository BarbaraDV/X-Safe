export async function load({ locals, params }) {
  const incident = locals.pb.collection("incidents").getOne(params.incidentId, {
    expand: "author,comments.author",
  });
  return {
    incident,
  };
}
