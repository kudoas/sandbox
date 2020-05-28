import { NextFunction, Request, Response } from "express";
import { validationResult } from "express-validator";

import Post from "./../models/post";

export const getPosts = (req: Request, res: Response, next: NextFunction): void => {
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
};

export const createPost = (req: Request, res: Response, next: NextFunction): void => {
  const errors = validationResult(req);
  if (!errors.isEmpty()) {
    const error: { statusCode?: number; message: string } = new Error(
      "Validation falied, enterd data is incorrect."
    );
    error.statusCode = 422;
    throw error;
  }
  const title = req.body.title;
  const content = req.body.content;
  const post = new Post({
    title: title,
    content: content,
    imageUrl: "images/honda.jpg",
    creator: { name: "Tsubasa" },
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
      if (!err.statusCode) {
        err.statusCode = 500;
      }
      next(err);
    });
};

export const getPost = (req: Request, res: Response, next: NextFunction): void => {
  const postId = req.params.postId;
  Post.findById(postId)
    .then((post) => {
      if (!post) {
        const error: { statusCode?: number; message?: string } = new Error("Could not find post");
        error.statusCode = 404;
        console.log(11111111);
        throw error;
      }
      res.status(200).json({ message: "Post fetched", post: post });
    })
    .catch((err) => {
      if (!err.statusCode) {
        err.statusCode = 500;
      }
      next(err);
    });
};
