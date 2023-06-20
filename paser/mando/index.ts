import { lex, parse } from "./src/parser";

const markdown = `
text text
`;
const token = lex(markdown);
const html = parse(token);
console.log(html);
