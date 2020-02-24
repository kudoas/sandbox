const express = require("express");

const path = require("path");

const rootDir = require("../util/path");

const router = express.Router();

// pathは同じでも違うメソッドなら使用可能

// /admin/add-product => GET
// router.get("/add-product", (req, res, next) => {
//   console.log("In another middleware1");
//   res.send(
//     '<form action="/admin/add-product"><input type="text" name="title"><button type="submit">Add Product</button></form>'
//   );
// });

router.get("/add-product", (req, res, next) => {
  res.sendFile(path.join(rootDir, "views", "shop.html"));
});

// post: urlに直接アクセスするのはだめ
// /admin/add-product => POST
router.post("/add-product", (req, res, next) => {
  // bodyParserを使ってエンコードする必要がある
  console.log(req.body);
  res.redirect("/");
});

module.exports = router;
