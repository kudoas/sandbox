class SampleJob < ApplicationJob
  queue_as :default

  def perform(name, message)
    # 実際の処理をシミュレート（5秒待機）
    sleep(5)

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
  end
end
