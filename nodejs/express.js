import express from "express";

const app = express();

app.get("/", (req, res) => {
  res.send({ express: true });
});

app.listen(3000, () => {
  console.log("nodejs/express listening on port 3000");
});
