import { Token } from "./models/token";

const TEXT = "text";
const STRONG = "strong";

// \ エスケープなので次の文字を文字として認識する
// .*?
// ドット (.) は任意の1文字にマッチし、アスタリスク (*) は直前のパターンが0回以上の繰り返しにマッチすることを表します。
// クエスチョンマーク (?) の存在により、このパターンは非貪欲マッチ（non-greedy match）を表します。
// 非貪欲マッチは、可能な限り短い範囲にマッチさせることを意味します。つまり、最初にマッチする部分から検索を終了します。
const STRONG_ELM_REGXP = /\*\*(.*?)\*\*/;

const getTextElement = (id: number, text: string, parent: Token): Token => ({
  id,
  elmType: TEXT,
  content: text,
  parent,
});

const genStrongElement = (id: number, parent: Token): Token => ({
  id,
  elmType: STRONG,
  content: "",
  parent,
});

const matchWithStrongRegxp = (text: string) => text.match(STRONG_ELM_REGXP);

export { getTextElement, genStrongElement, matchWithStrongRegxp };
