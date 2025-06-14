class EmailJob
  include Sidekiq::Job
  include ProgressTrackable

  # 高優先度キューを使用し、リトライ回数を5回に設定
  sidekiq_options queue: :high_priority, retry: 5, backtrace: true

  def perform(email, subject, body)
    total_steps = 4

    # ステップ1: メール設定
    update_progress(1, total_steps, "メール設定を準備中...")
    sleep(1)

    # ステップ2: 内容生成
    update_progress(2, total_steps, "メール内容を生成中...")
    sleep(1)

    # ステップ3: 送信処理
    update_progress(3, total_steps, "メールを送信中...")
    sleep(1)

    # ログに出力
    Rails.logger.info "📧 EmailJob completed!"
    Rails.logger.info "Email: #{email}"
    Rails.logger.info "Subject: #{subject}"
    Rails.logger.info "Body: #{body}"
    Rails.logger.info "Sent at: #{Time.current}"

    # 結果をファイルに保存（デモ用）
    result_file = Rails.root.join("tmp", "job_results.txt")
    File.open(result_file, "a") do |file|
      file.puts "#{Time.current}: Email sent to #{email} - Subject: #{subject}"
    end

    # ステップ4: 完了
    complete_progress("#{email} へのメール送信が完了しました！")

    # 完了通知を送信
    broadcast_job_completion("EmailJob", email, subject)
  end

  private

  def broadcast_job_completion(job_type, email, subject)
    # Turbo Streamsを使用してページ更新
    Turbo::StreamsChannel.broadcast_prepend_to(
      "job_notifications",
      target: "notifications",
      partial: "shared/job_notification",
      locals: {
        job_type: job_type,
        name: email,
        message: subject,
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
