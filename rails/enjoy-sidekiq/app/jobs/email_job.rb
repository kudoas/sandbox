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
  end
end
