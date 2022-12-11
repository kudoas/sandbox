const path = require("path");

const express = require("express");
const bodyParser = require("body-parser");
// handlerbar
const expressHbs = require("express-handlebars");

const app = express();

// handlerbar
// app.engine(
//   "hbs",
//   expressHbs({
//     layoutsDir: "views/layouts/",
//     defaultLayout: "main-layout",
//     extname: "hbs"
//   })
// );
// app.set("view engine", "handlebars");

// ejs
// app.set("view engine", "ejs");
// particleを利用可能(Reactでいうところのコンポーネント的な)

// expressが対応しているtemplate engineなら簡単に導入可能
app.set("view engine", "pug");
// viewsフォルダをviewsに指定
app.set("views", "views");

const adminData = require("./routes/admin");
const shopRoutes = require("./routes/shop");

app.use(bodyParser.urlencoded({ extended: false }));
app.use(express.static(path.join(__dirname, "public")));

app.use("/admin", adminData.routes);
app.use(shopRoutes);

app.use((req, res, next) => {
  // res.status(404).sendFile(path.join(__dirname, "views", "404.html"));
  res.status(404).render("404", { pageTitle: "Page Not Found" });
});

app.listen(3000);
