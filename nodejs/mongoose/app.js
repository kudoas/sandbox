const path = require("path");

const express = require("express");
const bodyParser = require("body-parser");
const mongoose = require("mongoose");

const errorController = require("./controllers/error");
const User = require("./models/user");

const app = express();

app.set("view engine", "ejs");
app.set("views", "views");

const adminRoutes = require("./routes/admin");
const shopRoutes = require("./routes/shop");

app.use(bodyParser.urlencoded({ extended: false }));
app.use(express.static(path.join(__dirname, "public")));

app.use((req, res, next) => {
  User.findById("5ea239c88a82d6621d9753a4")
    .then((user) => {
      req.user = user;
      next();
    })
    .catch((err) => console.log(err));
});

app.use("/admin", adminRoutes);
app.use(shopRoutes);

app.use(errorController.get404);

mongoose
  .connect(
    // "mongodb+srv://kudoa:usernamekudoa@cluster0-m1vro.mongodb.net/shop?retryWrites=true&w=majority"
    "mongodb://localhost:27017/shop",
    { auth: { user: "testuser", password: "password" } }
  )
  .then((result) => {
    // console.log(result);
    User.findOne().then((user) => {
      if (!user) {
        const user = new User({
          name: "kudoa",
          email: "kudoa@xxx.com",
          cart: {
            items: [],
          },
        });
        user.save();
      }
    });
    console.log("Connected mongodb!!!");
    app.listen(3000);
  })
  .catch((err) => {
    console.log(err);
  });
