import * as mongoose from "mongoose";

export default interface IPostDocument extends mongoose.Document {
  postId: String;
  title: string;
  imageUrl: string;
  content: string;
  creator: {
    name: string;
  };
  created_at: Date;
  updated_at: Date;
}
