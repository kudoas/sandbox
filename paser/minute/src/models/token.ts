export type Token = {
  id: number; // unique id, マージ済みトークンの位置を調べるために使う
  parent: Token; // 親トークン
  elmType: string; // 要素の種別
  content: string; // トークンの中身
};
