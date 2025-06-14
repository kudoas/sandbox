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

Rails 8 + Sidekiq + Turbo Streams を使ったプログレスバー付きリアルタイム非同期処理のサンプルアプリケーションです。

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
- **5段階のプログレス表示**:
  1. ジョブ初期化
  2. データ準備
  3. メイン処理
  4. 結果保存
  5. 完了
- デフォルトキュー、リトライ3回
- **リアルタイムプログレスバー更新**

### EmailJob（高優先度キュー）
- メール送信をシミュレート
- **4段階のプログレス表示**:
  1. メール設定準備
  2. 内容生成
  3. 送信処理
  4. 完了
- 高優先度キュー、リトライ5回、バックトレース有効
- **リアルタイムプログレスバー更新**

### 📊 プログレスバー機能
- **リアルタイム更新**: Turbo Streamsで進行状況を自動更新
- **視覚的フィードバック**: アニメーション付きプログレスバー
- **詳細情報表示**: 現在のステップ、進行率、メッセージ
- **自動非表示**: 完了後に自動でプログレスバーが消える
- **スピナーアニメーション**: 処理中を示すローディングアニメーション

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
- **Redis**: ジョブキューのストレージ + プログレス情報保存 + Turbo Streamsブロードキャスト
- **Turbo Streams**: サーバープッシュによるリアルタイムDOM更新
- **ProgressTrackable**: プログレス管理用Concern
- **Docker**: コンテナ化された開発環境

## 📝 プログレスバーの仕組み

### 1. プログレス管理Concern
```ruby
module ProgressTrackable
  def update_progress(current_step, total_steps, message = nil)
    # Redisにプログレス情報を保存
    # Turbo Streamsでプログレスバーを更新
  end
end
```

### 2. ジョブでのプログレス更新
```ruby
class SampleJob
  include ProgressTrackable
  
  def perform(name, message)
    update_progress(1, 5, "初期化中...")
    # 処理...
    update_progress(2, 5, "データ準備中...")
    # 処理...
    complete_progress("完了！")
  end
end
```

### 3. Turbo Streamsでのリアルタイム更新
```ruby
# プログレスバーの更新
Turbo::StreamsChannel.broadcast_replace_to(
  "job_notifications",
  target: "progress_#{job_type}_#{job_id}",
  partial: "shared/progress_bar"
)
```

### 4. プログレス情報の永続化
- **Redis**: プログレス情報を一時保存（5分でexpire）
- **キー形式**: `job_progress:JobClass:job_id`
- **データ**: JSON形式で進行率、ステップ、メッセージを保存

## 🎨 UI/UX 特徴

### プログレスバーデザイン
- **グラデーション**: 美しいブルーグラデーション
- **アニメーション**: スムーズな進行アニメーション
- **スピナー**: 処理中を示すローディングスピナー
- **パルス効果**: プログレスバー内のパルスアニメーション

### レスポンシブ対応
- **モバイルフレンドリー**: 小画面でも見やすいデザイン
- **アクセシビリティ**: 色だけでなくテキストでも進行状況を表示

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

### プログレス情報の確認
```bash
# Redisでプログレス情報を確認
docker compose exec web rails console
> redis = Redis.new(url: ENV['REDIS_URL'])
> redis.keys("job_progress:*")
> redis.get("job_progress:SampleJob:job_id")
```

### Turbo Streamsの動作確認
```bash
# Railsコンソールで手動ブロードキャスト
docker compose exec web rails console
> Turbo::StreamsChannel.broadcast_prepend_to("job_notifications", target: "progress_area", html: "<div>Test</div>")
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

このサンプルを参考に、あなたのアプリケーションにプログレスバー付きのリアルタイム非同期処理を導入してみてください！
