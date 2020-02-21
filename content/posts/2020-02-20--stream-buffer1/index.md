---
title: node.jsでのstreamとbufferについて
category: "nodejs"
cover: ../cover-images/web-image.jpeg
author: Kudoa
---

nodejsでの重要な概念である**stream**と**buffer**についてのメモ  
正直まだ1ぐらいしか分かってないが、これからの学習に橋渡しのために残しておく

## Streamを制するものは、Node.jsを制す

> Jxckさんの記事より引用
> [Node.js の Stream API で「データの流れ」を扱う方法](http://jxck.hatenablog.com/entry/20111204/1322966453)

NodeのプログラミングにおいてStreamは非常に重要な要素。StreamなしではNodejsを語れない。

NodejsではデータがI/O処理をデータまるごと行うだけでなく、破片(chuck)にして行うことができる。
streamではchuckを読み込むごとに使用することができる。一方、bufferを使用する場合はbufferにchuckを蓄えてから使用できる。

![stream-img](https://kudoa-image-store.s3-ap-northeast-1.amazonaws.com/nodejs/stream-buffer.jpeg)

試しにこんなコードを書いてみた。
以下はformのmessageをchuckで読み込んで、bodyに保存しパースしてテキストファイルに書き込む流れ

```javascript
const http = require("http");
const fs = require("fs");

// event loop: 実行できるコードがある限り実行しづつける
const server = http.createServer((req, res) => {
  // console.log(req.url, req.method, req.headers);
  // process終了後にループを抜ける
  // process.exit( );

  // urlによるルーティングの設定
  const url = req.url;
  const method = req.method;

  if (url === "/") {
    res.write("<html>");
    res.write("<header><title>Enter Message</title></header>");
    res.write(
      "<body><form action='/message' method='POST'><input type='text' name='message'><button type='submit'>Send</button></form></body>"
    );
    res.write("</html>");
    return res.end();
  }

  // formでPOSTした際のリダイレクト先を指定
  if (url === "/message" && method === "POST") {
    // Parsing Request Bodies(stream, buffer)
    const body = [];
    req.on("data", chuck => {
      console.log("chuck", chuck);
      body.push(chuck);
    });
    req.on("end", () => {
      // bufferに読み込まれたchuckを文字列に戻す
      const parsedBody = Buffer.concat(body).toString();
      console.log("parsedBody", parsedBody);
			
      const message = parsedBody.split("=")[1];
      // writeFileSync: 同期的にファイル書き込みを行う
      fs.writeFileSync("messagge.txt", message);
    });
    // 302(リダイレクト)を返す
    res.statusCode = 302;
    res.setHeader("Location", "/");
    return res.end();
  }
  // headerを設定しすると設定に合った書き込みが可能
  res.setHeader("Content-Type", "text/html");
  res.write("<html>");
  res.write("<header><title>First Page</title></header>");
  res.write("<body><h1>Hello from my Node.js Server!</h1></body>");
  res.write("</html>");
  res.end();
});

server.listen(3000);
```

formに`hoge`と入力すると以下の内容がコンソールに出力される

```txt
chuck <Buffer 6d 65 73 73 61 67 65 3d 68 6f 67 65>
parsedBody message=hoge
```

chuckは2文字ずつの英数字で書かれていた。
bufferに溜まっているのでこれをパースして、`hoge`部分だけ抜き出してテキストファイルに書き込む

ここだけ切り取ってもまだまだ深そうな分野である。