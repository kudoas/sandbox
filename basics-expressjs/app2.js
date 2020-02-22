const express = require("express");

const app = express();

// nextを指定しなければ2回め以降のmiddlewareは更新されない(呼び出しはされる)
app.use("/add-product", (req, res, next) => {
  console.log("In another middleware1");
  res.send("<h1>The Add Product Page</h1>");
});

app.use("/", (req, res, next) => {
  console.log("In another middleware2");
  res.send("<h1>Hello from Express!</h1>");
});

app.listen(3000);
