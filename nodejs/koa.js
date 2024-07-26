import Koa from "koa";
import Router from "@koa/router";

const app = new Koa();
const router = new Router();

router.get("/", async (ctx) => {
  ctx.body = { koa: true };
});

app.use(router.routes());

app.listen(3002);
console.log("nodejs/koa listening on port 3002");
