import * as bodyParser from "body-parser";
import * as dotenv from "dotenv";
import * as express from "express";
import { NextFunction, Request, Response } from "express";
import * as mongoose from "mongoose";

import Post from "./models/post";

dotenv.config();

const app: express.Express = express();

app.use(bodyParser.json());
app.use((req: Request, res: Response, next: NextFunction): void => {
  // CORS: development localhost
  res.setHeader("Access-Control-Allow-Origin", "*");
  res.setHeader("Access-Control-Allow-Methods", "GET, POST, PUT, PATCH, DELETE");
  res.setHeader("Access-Control-Allow-Headers", "Content-Type, Authorization");
  next();
});

interface error {
  statusCode?: number;
  message?: string;
}

let error: error;

app.get("/posts", (req: Request, res: Response, next: NextFunction): void => {
  Post.find()
    .then((posts) => {
      res.status(200).json({ message: "Post fetched", posts: posts });
    })
    .catch((err) => {
      if (!err.statusCode) {
        err.statusCode = 500;
      }
      next(err);
    });
});

app.post("/post", (req: Request, res: Response, next: NextFunction): void => {
  const title = req.body.title;
  const url = req.body.url;
  const text = req.body.text;
  const post = new Post({
    title: title,
    url: url,
    text: text,
  });
  post
    .save()
    .then((result) => {
      console.log(result);
      // Content-type: application/json
      res.status(201).json({
        message: "Post created successfully",
        post: result,
      });
    })
    .catch((err) => {
      console.log(err);
    });
});

mongoose
  .connect(process.env.MONGO_DB_URL)
  .then((result) => {
    app.listen(process.env.PORT || 8080);
    console.log("connected!");
  })
  .catch((err) => console.log(err));
