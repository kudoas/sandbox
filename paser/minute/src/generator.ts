import { Token } from "./models/token";
import { MergedToken } from "./models/merged_token";

const _generateHtmlString = (tokens: Array<Token | MergedToken>) =>
  tokens
    .map((t) => t.content)
    .reverse()
    .join("");

const generate = (asts: Token[][]) => {
  const htmlStrings = asts.map((lineTokens) => {
    let rearrangeAst: Array<Token | MergedToken> = lineTokens.reverse();
    return _generateHtmlString(rearrangeAst);
  });
  return htmlStrings.join("");
};

export { generate };
