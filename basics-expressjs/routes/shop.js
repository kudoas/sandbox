const express = require("express");

const router = express.Router;

// getの場合は正確なパスである必要あり
router.get("/", (req, res, next) => {
  console.log("In another middleware2");
  res.send("<h1>Hello from Express!</h1>");
});

module.exports = router;
