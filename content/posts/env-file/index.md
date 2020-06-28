---
title: Nodejsで環境変数を設定できるdotenvの紹介
category: "nodejs"
createdAt: "2020-05-26"
cover: background.png
---

Node.jsでAPIを作っていた時にProduction KeyやDBのパスワードをハードコーディングしてGitHubに上げてしまうのは気が引けた。調べてみるとNodejsには環境変数を簡単に設定できる**dotenv**というモジュールがあったので使い方のメモ

## dotenvの使い方

`dotenv`をインストールするだけで使用できる。

```bash
$ yarn add dotenv
```

そしてメインのファイルで以下のように使用する。

```typescript
import * as dotenv from "dotenv";

dotenv.config();
```

たったのこれだけでOK。こうすることで開発環境では.envファイルから`process.env`経由で環境変数が読み込まれ、本番環境ではホスティング先で環境変数を設定すればよい。

## 実装例

以下の例はMongooseを利用したMongoDBとの接続方法について。

### フォルダ構造

```reStructuredText
root
├── index.ts
└── .envt
```

### .env

```bash
MONGO_DB_URL = "DB_URL"
```

### index.ts

```typescript
import * as mongoose from "mongoose";
import * as dotenv from "dotenv";

dotenv.config();

mongoose
  .connect(process.env.MONGO_DB_URL)
  .then((result) => {
    app.listen(8080);
  })
  .catch((err) => console.log(err));
```


GitHubにpushする際は`.env`をignoreすることをお忘れなく。