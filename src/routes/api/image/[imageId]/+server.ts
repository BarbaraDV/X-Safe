export async function GET({ locals, params }) {
  const image = await locals.pb.collection("images").getOne(params.imageId);
  const url = locals.pb.getFileUrl(image, image.file, {
    token: await locals.pb.files.getToken(),
  });
  const data = await fetch(url);
  return new Response(data.body);
}
