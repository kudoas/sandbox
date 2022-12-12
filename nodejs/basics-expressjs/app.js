const path = require("path");

const express = require("express");
const bodyParser = require("body-parser");

const app = express();

const adminRoutes = require("./routes/admin");
const shopRoutes = require("./routes/shop");

// nextを指定しなければ2回め以降のmiddlewareは更新されない(呼び出しはされる)
// app.use("/", (req, res, next) => {
//   console.log("This is always runs!");
//   next();
// });

app.use(bodyParser.urlencoded({ extended: false }));
app.use(express.static(path.join(__dirname, "public")));

// admin.jsへ
app.use("/admin", adminRoutes);

// shop.jsへ
app.use(shopRoutes);

// 404
// app.use((req, res, next) => {
//   res.status(404).send("<h1>Page not found</h1>");
// });

app.use((req, res, next) => {
  res.status(404).sendFile(path.join(__dirname, "views", "404.html"));
});

app.listen(3000);
