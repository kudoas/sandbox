import {
  genStrongElement,
  genTextElement,
  matchWithStrongRegxp,
} from "./lexer";
import { Token } from "./models/token";

const rootToken: Token = {
  id: 0,
  elmType: "root",
  content: "",
  parent: {} as Token,
};

export const parse = (markdownRow: string) => _tokenizeText(markdownRow);

const _tokenizeText = (
  textElement: string,
  initialId: number = 0,
  initialRoot: Token = rootToken
): Token[] => {
  let elements: Token[] = [];
  let parent: Token = initialRoot;
  let id = initialId;

  // originalText: "****", p: rootToken
  const _tokenize = (originalText: string, p: Token) => {
    let processingText = originalText;
    parent = p;

    while (processingText.length !== 0) {
      id += 1;
      // 正規表現で markdown を検知する処理
      // example. matchArray [ '****', '', index: 0, input: '****', groups: undefined ]
      const matchArray = matchWithStrongRegxp(
        processingText
      ) as RegExpMatchArray;

      // 正規表現にマッチしない場合、確定でTEXT
      // example. "normal"
      if (!matchArray) {
        const onlyText = genTextElement(id, processingText, parent);
        processingText = "";
        elements.push(onlyText);
        break;
      }

      // 正規表現にマッチするがマッチした最初の文字が文頭ではない場合、マッチするまでの文字はTEXT
      // example. "normal**bold**"
      if (matchArray.index! > 0) {
        id += 1;
        const text = processingText.slice(0, matchArray.index!);
        const textElm = genTextElement(id, text, parent);
        elements.push(textElm);
        processingText = processingText.replace(text, "");
        continue;
      }

      const elm = genStrongElement(id, parent);
      parent = elm;
      elements.push(elm);

      // match した配列を processingText から消す
      processingText = processingText.replace(matchArray[0], "");
      // 残りを再帰に回す
      _tokenize(matchArray[1], parent);
      parent = p;
    }
  };
  _tokenize(textElement, parent);
  return elements;
};
