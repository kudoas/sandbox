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

Rails 8 + Sidekiq + Turbo Streams を使ったリアルタイム更新付き非同期処理のサンプルアプリケーションです。

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
docker compose up --build
```

3. ブラウザでアクセス
- メインアプリ: http://localhost:3000
- Sidekiq Web UI: http://localhost:3000/sidekiq

## 🎯 機能

### SampleJob（デフォルトキュー）
- 名前とメッセージを入力してジョブを実行
- 5秒間の処理をシミュレート
- デフォルトキュー、リトライ3回
- **完了時にTurbo Streamsでリアルタイム更新**

### EmailJob（高優先度キュー）
- メール送信をシミュレート
- 3秒間の処理時間
- 高優先度キュー、リトライ5回、バックトレース有効
- **完了時にTurbo Streamsでリアルタイム更新**

### 🔄 Turbo Streams リアルタイム更新
- **通知の自動追加**: ジョブ完了時に通知エリアに自動追加
- **結果リストの更新**: ジョブ結果リストも自動更新
- **ページ更新不要**: JavaScriptなしでリアルタイム更新
- **通知音**: ジョブ完了時に音で通知（ブラウザ対応時）
- **自動フェード**: 5秒後に通知が薄くなる
- **手動削除**: ×ボタンで通知を手動削除可能

### Sidekiq Web UI
- ジョブの実行状況をリアルタイム監視
- キューの状態を確認
- 失敗したジョブの再実行
- パフォーマンス統計

## 🏗️ アーキテクチャ

- **Rails**: Webアプリケーション
- **Sidekiq**: バックグラウンドジョブ処理（直接利用）
- **Redis**: ジョブキューのストレージ + Turbo Streamsのブロードキャスト
- **Turbo Streams**: サーバープッシュによるリアルタイムDOM更新
- **Docker**: コンテナ化された開発環境

## 📝 Turbo Streams の仕組み

### 1. ジョブ完了時の処理フロー
```ruby
# ジョブ内でTurbo Streamsにブロードキャスト
Turbo::StreamsChannel.broadcast_prepend_to(
  "job_notifications",
  target: "notifications",
  partial: "shared/job_notification",
  locals: { ... }
)
```

### 2. フロントエンドでの受信
```erb
<!-- Turbo Streamsを購読 -->
<%= turbo_stream_from "job_notifications" %>

<!-- 更新対象のDOM要素 -->
<div id="notifications">
  <!-- ここに通知が自動追加される -->
</div>
```

### 3. 技術スタック
- **Turbo Streams**: サーバープッシュDOM更新
- **Action Cable**: WebSocket通信（Turbo Streamsが内部利用）
- **Redis**: メッセージブローカー
- **パーシャルビュー**: 再利用可能なHTML部品

### 4. Turbo Streamsの利点
- **シンプル**: JavaScriptコード不要
- **宣言的**: ERBテンプレートで完結
- **高性能**: 必要な部分のみ更新
- **保守性**: サーバーサイドで完結

## 🔧 開発

### ログの確認
```bash
# Railsアプリのログ
docker compose logs web

# Sidekiqのログ
docker compose logs sidekiq

# Redisのログ
docker compose logs redis
```

### コンテナに入る
```bash
# Railsコンテナ
docker compose exec web bash

# Sidekiqコンテナ
docker compose exec sidekiq bash
```

### ジョブの手動実行
```bash
# Railsコンソールでジョブを実行
docker compose exec web rails console
> SampleJob.perform_async("Test User", "Manual job execution")
> EmailJob.perform_async("test@example.com", "Test Subject", "Test Body")
```

### Turbo Streamsの動作確認
```bash
# Railsコンソールで手動ブロードキャスト
docker compose exec web rails console
> Turbo::StreamsChannel.broadcast_prepend_to("job_notifications", target: "notifications", html: "<div>Test</div>")
```

### キューの確認
```bash
# Sidekiqの統計情報
docker compose exec web rails console
> Sidekiq::Stats.new.queues
> Sidekiq::Queue.new("default").size
> Sidekiq::Queue.new("high_priority").size
```

## 🎉 楽しんでください！

このサンプルを参考に、あなたのアプリケーションにTurbo Streamsを使ったリアルタイム更新機能を導入してみてください！
