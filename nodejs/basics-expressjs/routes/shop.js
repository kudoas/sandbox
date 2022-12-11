const path = require("path");

const express = require("express");

const rootDir = require("../util/path");

const router = express.Router();

// getの場合は正確なパスである必要あり
// router.get("/", (req, res, next) => {
//   console.log("In another middleware2");
//   res.send("<h1>Hello from Express!</h1>");
// });

router.get("/", (req, res, next) => {
  // 直接パスを指定できないので以下をやる
  res.sendFile(path.join(rootDir, "views", "add-product.html"));
});

module.exports = router;
