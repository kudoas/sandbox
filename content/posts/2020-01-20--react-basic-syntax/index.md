---
title: React.jsの特徴とJavaScriptの頻出構文
category: "reactjs" 
---

React.jsにおける特徴と非常によく使う文法事項のメモ

## 目次

- [今回使用するJavaScript](#1)

- [React.jsの特徴](#2)
    - Component指向形のフレームワーク
    - 仮想DOMによる高速なレンダリング


- [React.jsでよく使う文法](#3)
   - importとexport
   - constとlet
   - map()
   - 三項演算子
   - switch文
   - bind()

- [まとめ](#4)

- [参考文献・URL](#5)

<a id='1'></a>

## 今回使用するJavaScript

今回の記事では<u>ES7(ES2016)</u>以降の文法を使用します。
古いバージョンだと書き方が違うので注意して下さい。

<a id='2'></a>

## React.jsの特徴

ここではReact.jsの重要な要素である**Component指向型のフレームワーク**と**仮想DOMによる高速なレンダリング**をしている点について紹介します。

### Component指向型のフレームワーク

React.jsは**Component指向型のフレームワーク**と言われています。
Componentとは部品という意味です。
この部品を組み合わせて1つ画面を作るのがReact.jsの基本的な考え方となっています。

そもそもWebページはさまざまな部品によって構成されています。

![image1](https://kudoa-image-store.s3-ap-northeast-1.amazonaws.com/react-basic-syntax/react-component.jpeg)

通常のHTMLでデザインする場合、Header部分は\<Header>タグや\<nav>タグなどを使い、List-itemは\<li>タグや\<ul>タグなどを使います。
React.jsではComponentを1つのJavaScriptファイルや関数、クラスとして定義します。
これにより必要なComponentを一度作ってしまえば再利用が可能になるため、効率よく開発ができます。

### 仮想DOMによる高速なレンダリング

**DOM (Document Object Model)**とはブラウザ上でページを表示する際に作られるものです。
このDOMをもとに画面を表示する(レンダリング)を行います。

通常のDOM生成からレンダリングまでの手順はHTMLファイルが変更されるたびに行われます。

![image2](https://kudoa-image-store.s3-ap-northeast-1.amazonaws.com/react-basic-syntax/v-dom1.jpeg)

しかしこの場合だと新しいファイルが読み込まれるたびにすべてDOMへ変換します。
そして毎回最初からこの工程を行うので時間がかかってしまいます。
特にReact.jsのようなComponent指向型フレームワークでは画面上で変更される部分は一部なので、そこだけ表示を変更する方が効率的です。

そこでReact.jsではDOMを生成する前に**仮想DOM**を生成します。
仮想DOMが生成された時点で新しい仮想DOMと古い仮想DOMと比較を行い、そのdiffだけDOMを書き換えます。

![image3](https://kudoa-image-store.s3-ap-northeast-1.amazonaws.com/react-basic-syntax/v-dom2.jpeg)

すべてのDOMを作り直すよりも高速で画面を表示できます。

<a id='3'></a>

## React.jsでよく使う文法

ここからはReact.jsで非常によく使うJavaScriptの文法について紹介。使い方も例示します。

### importとexport

React.jsではさまざまなライブラリーやモジュールをインポートします。
その際には以下のように行います。

```javascript
// import Module名 from Moduleのパス
import React from 'react'; // React.jsではファイルごとに必須
import { Route, Switch } from "react-router-dom"; // インストール済みのパッケージはそのまま記述可能

import ImportedContainer from "./containers/ImportedContainer"; // 自分で作成したものは相対パスを使用
```

npmでプロジェクトディレクトリにインストールしたものは`import React from 'react';`のようにインポートします。
自分で作成したモジュール名はインポートする際には好きな名前をつけることができます。

また他のファイルからインポートするためには、インポートするファイルをエクスポートする必要があります。

```jsx
import React, { Component } from "react";
import ImportedContainer from "./containers/ImportedContainer";

class App extends Component {
  render() {
    return (
      <div>
				<ImportedContainer />
      </div>
    );
  }
};

// export default class名 or function名
export default App;
```

エクスポートしなければ他のファイルからインポートできないので、うまく読み込めない際は確認して下さい。

### constとlet

変数を定義する際にconstもしくはletを宣言します。

```javascript
const Person1 = 'Fuginami';
let Person2 = 'Hashimoto';
```

これらには使い分けが存在します。
constで宣言した変数は再代入ができません。

```javascript
const Person1 = 'Fuginami';

Person1 = 'Tyousyuu'; // Error
console.log(Person1);
```

変数を定義する際は<u>再代入するものなどはletで定義</u>し、<u>それ以外はconstで定義</u>するのが一般的です。
React.jsでも基本的にはconstを使い、状況に応じて変数が変わる可能性のあるものだけletを使います。

### map()

map()は配列対して使用するメソッドです。
配列の各要素1つずつに対して関数(コールバック関数)を適用し、新しい配列を作ることができます。

```javascript
const Array = [1, 3, 5, 7, 9]

const doubleNumber = (x) => {
	return x*2
}
doubleArray = Array.map(doubleNumber)
// [2, 6, 10, 14, 18]
```

また別で関数を定義せず、直接map()に関数をいれることも可能です。
この方がコードを短く書けるので使用頻度は多いかもしれません。

```JavaScript
const Array = [1, 3, 5, 7, 9]

doubleArray = Array.map(num => num*100) // Arrayの1つ1つの要素がnumに代入されて関数が適用
// [100, 300, 500, 700, 900]
```

map()はReact.jsでは非常によく使います。
例えば、何らかのAPIを使ってデータを取ってきた(フェッチ)際にそのデータをComponentに渡して出力したいケースに使用します。

```jsx
import React, { Component } from "react";
import Order from "./components/Order";

class Orders extends Component {
  state = {
    orders: []
  };

  componentDidMount() {
		// ここでAPIからデータを受けとってstateのordersに入れておく
    ...
  }

  render() {
    return (
      <div>
        <!-- 先程APIから受けとってordersに入れて要素をOrderというComponentに代入する-->
        {this.state.orders.map(order => (
          <Order
            id={order.id}
						item={order.items} // ここで割り当てる
          />
        ))}
      </div>
    );
  }
}
```

配列を取り出してComponentに入れる作業をたった数行で行えるので非常に重宝します。

### 三項演算子

三項演算子はif文を省略して書く記法です。
通常、if文で条件式がtrue時とfalse時を分岐させて記述します。

```javascript
const age = 22;

if (age >= 18) {
  console.log("You can ride a motorcycle.")
} else {
  console.log("You can't ride a motorcycle!!!")
}
```

これを三項演算子では一行で書けます。

```javascript
// 条件式？条件がtrueのときに実行する処理：条件式がfalseのときに実行する処理
age >= 18 ? console.log("You can ride a motorcycle.") : console.log("You can't ride a motorcycle!!!")
```

これもReact.jsではよく使います。
三項演算子を使うとアプリケーションの状態に応じてHTMLに作用させるCSSを変更できます。

例えばModalを作りたい場合は以下のようにCSSを割り当てます。

```jsx
<Modal style={{
  // transform: 表示位置の移動
	transform: this.state.show ? "translateY(0)" : "translateY(-100vh)", 
  // opacity: 透明度
	opacity: this.state.show ? "1" : "0" 
  }} />
```

アプリケーションに何らかの変化(ボタンを押すなど)が起きてthis.state.showがtrueになった場合はModalのCSSに`transform: translateY(0);`と`opacity: 1`となり画面に表示させます。
falseのときは`transform: translateY(-100vh);`と`opacity: 0`なので画面外に透明な状態で隠れていることになります。

![image4](https://kudoa-image-store.s3-ap-northeast-1.amazonaws.com/react-basic-syntax/modal.jpeg)

このようにアプリの状態に応じた処理の変更がする際によく使用されます。

### switch文

三項演算子の場合はtrueかfalseの場合の処理しかできません。
もっと条件分岐させたい場合はswtich文を使います。

```javascript
switch(条件) {
  case 値1:
    「条件=値1」である場合に実行される命令;
    break;
  case 値2:
    「条件=値2」である場合に実行される命令;
    break;
  ...
  case 値n:
    「条件=値n」である場合に実行される命令;
    break;
  default:
    条件の値がどの条件にも合致しない場合に実行される命令;
}
```

具体的には以下のように使います。

```javascript
const input = 'text'

switch(input) {
  case 'text':
    console.log('this is text');
  case 'number':
    console.log('this is number');
    break;
  default:
    console.log('What is this?');
}
```

React.jsではフォーム画面で入力のタイプ(数字、テキスト、E-mialなど)によって違う処理をしたい場合などに使用します。
三項演算子だけでカバーしきれないものはswitch文を使うと良いでしょう。

### bind()

bind()は関数に設定を追加して新しい関数を生成する関数です。
なかなか分かりづらいので、最初に例を示します。

```javascript
class Person {
  // プロパティ：クラス内の情報. constructor()はプロパティを定義するための関数
  constructor() {
   this.name = 'Yui' 
  }

  saySomething = () => {
    console.log('Hi, I am ' + this.name)
  }
}

class Dog {
  constructor() {
   this.name = 'Pochi' 
  }
}

const yui = new Person()
const pochi = new Dog()

// bind()は関数を返しているので最後の実行の()は忘れずに
yui.saySomething.bind(pochi)() // Hi, I am Pochi
```

`const pochi = new Dog()`で定義したDogクラスのプロパティをbind()を用いてPersonクラスに紐付けています。
それによってPersonのプロパティで`name='Yui'`と定義したものが`'Pochi'` に上書きされました。
このようにbind()は関数と関数をbind(紐付け)できます。

React.jsではイベントハンドラを定義する際に使用します。
公式ドキュメントにはこのように書かれています。

> JSX のコールバックにおける `this` の意味に注意しなければなりません。JavaScript では、クラスのメソッドはデフォルトでは[バインド](https://developer.mozilla.org/en/docs/Web/JavaScript/Reference/Global_objects/Function/bind)されません。`this.handleClick` へのバインドを忘れて `onClick` に渡した場合、実際に関数が呼ばれた時に `this` は `undefined` となってしまいます。
>
> イベント処理：https://ja.reactjs.org/docs/handling-events.html

つまり何らかの関数をクラス型のコンポーネントで使用する際にはconctructorで`handleClick`をbindする必要があります。

```jsx
class LoggingButton extends React.Component {
  constructor(props) {
    super(props);
    // ここでthisをbind
    this.handleClick = this.handleClick.bind(this);
  }
  
  handleClick() {
    console.log('this is:', this);
  }

  render() {
    return (
      <button onClick={this.handleClick}>
        Click me
      </button>
    );
  }
}
```

ただし、毎回constructor()でbind()を使ってthisにイベントハンドラを紐付けるのは面倒くさいです。
React.jsではbind()を使わなくて済む解決策があります。

1つはアロー関数を使う方法です。

```jsx
class LoggingButton extends React.Component {
  handleClick() {
    console.log('this is:', this);
  }

  render() {
    return (
      // eventを渡す必要あり
      <button onClick={(e) => this.handleClick(e)}>
        Click me
      </button>
    );
  }
}
```

この他にも回避する方法をさまざまあります。
興味があれば、以下の記事を参考にしてみてください。

> [Reactをes6で使う場合のbindの問題](https://qiita.com/cubdesign/items/ee8bff7073ebe1979936)

<a id='4'></a>

## まとめ

他にもいっぱい使うものはあると思うので要勉強です。

<a id='5'></a>

## 参考文献・URL

- [Virtual DOMの仕事ってなに?  (Reactの表示速度がはやい理由)](https://qiita.com/risagon/items/019942c60e9c3e6c05a5)

- [(初心者向け) JavaScriptのクラス (ES6対応)](https://qiita.com/tadnakam/items/ae8e0e95107e1427983f)

- [React公式イベント処理](https://ja.reactjs.org/docs/handling-events.html)