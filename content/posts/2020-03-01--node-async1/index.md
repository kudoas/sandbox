---
title: 非同期処理でハマったこと
category: "nodejs"
---

JavaSctipt特有の非同期通信メモ

## 問題点

ローカルファイルのjson形式のデータをパースして受け取る。
そのjsonファイルをタイトルが入ってる。

```json
// data/products.json
[{"title":"hoge"},{"title":"hgoe"}] 
```

```javascript
// models
const fs = require("fs");
const path = require("path");

module.exports = class Product {
  constructor(t) {
    this.title = t;
  }
  
  static fetchAll(cb) {
    const p = path.join(
      path.dirname(process.mainModule.filename),
      "data", 
      "products.json"
    );
    fs.readFile(p, (err, fileContent) => {
       if (err) {
         return [];
       }
       return JSON.parse(fileContent);
     });
  }
};
```

これをcontrollersで受け取り、viewにデータを渡す。

```javascript
// controllers
exports.getProducts = (req, res, next) => {
  Products.fetchAll(products => {
    res.render("shop", {
      prods: products,
      pageTitle: "Shop",
      path: "/",
      hasProducts: products.length > 0,
      activeShop: true,
      productCSS: true
    });
  });
};
```

ただし、これのコードではjsonを受け取れすにundefinedになる。
問題は以下のコード。

```javascript
static fetchAll() {
    const p = path.join(
      path.dirname(process.mainModule.filename),
      "data", 
      "products.json"
    );
  	// ①：ここが非同期処理なので処理が終わる前に通過してundefinedになる
    fs.readFile(p, (err, fileContent) => {
       // 内部の条件で値を返すが、fetchAll関数自体が値を返すわけではない。
       if (err) {
         return [];
       }
       return JSON.parse(fileContent);
     });
  }
```

データを読み込む処理①は条件内で値を返すが、関数自体が値を返している訳ではない。node.jsでは非同期通信なので①が終了して関数が値を返す前に、controllerがviewにデータを渡す処理が実行されてしまう。

## 解決案

これを解決するためには色々方法論がある。

- callback
- async/await Promise

### callback

fetchAllの引数にcallback関数を受け取る方法

```javascript
// callbackで受け取る
static fetchAll(cb) {
    const p = path.join(
      path.dirname(process.mainModule.filename),
      "data", 
      "products.json"
    );
    fs.readFile(p, (err, fileContent) => {
       // 内部の条件で値を返すが、fetchAll関数自体が値を返すわけではない。
       if (err) {
         return cb([]);
       }
       return cb(JSON.parse(fileContent));
     });
  }
```

### async/await Promise

Promiseのasync/awaitを使った方法。これを使うと非同期処理を同期処理っぽく書くことができる。

```javascript
// models
static async fetchAll() {
	const p = path.join(rootDir, "data", "products.json");
  // awaitで処理が終わるまで待ってもらう
	let data = await fs.readFile(p);
	return JSON.parse(data);
}
```

```javascript
// controller
exports.getProducts = async (req, res) => {
	const products = await Product.fetchAll();
	res.render('shop', { prods: products, docTitle: 'Shop', path: '/' });
}
```

## 所感

非同期処理は他の同期的な言語にない使用であり、なかなか理解が難しい。
後で分かったらまたまとめよう...

## 参考文献

- [How do I return the response from an asynchronous call?](https://stackoverflow.com/questions/14220321/how-do-i-return-the-response-from-an-asynchronous-call)
- [ラーメンで理解するasync/await](https://qiita.com/7tsuno/items/6d5a27ffe9143b35defe)
- [Promiseの使い方、それに代わるasync/awaitの使い方](https://qiita.com/suin/items/97041d3e0691c12f4974)

