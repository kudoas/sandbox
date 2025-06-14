class SampleJob
  include Sidekiq::Job
  include ProgressTrackable

  # Sidekiqのオプション設定
  sidekiq_options queue: :default, retry: 3

  def perform(name, message)
    total_steps = 5

    # ステップ1: 初期化
    update_progress(1, total_steps, "ジョブを初期化しています...")
    sleep(1)

    # ステップ2: データ準備
    update_progress(2, total_steps, "データを準備しています...")
    sleep(1)

    # ステップ3: メイン処理
    update_progress(3, total_steps, "メイン処理を実行中...")
    sleep(2)

    # ステップ4: 結果保存
    update_progress(4, total_steps, "結果を保存しています...")

    # ログに出力
    Rails.logger.info "🎉 SampleJob completed!"
    Rails.logger.info "Name: #{name}"
    Rails.logger.info "Message: #{message}"
    Rails.logger.info "Processed at: #{Time.current}"

    # 結果をファイルに保存（デモ用）
    result_file = Rails.root.join("tmp", "job_results.txt")
    File.open(result_file, "a") do |file|
      file.puts "#{Time.current}: Job completed for #{name} - #{message}"
    end
    sleep(1)

    # ステップ5: 完了
    complete_progress("#{name} さんのジョブが完了しました！")

    # 完了通知を送信
    broadcast_job_completion("SampleJob", name, message)
  end

  private

  def broadcast_job_completion(job_type, name, message)
    # Turbo Streamsを使用してページ更新
    Turbo::StreamsChannel.broadcast_prepend_to(
      "job_notifications",
      target: "notifications",
      partial: "shared/job_notification",
      locals: {
        job_type: job_type,
        name: name,
        message: message,
        completed_at: Time.current.strftime("%Y-%m-%d %H:%M:%S")
      }
    )

    # 結果リストも更新
    broadcast_results_update
  end

  def broadcast_results_update
    # 結果リストを更新
    job_results = read_job_results
    Turbo::StreamsChannel.broadcast_replace_to(
      "job_notifications",
      target: "job_results_content",
      partial: "jobs/results_content",
      locals: { job_results: job_results }
    )
  end

  def read_job_results
    result_file = Rails.root.join("tmp", "job_results.txt")
    return [] unless File.exist?(result_file)

    File.readlines(result_file).reverse.first(10)
  end
end
