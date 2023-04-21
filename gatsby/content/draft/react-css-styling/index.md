---
title: React Componentでのスタイリングの紹介
category: "reactjs"
createdAt: "2020-02-23"
cover: background.jpeg
---

## はじめに

WebデザインをリッチにするためにはCSSは欠かせません。

React.jsではJSXという記法を用いてHTMLを記述しますが、そのHTMLにCSSを当てる方法は様々あります。
この記事ではそのCSSを当てる方法を4つ紹介します。

## 目次

- [シンプルなCSSファイルの読み込み](#1)
- [CSS Modules](#2)
  - CSS Modulesを使用したclass名によるスタイリング
- [CSS in JS](#3)
  - styled-jsx
  - styled-component
- [まとめ](#4)
- [参考記事](#5)

<a id=1></a>

## シンプルなCSSファイルの読み込み

`create-react-app`でReact Applicationを作るとsrcフォルダ内のindex.cssやapp.cssはシンプルなCSSファイルの読み込みをすることでスタイルを当てています。
`create-react-app==3.3.0`のコマンドで生成されるindex.jsはデフォルトでは以下のようになっています。

```jsx
import React from "react";
import ReactDOM from "react-dom";
// カレントディレクトリにあるindex.cssの読み込み
import "./index.css";
import App from "./App";
import * as serviceWorker from "./serviceWorker";

ReactDOM.render(<App />, document.getElementById("root"));
...
serviceWorker.unregister();
```

`import "./index.css";`の部分でindex.cssファイル読み込み、ファイル全体にcssを当てています。
index.cssファイルにはcssが記述されており、単純にそれを読み込んでいるだけのシンプルな方法です。

```css
body {
  margin: 0;
  font-family: -apple-system, BlinkMacSystemFont, "Segoe UI", "Roboto", "Oxygen",
    "Ubuntu", "Cantarell", "Fira Sans", "Droid Sans", "Helvetica Neue",
    sans-serif;
  -webkit-font-smoothing: antialiased;
  -moz-osx-font-smoothing: grayscale;
}

code {
  font-family: source-code-pro, Menlo, Monaco, Consolas, "Courier New",
    monospace;
}
```

この方法はページ全体にglobal cssとして使用する場合には問題になりません。
しかし、tagやclassを指定したり疑似要素にスタイルを当てる場合には、どこでどのような上書きされるかわからなくなる可能性があるためあまり向いていません。

<a id=2></a>

## CSS Modules

### CSS Modulesを使用したclass名によるスタイリング

CSS ModulesはCSS設計でよく使われます。

React.jsで使用する場合は[css-loader](https://github.com/webpack-contrib/css-loader)のCSS Modules、[react-css-modules](https://github.com/gajus/react-css-modules)または[babel-plugin-react-css-modules](https://github.com/gajus/babel-plugin-react-css-modules) のいずれかを組み入れる必要がありますが、`react-create-app`でApplicationを立ち上げた際はデフォルトで使用できます。
CSSをそれぞれの要素に当てる際にはdivタグでclass名をつけたりしますが、CSSはグローバルスコープなので最初に書かれたclass名が上書きされてしまいます。

そのため他の名前と干渉しないようにユニークな名前を考える必要性がありますが､毎回それを気にして命名するのは大変です。
CSSファイル名の語尾を`.module.css`とするとCSS modulesが適応され自動でユニークなclass名を割り振ります。

> A **CSS Module** is a CSS file in which all class names and animation names are scoped locally by default.
>
> [CSS Modules](https://github.com/css-modules/css-modules)

以下のように使用できます。

```css
/* Post.module.css */
.Post {
  width: 250px;
  padding: 16px;
  text-align: center;
  border: 1px solid #eee;
  box-shadow: 0 2px 3px #ccc;
  margin: 10px;
  box-sizing: border-box;
  cursor: pointer;
}

.Post:hover,
.Post:active {
  background-color: #c0ddf5;
}

.Author {
  margin: 16px 0;
  color: #ccc;
}
```

```jsx
// Post.js
import React from "react";

import classes from "./Post.module.css";

const post = props => (
  <article className={classes.Post} onClick={props.clicked}>
    <h1>{props.title}</h1>
    <div className={classes.Author}>{props.author}</div>
  </article>
);

export default post;
```

ブラウザ上では表示されたHTMLを確認するとarticleタグに`class="Post_Post__2uao7"`とクラス名が振られています。
これがCSS Modules側で自動的に振られたユニークなIDです。

<a id=3></a>

## CSS in JS

CSS ModulesではCSSファイルを別に用意しますが、**CSS in JS**ではJavaScriptファイルに直接CSSを記述できます。
**CSS in JS**ではCSSファイルを用意する必要のないためフォルダ構造がスッキリ整理できますが、CSSがJSファイル内に書かれているため慣れないと理解しにくいのがデメリットになります。

今回は**syled-jsx**と**styled-component**の2つを紹介します。

### styled-jsx

jsx内に書かれるタグへCSSを適用できます。

`create-react-app`で作成したApplicationに適応するためにはpackageのインストールとejectして設定ファイルに書き込む必要があります。

```bash
# create-react-appを使用してApplicationを作成した後
# packageのインストールとeject
$ npm install --save styled-jsx
$ npm run eject
```

```json
// babelの設定の編集
"babel": {
    "presets": [
      "react-app"
    ],
  	// 以下を追加
    "plugins": [
      "styled-jsx/babel"
    ]
},
```

これで\<style jsx>タグが使用できるようになります。
このタグを使い、直接JSXにスタイルを指定して当てることができます。

```jsx
// Post.js
import React from "react";

const post = props => (
  <React.Fragment>
    <article className="Post" onClick={props.clicked}>
      <h1>{props.title}</h1>
      <div className="Author">{props.author}</div>
    </article>
    <style jsx>
      {`
        .Post {
          width: 250px;
          padding: 16px;
          text-align: center;
          border: 1px solid #eee;
          box-shadow: 0 2px 3px #ccc;
          margin: 10px;
          box-sizing: border-box;
          cursor: pointer;
        }
        .Post:hover,
        .Post:active {
          background-color: #c0ddf5;
        }
        .Author {
          margin: 16px 0;
          color: #ccc;
        }
      `}
    </style>
  </React.Fragment>
);

export default post;
```

### styled-component

これは文字通り、CSSをComponentとして記述して割り当てることができるものです。
styled-componentを使用できるpackageは様々ありますが、今回は後発でほぼすべての機能を使用できる**emotion**を使用します。

> Emotion is a performant and flexible CSS-in-JS library. Building on many other CSS-in-JS libraries, it allows you to style apps quickly with string or object styles. 
>
> [emotion](https://github.com/emotion-js/emotion)

Reactでは特に設定しなくてもinstallすれば使えます。

```bash
$ npm install --save @emotion/core @emotion/styled
```

以下のようにCSSでコンポーネントを作成して適応できます。

```jsx
import React from "react";
import styled from "@emotion/styled";

const Article = styled.article`
  width: 250px;
  padding: 16px;
  text-align: center;
  border: 1px solid #eee;
  box-shadow: 0 2px 3px #ccc;
  margin: 10px;
  box-sizing: border-box;
  cursor: pointer;
  :hover,
  :active {
    background-color: #c0ddf5;
  }
  .Author {
    margin: 16px 0;
    color: #ccc;
  }
`;

const post = props => (
  <Article className="Post" onClick={props.clicked}>
    <h1>{props.title}</h1>
    <div className="Author">{props.author}</div>
  </Article>
);

export default post;
```

articleタグとその中のclassにCSSを割り当てる`<Article />`というComponentを作り、使用しています。
このようにCSSをComponentとして定義して使用するのがstyled-componentです。
emotionでは他にもstyled-componentにpropsを渡して使い分けを行ったり、適応するタグを変更したりする等々、様々な機能があります。

詳しくはemtionの公式ドキュメントをごらんください。
使い方とそのコードが書かれておりサイト内で試すことができるので分かりやすいです。

> `styled` is a way to create React components that have styles attached to them. It’s available from [@emotion/styled](https://emotion.sh/docs/@emotion/styled). `styled` was heavily inspired by [styled-components](https://www.styled-components.com/) and [glamorous](https://glamorous.rocks/)
>
> [Emotion (Styled Components)](https://emotion.sh/docs/styled)

<a id=4></a>

## まとめ

今回はReact.jsで使われているCSSスタイリングの話でした。
CSSスタイリングはエンジニアそれぞれの考え方によっても何を使うか変わっていて、参考記事にもあげましたがCSS ModulesかCSS in JSを使うかは意見が分かれているようです。
ただしどの書き方もコードをスッキリ書ける工夫が詰まっていて学ぶだけでも面白いので、ぜひこの機会に色々手を出してみると知見も深まって楽しいかと思います。

<a id=5></a>

## 参考記事

- [styled-jsx](https://www.npmjs.com/package/styled-jsx#getting-started)
- [Reactのコンポーネントのスタイリングをどうやるか](https://qiita.com/lightnet328/items/218eb1c4a347302cc340#3-css-modules)
- [ReactにおけるCSS ModulesとCSS in JSの話](https://blog.ikeryo1182.com/react-style/)
- [堅牢なCSSをReactに手軽に実装できるstyled-jsx](https://inside.dmm.com/entry/2018/05/14/react-styled-jsx)