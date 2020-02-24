const express = require("express");

const router = express.Router();

router.get("/add-product", (req, res, next) => {
  console.log("In another middleware1");
  res.send(
    '<form action="/product"><input type="text" name="title"><button type="submit">Add Product</button></form>'
  );
});

// post: urlに直接アクセスするのはだめ
router.post("/product", (req, res, next) => {
  // bodyParserを使ってエンコードする必要がある
  console.log(req.body);
  res.redirect("/");
});

module.exports = router;
