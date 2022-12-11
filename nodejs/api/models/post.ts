import * as mongoose from "mongoose";

import IPostDocument from "../interfaces/IPostDocument";

const Schema = mongoose.Schema;

const postSchema = new Schema(
  {
    title: {
      type: String,
      required: true,
    },
    imageUrl: {
      type: String,
      required: true,
    },
    content: {
      type: String,
      required: true,
    },
    creator: {
      type: Object,
      required: String,
    },
  },
  { timestamps: true }
);

export default mongoose.model<IPostDocument>("Post", postSchema);
