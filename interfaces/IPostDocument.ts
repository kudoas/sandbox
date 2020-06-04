import * as mongoose from "mongoose";

export default interface IPostDocument extends mongoose.Document {
  postId: String;
  title: string;
  url: string;
  text: string;
  created_at: Date;
  updated_at: Date;
}
