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
      // matchArray [ '****', '', index: 0, input: '****', groups: undefined ]
      const matchArray = matchWithStrongRegxp(
        processingText
      ) as RegExpMatchArray;

      id += 1;

      if (!matchArray) {
        const onlyText = genTextElement(id, processingText, parent);
        processingText = "";
        elements.push(onlyText);
        break;
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
