---
title: chromeで表示したHTMLにCSSが反映されないときの対処法
category: "chrome"
createdAt: "2020-03-23"
cover: background.jpeg
---

Chrome上でWebデザインをいじっている時にCSSを変更しても反映されないときのメモ

CSSのキャッシュが悪さをしている場合があるのでキャッシュを削除すればよい。
検証画面を開いた状態で、右上の回転矢印マークを右クリックして**キャッシュの消去とハードの読み込み**を押す。
これでキャッシュが削除され、CSSが新たに適応される。

![proceedure](https://kudoa-image-store.s3-ap-northeast-1.amazonaws.com/basic-web/%E3%82%B9%E3%82%AF%E3%83%AA%E3%83%BC%E3%83%B3%E3%82%B7%E3%83%A7%E3%83%83%E3%83%88+2020-03-23+21.47.57.png)