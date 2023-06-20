type Token = {
  type: "text";
  raw: string;
  text: string;
};

// only text
export const lex = (src: string): Token => {
  return {
    type: "text",
    raw: src,
    text: src,
  };
};

// convert token to html
export const parse = (token: Token): string => {
  let out = "";
  switch (token.type) {
    case "text": {
      out += token.text;
    }
  }

  return out;
};
