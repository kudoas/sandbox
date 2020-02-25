const path = require("path");

const express = require("express");
const bodyParser = require("body-parser");

const app = express();

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
  res.status(404).sendFile(path.join(__dirname, "views", "404.html"));
});

app.listen(3000);
