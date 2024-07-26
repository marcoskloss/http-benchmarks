import Fastify from "fastify";

const fastify = Fastify();

fastify.get("/", async (req, reply) => {
  return { fastify: true };
});

fastify.listen({ port: 3001 }).then(() => {
  console.log("nodejs/fastify listening on port 3001");
});
