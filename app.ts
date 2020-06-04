import * as express from "express";
import { NextFunction, Request, Response } from "express";
import * as bodyParser from "body-parser";
import * as dotenv from "dotenv";

dotenv.config();

const app: express.Express = express();

app.use(bodyParser.json());
app.use((req: Request, res: Response, next: NextFunction) => {
  // CORS: development localhost
  res.setHeader("Access-Control-Allow-Origin", "*");
  res.setHeader("Access-Control-Allow-Methods", "GET, POST, PUT, PATCH, DELETE");
  res.setHeader("Access-Control-Allow-Headers", "Content-Type, Authorization");
  next();
});

app.get("/", (req: Request, res: Response, next: NextFunction) => {
  res.json([
    {
      title: "Learn to JSON",
      createAt: "2020-05-01",
      url: "https://www.youtube.com/embed/M7lc1UVf-VE",
      text: "素晴らしい",
    },
    {
      title: "Learn to SJON",
      createAt: "2020-05-03",
      url: "https://www.youtube.com/embed/M7lc1UVf-VE",
      text: "素晴らしい",
    },
    {
      title: "Learn to hogeho",
      createAt: "2020-05-01",
      url: "https://www.youtube.com/embed/M7lc1UVf-VE",
      text: "素晴らしい",
    },
  ]);
});

app.listen(8000 || process.env.PORT);
