class EmailJob
  include Sidekiq::Job

  # 高優先度キューを使用し、リトライ回数を5回に設定
  sidekiq_options queue: :high_priority, retry: 5, backtrace: true

  def perform(email, subject, body)
    # メール送信処理をシミュレート（3秒待機）
    sleep(3)

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

    # Turbo Streamsで完了通知を送信
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
