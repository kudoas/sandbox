---
title: Nodejsで環境変数を設定できるdotenvの紹介
category: "nodejs"
cover: ../cover-images/web-image.jpeg
author: Kudoa
---

Node.jsでAPIを作っていた時にProduction KeyやDBのパスワードをハードコーディングしてGitHubに上げてしまうのは気が引けた。調べてみるとNodejsには環境変数を簡単に設定できる**dotenv**というモジュールがあったので使い方のメモ

## dotenvの使い方

yarnで`dotenv`をインストールするだけで使用できる。

```bash
$ yarn add dotenv
```

あとはメインのファイルで以下のように使用する。

```typescript
import * as dotenv from "dotenv";

dotenv.config();
```

たったのこれだけ。こうすることで開発環境では.envファイルから`process.env`経由で環境変数が読み込まれるし、運用環境ではホスティング先で環境変数を設定すればよい。

以下の例はMongooseを利用したMongoDBとの接続方法について。

### .envファイル

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

TypeScriptだとこのような感じで使用できる。
pushする際は`.env`ファイルをignoreすることをお忘れなく。