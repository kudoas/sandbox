import { Token } from "./models/token";
import { MergedToken } from "./models/merged_token";

const _generateHtmlString = (tokens: Array<Token | MergedToken>) =>
  tokens
    .map((t) => t.content)
    .reverse()
    .join("");

const isAllElmParentRoot = (tokens: Array<Token | MergedToken>) =>
  tokens.map((t) => t.parent.elmType).every((v) => v === "root");

// const _getInsertPosition = (content: string) => {
//   let state = 0;
//   const closeTagParentheses = ["<", ">"];
//   let position = 0;
//   console.log(content);
//   content.split("").some((c, i) => {
//     if (state === 1 && c === closeTagParentheses[state]) {
//       position = i;
//       return;
//     }
//     if (state === 0 && c === closeTagParentheses[state]) {
//       state++;
//     }
//   });
//   return position++;
// };

const _createMergedContent = (
  currentToken: Token | MergedToken,
  parentToken: Token | MergedToken
) => {
  let content = "";
  switch (parentToken.elmType) {
    case "strong":
      content = `<strong>${currentToken.content}</strong>`;
      break;
    case "merged":
    // const position = _getInsertPosition(parentToken.content);
    // content = `${parentToken.content.slice(0, position)}${
    //   currentToken.content
    // }${parentToken.content.slice(position)}`;
  }
  return content;
};

const generate = (asts: Token[][]) => {
  const htmlStrings = asts.map((lineTokens) => {
    // FILO
    let reversedAsts: Array<Token | MergedToken> = lineTokens.reverse();
    while (!isAllElmParentRoot(reversedAsts)) {
      let index = 0;
      while (index < reversedAsts.length) {
        // 親要素がルートなら何もしない
        if (reversedAsts[index].parent?.elmType === "root") {
          index++;
          continue;
        }
        const currentToken = reversedAsts[index];
        // currentTokenの親要素を探す
        // 一旦自分自身は除外する
        reversedAsts = reversedAsts.filter((_, t) => t !== index);
        const parentIndex = reversedAsts.findIndex(
          (t) => t.id === currentToken.parent.id
        );
        const parentToken = reversedAsts[parentIndex];
        const mergedToken: MergedToken = {
          id: parentToken.id,
          elmType: "merged",
          content: _createMergedContent(currentToken, parentToken),
          parent: parentToken.parent,
        };
        // parentToken は mergedToken に置き換える
        reversedAsts.splice(parentIndex, 1, mergedToken);
      }
    }
    return _generateHtmlString(reversedAsts);
  });
  return htmlStrings.join("");
};

export { generate };
