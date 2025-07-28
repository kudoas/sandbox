class ScheduledJob
  include Sidekiq::Job

  # スケジュールされたジョブ用の設定
  sidekiq_options queue: :scheduled, retry: 2

  def perform(task_type = "health_check")
    case task_type
    when "health_check"
      perform_health_check
    when "cleanup"
      perform_cleanup
    when "daily_summary"
      perform_daily_summary
    else
      Rails.logger.warn "Unknown scheduled task type: #{task_type}"
    end
  end

  private

  def perform_health_check
    # システムヘルスチェックをシミュレート
    sleep(2)

    status = {
      timestamp: Time.current,
      memory_usage: "#{rand(50..80)}%",
      cpu_usage: "#{rand(10..40)}%",
      queue_size: 0,  # Queue size monitoring disabled for stability
      scheduled_queue_size: 0
    }

    Rails.logger.info "🏥 Health Check completed!"
    Rails.logger.info "Status: #{status}"

    # 結果をファイルに保存
    result_file = Rails.root.join("tmp", "job_results.txt")
    File.open(result_file, "a") do |file|
      file.puts "#{Time.current}: [SCHEDULED] Health Check - Memory: #{status[:memory_usage]}, CPU: #{status[:cpu_usage]}"
    end

    # Turbo Streamsで完了通知を送信
    broadcast_scheduled_job_completion("Health Check", status)
  end

  def perform_cleanup
    # 古い結果ファイルのクリーンアップをシミュレート
    sleep(1)

    result_file = Rails.root.join("tmp", "job_results.txt")
    if File.exist?(result_file)
      lines = File.readlines(result_file)
      # 最新50件のみ保持
      if lines.length > 50
        File.write(result_file, lines.last(50).join)
        cleaned_count = lines.length - 50
      else
        cleaned_count = 0
      end
    else
      cleaned_count = 0
    end

    Rails.logger.info "🧹 Cleanup completed!"
    Rails.logger.info "Cleaned #{cleaned_count} old entries"

    # 結果をファイルに保存
    File.open(result_file, "a") do |file|
      file.puts "#{Time.current}: [SCHEDULED] Cleanup - Removed #{cleaned_count} old entries"
    end

    # Turbo Streamsで完了通知を送信
    broadcast_scheduled_job_completion("Cleanup", { cleaned_count: cleaned_count })
  end

  def perform_daily_summary
    # 日次サマリーをシミュレート
    sleep(3)

    summary = {
      date: Date.current,
      total_jobs_today: rand(10..50),
      successful_jobs: rand(8..45),
      failed_jobs: rand(0..5),
      avg_processing_time: "#{rand(1.0..5.0).round(2)}s"
    }

    Rails.logger.info "📊 Daily Summary completed!"
    Rails.logger.info "Summary: #{summary}"

    # 結果をファイルに保存
    result_file = Rails.root.join("tmp", "job_results.txt")
    File.open(result_file, "a") do |file|
      file.puts "#{Time.current}: [SCHEDULED] Daily Summary - Total: #{summary[:total_jobs_today]}, Success: #{summary[:successful_jobs]}, Failed: #{summary[:failed_jobs]}"
    end

    # Turbo Streamsで完了通知を送信
    broadcast_scheduled_job_completion("Daily Summary", summary)
  end

  def broadcast_scheduled_job_completion(task_name, data)
    # Turbo Streamsを使用してページ更新
    Turbo::StreamsChannel.broadcast_prepend_to(
      "job_notifications",
      target: "notifications",
      partial: "shared/scheduled_job_notification",
      locals: {
        task_name: task_name,
        data: data,
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
