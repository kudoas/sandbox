import * as bodyParser from "body-parser";
import * as dotenv from "dotenv";
import * as express from "express";
import * as morgan from "morgan";
import { NextFunction, Request, Response } from "express";
import * as mongoose from "mongoose";

import Post from "./models/post";

dotenv.config();

const app: express.Express = express();

// middlewares
app.use(morgan("dev")); // logging
app.use(bodyParser.json());
app.use((_, res: Response, next: NextFunction) => {
  res.setHeader("Access-Control-Allow-Origin", "*");
  res.setHeader("Access-Control-Allow-Methods", "GET, POST, PUT, PATCH, DELETE");
  res.setHeader("Access-Control-Allow-Headers", "Content-Type, Authorization");
  next();
});

// get all posts
app.get("/post", async (req: Request, res: Response, next: NextFunction) => {
  try {
    const posts = await Post.find();
    res.status(200).json(posts);
  } catch (err) {
    res.status(500).send({ message: "Internal Server Error" });
  }
});

// get post
app.get("/post/:id", async (req: Request, res: Response, next: NextFunction) => {
  try {
    const post = await Post.findOne({ _id: req.params.id });
    res.status(200).json({ post: post });
  } catch {
    res.status(404).send({ message: res.statusMessage });
  }
});

// create post
app.post("/post", async (req: Request, res: Response, next: NextFunction) => {
  const title = req.body.title;
  const url = req.body.url;
  const text = req.body.text;
  const post = new Post({
    title: title,
    url: url,
    text: text,
  });
  try {
    await post.save();
    res.status(201).json({
      message: "Post created successfully",
      post: post,
    });
  } catch {
    res.status(500).send({ message: "Internal Server Error" });
  }
});

// update post
app.patch("/post/:id", async (req: Request, res: Response, next: NextFunction) => {
  try {
    const post = await Post.findOne({ _id: req.params.id });

    // form validation
    if (!req.body.title) {
      res.status(400).send({ message: "Missing title parameter" });
      return;
    }
    if (!req.body.url) {
      res.status(400).send({ message: "Missing url parameter" });
      return;
    }
    if (!req.body.text) {
      res.status(400).send({ message: "Missing text parameter" });
      return;
    }
    post.title = req.body.title;
    post.url = req.body.url;
    post.text = req.body.text;
    await post.save();
    res.status(201).json({
      message: "Post created successfully",
      post: post,
    });
  } catch {
    res.status(500).send({ message: "Internal Server Error" });
  }
});

// delete post
app.delete("/posts/:id", async (req: Request, res: Response, next: NextFunction) => {
  try {
    await Post.deleteOne({ _id: req.params.id });
    res.status(204).send();
  } catch {
    res.status(404);
    res.send({ error: "Post doesn't exist!" });
  }
});

// build server
mongoose
  .connect(process.env.MONGO_DB_URL)
  .then(() => {
    app.listen(process.env.PORT || 8080);
    console.log("Connected MongoDB!");
  })
  .catch((err) => console.log(err));
