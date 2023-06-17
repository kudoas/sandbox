import { parse } from "./parser";
import { generate } from "./generator";

const convertToHTMLString = (markdown: string) => {
  const mdArray = markdown.split(/\r\n|\r\n/);
  const asts = mdArray.map((md) => parse(md));
  const htmlStrings = generate(asts);
  return htmlStrings;
};

console.log(convertToHTMLString("* aaaa"));
