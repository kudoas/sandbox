import * as mongoose from "mongoose";

import IPostDocument from "../interfaces/IPostDocument";

const Schema = mongoose.Schema;

const postSchema = new Schema(
  {
    title: {
      type: String,
      required: true,
    },
    url: {
      type: String,
      required: true,
    },
    text: {
      type: String,
      required: true,
    },
  },
  { timestamps: true }
);

export default mongoose.model<IPostDocument>("Post", postSchema);
