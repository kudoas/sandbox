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

Rails 8 + Sidekiq を直接利用した非同期処理のサンプルアプリケーションです。

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

### EmailJob（高優先度キュー）
- メール送信をシミュレート
- 3秒間の処理時間
- 高優先度キュー、リトライ5回、バックトレース有効

### Sidekiq Web UI
- ジョブの実行状況をリアルタイム監視
- キューの状態を確認
- 失敗したジョブの再実行
- パフォーマンス統計

## 🏗️ アーキテクチャ

- **Rails**: Webアプリケーション
- **Sidekiq**: バックグラウンドジョブ処理（直接利用）
- **Redis**: ジョブキューのストレージ
- **Docker**: コンテナ化された開発環境

## 📝 Sidekiq直接利用の特徴

### ApplicationJobとの違い
- `include Sidekiq::Job` でSidekiqを直接継承
- `perform_async` でジョブを非同期実行
- `sidekiq_options` で詳細な設定が可能

### 設定例
```ruby
class SampleJob
  include Sidekiq::Job
  
  sidekiq_options queue: :default, retry: 3
  
  def perform(name, message)
    # ジョブの処理
  end
end

# ジョブの実行
SampleJob.perform_async("名前", "メッセージ")
```

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

### キューの確認
```bash
# Sidekiqの統計情報
docker compose exec web rails console
> Sidekiq::Stats.new.queues
> Sidekiq::Queue.new("default").size
> Sidekiq::Queue.new("high_priority").size
```

## 🎉 楽しんでください！

このサンプルを参考に、あなたのアプリケーションにSidekiqを直接導入してみてください！
