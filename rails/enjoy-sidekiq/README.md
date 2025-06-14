# README

This README would normally document whatever steps are necessary to get the
application up and running.

Things you may want to cover:

* Ruby version

* System dependencies

* Configuration

* Database creation

* Database initialization

* How to run the test suite

* Services (job queues, cache servers, search engines, etc.)

* Deployment instructions

* ...

# Enjoy Sidekiq 🚀

Rails 8 + Sidekiq を使った非同期処理のサンプルアプリケーションです。

## 🛠️ セットアップ

### 前提条件
- Docker
- Docker Compose

### 起動方法

1. リポジトリをクローン
```bash
git clone <repository-url>
cd enjoy-sidekiq
```

2. Docker Composeでサービスを起動
```bash
docker-compose up --build
```

3. ブラウザでアクセス
- メインアプリ: http://localhost:3000
- Sidekiq Web UI: http://localhost:3000/sidekiq

## 🎯 機能

### サンプルジョブ
- 名前とメッセージを入力してジョブを実行
- 5秒間の処理をシミュレート
- 結果をファイルに保存してWebページに表示

### Sidekiq Web UI
- ジョブの実行状況を監視
- キューの状態を確認
- 失敗したジョブの再実行

## 🏗️ アーキテクチャ

- **Rails**: Webアプリケーション
- **Sidekiq**: バックグラウンドジョブ処理
- **Redis**: ジョブキューのストレージ
- **Docker**: コンテナ化された開発環境

## 📝 使い方

1. http://localhost:3000 にアクセス
2. フォームに名前とメッセージを入力
3. 「ジョブを実行」ボタンをクリック
4. 5秒後にページを更新して結果を確認
5. Sidekiq Web UI (http://localhost:3000/sidekiq) でジョブの状況を監視

## 🔧 開発

### ログの確認
```bash
# Railsアプリのログ
docker-compose logs web

# Sidekiqのログ
docker-compose logs sidekiq

# Redisのログ
docker-compose logs redis
```

### コンテナに入る
```bash
# Railsコンテナ
docker-compose exec web bash

# Sidekiqコンテナ
docker-compose exec sidekiq bash
```

### ジョブの手動実行
```bash
# Railsコンソールでジョブを実行
docker-compose exec web rails console
> SampleJob.perform_later("Test User", "Manual job execution")
```

## 🎉 楽しんでください！

このサンプルを参考に、あなたのアプリケーションに非同期処理を導入してみてください！
