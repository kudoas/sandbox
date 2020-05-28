import * as express from "express";
import { body } from "express-validator";

import * as feedController from "../controllers/feed";

export const router = express.Router();

// GET /feed/posts
router.get("/posts", feedController.getPosts);

// POST /feed/post
router.post(
  "/post",
  [body("title").trim().isLength({ min: 5 }), body("content").trim().isLength({ min: 5 })],
  feedController.createPost
);

router.get("/post/:postId", feedController.getPost);

router.put(
  "/post/:postId",
  [body("title").trim().isLength({ min: 5 }), body("content").trim().isLength({ min: 5 })],
  feedController.updatePost
);
