---
title: React.jsのライフサイクルメソッドについて
category: "reactjs"
---

## はじめに

この記事ではReactのライフサイクルについて紹介します。

## 目次

- [Reactのライフサイクルとは](#1)
- [ライフサイクルメソッドとは](#2)
- [ライフサイクルメソッドの流れ](#3)
- [Mounting時のライフサイクルメソッドの実行](#4)
- [まとめ](#5)
- [参考URL](#6)

<a id=1></a>

## Reactのライフサイクルとは

<img src='https://kudoa-image-store.s3-ap-northeast-1.amazonaws.com/lifecycle/component-interval.jpeg' width=100%>

Reactのコンポーネントにはライフサイクル(lifecycle)と呼ばれる時間の流れがあります。
イメージとしては日が昇り、日中に活動し、夜になり就寝するという感じです。

このライフサイクルは3つの状態を遷移します。
それは**Mounting**・**Updateing**・**Unmouting**の3つです。
具体的には以下のようなものです。

|           | タイミングと内容                                             | 時間帯 |
| --------- | ------------------------------------------------------------ | ------------------- |
| Mounting  | コンポーネントがユーザーにレンダリングされるまでの仕込みの期間 | 日の出              |
| Updating  | コンポーネントがユーザに表示されており、ユーザーが操作できる期間 | 日中                |
| Unmouting | 他のコンポーネントに切り替え前に現在表示されているコンポーネントを破棄するための期間 | 日暮れ              |

<a id=2></a>

## ライフサイクルメソッドとは

上で説明したライフサイクルにはそれに付随した**ライフサイクルメソッド**というものがあります。
これらのメソッドはライフサイクルに合わせて順番に呼ばれます。

ただし、これらはクラスベースのコンポーネントでしか使用できないことに注意してください。

#### (補足)：関数ベースコンポーネントとクラスベースコンポーネント

```jsx
// function based component
import React from 'react';

const welcome = (props) => {
  return <h1>Hello, {props.name}</h1>
}
```

```jsx
// class based component
import React, { Component } from 'react';

class Welcome extends Component {
  render() {
    return <h1>Hello, {this.props.name}</h1>;
  }
}
```

> 先程、ライフサイクルメソッドはクラスベースコンポーネントでしか使用できないと説明しましたが、関数ベースのコンポーネントでは新機能があり、**useEffect()**というフックで一部代用できます。
> これは後に説明するcomponentDidMount, componentDidUpdate, componentWillUnmountが3つ合わさったようなものです。
> 興味がある方は以下を参照してください。
>
> React Document：[副作用フックの利用法](https://ja.reactjs.org/docs/hooks-effect.html)

<a id=3></a>

## ライフサイクルメソッドの流れ

ライフサイクルメソッドには呼ばれる順番とその役割があります。
**Mounting**・**Updateing**・**Unmouting**時によばれる主なものには以下があります。

### Mounting

<img src='https://kudoa-image-store.s3-ap-northeast-1.amazonaws.com/lifecycle/component-mount.jpeg' width=100%>

- constructor

Mounting時に一番最初に呼ばれるメソッドです。
JSXでは使われることがほとんどないため、見かけることはまれです。 
super(props)を呼び、更新がする必要性があれば更新できます。

- getDerivedFromProps

このメソッドだけは`static`です。
render()が呼ばれる前にstateの更新があるかどうかを確認します。
もしあれば更新後の`state`、無ければ`null`を返します。

- render

このメソッドだけはコンポーネントで必須です。
ReactがJSXコードを評価して、仮想DOMを構築します。

- componentDidMount

このメソッドは重要なライフサイクルメソッドの一つです。
クラスベースのコンポーネントを操作する時に非常に多く使用します。
1回目のrender()が呼ばれた時に1度だけ呼ばれ、ネットワークへのリクエストなどはこのメソッド内で行います。

### Update

<img src='https://kudoa-image-store.s3-ap-northeast-1.amazonaws.com/lifecycle/component-update.jpeg' width=100%>

- getDerivedStateFromProps

Update時に一番最初に呼ばれるメソッドです。
コンポーネントのstateを初期化します。
ただし、使用頻度は低いです。

- shouldComponentUpdate

コンポーネントの評価と再レンダリングを継続するかどうかを決定します。
つまり更新をここでキャンセル出来ます。
ただし、更新をブロックするとコンポーネントが破損する可能性があるので慎重に使う必要があります。

- render

上記と同様にレンダリングするメソッドです。
これも必須のメソッドです。

- getSnapshotBeforeUpdate

更新が発生する直前のスクロールを位置を記憶して提供する等に使うことができます。
このメソッドも使用頻度は低いです。

- componentDidUpdate

更新が完了した際に呼ばれるメソッドです。
HTTPリクエストを作成できますが、無限ループにならないように注意する必要があります。

### Unmouting

<img src='https://kudoa-image-store.s3-ap-northeast-1.amazonaws.com/lifecycle/component-unmount.jpeg' width=100%>

- componentWillUnmount

コンポーネントがUnmountされるときに呼ばれるメソッドです。
アニメーションなどを設定した場合はここで破棄します。
それにより次に新しいコンポーネントのサイクルが始まった際にも、その分のリソースを削減できます。

<a id=4></a>

## Mounting時のライフサイクルメソッドの実行

メソッドの呼ばれる順番とそれぞれの役割について説明しました。

最後にMounting時のライフサイクルメソッドが呼ばれる過程を実際にコンソールに出力させて確認してみます。
以下のコードを実際にコンパイルしてみます。

```jsx
import React, { Component } from 'react';

class App extends Component {
  // constructor
  constructor(props) {
    super(props);
    console.log('[App.js] constructor');
  }

  state = {
    persons: [
      { name: 'Ayano', age: 22 },
      { name: 'Yuka', age: 21 },
      { name: 'Ayaka', age: 21 }
    ],
  };

  // getDerivedStateFromProps
  static getDerivedStateFromProps(props, state) {
    console.log('[App.js] getDerivedStateFromProps', props);
    return state;
  }

  // componentDidMount
  componentDidMount() {
    console.log('[App.js] componentDidMount');
  }

  // render
  render() {
    console.log('[App.js] render');

    return (
      <div>
        <h1>This is App.js</h1>
        <p>React lifecycle method description</p>
        <ul>
          <li>{this.state.persons[0].name}</li>
          <li>{this.state.persons[1].name}</li>
          <li>{this.state.persons[2].name}</li>
        </ul>
      </div>
    );
  }
}

export default App;
```

すると以下のようにレンダリングされると同時に、呼ばれたメソッド名もコンソールに出力されます。

![image5](https://cdn-ak.f.st-hatena.com/images/fotolife/k/kudoa/20191021/20191021231514_original.png)

Consoleにcontructor → getDerivedStateFromProps → render → componentDidMountの順に出力されています。
今回はこれらのメソッドに機能を追加しているわけでないですが、実際にライフサイクルが存在することは確認できました。

<a id=5></a>

## まとめ

今回はライフサイクルとライフサイクルメソッドについて紹介しました。
ライフサイクルメソッドを使いこなしてアプリ開発に活かしていこうと思います。

<a id=6></a>

## 参考URL

- [stateとライフサイクル](https://ja.reactjs.org/docs/state-and-lifecycle.html#___gatsby)
- [React(v16.4) コンポーネントライフサイクルメソッドまとめ](https://qiita.com/Julia0709/items/3c3fc8d29fd2e56ed7a9)
- [React コンポーネントのライフサイクルとメソッドの役割について](https://qiita.com/koseki/items/432cd54b37cf44865dbd)
- [What does Side effects mean in React?](https://www.reddit.com/r/reactjs/comments/8avfej/what_does_side_effects_mean_in_react/)
