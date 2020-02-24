// const http = require("http");

const express = require("express");

const app = express();

app.use((req, res, next) => {
  console.log("In the middleware");
  next(); // Allows the request to continue to the next middleware in line
});

app.use((req, res, next) => {
  console.log("In another middleware");
  // Headerにtext/htmlと指定しなくてもsend()の中身がstr型ならxpress側がdefaultで行ってくれる
  res.send("<h1>Hello from Express!</h1>");
});

// express側で省略できる
// const server = http.createServer(app);
// server.listen(3000);

app.listen(3000);
