class SimpleJob < ApplicationJob
  include Broadcastable
  queue_as :default

  def perform(message, delay_seconds = 5)
    job_data = {
      job_id: job_id,
      message: message,
      delay_seconds: delay_seconds,
      type: "simple"
    }

    # ジョブ開始をbroadcast
    broadcast_job_status("processing", job_data)
    broadcast_job_log("SimpleJob started: #{message}", job_id)

    Rails.logger.info "SimpleJob started: #{message}"

    # 1秒ごとに進行状況をbroadcast
    delay_seconds.times do |elapsed|
      sleep(1)
      elapsed_time = elapsed + 1
      remaining_time = delay_seconds - elapsed_time
      
      progress_message = "進行状況: #{elapsed_time}秒経過 / 残り#{remaining_time}秒 (#{message})"
      broadcast_job_log(progress_message, job_id)
    end

    # ジョブ完了をbroadcast
    broadcast_job_status("completed", job_data)
    broadcast_job_log("SimpleJob completed: #{message} after #{delay_seconds} seconds", job_id)

    Rails.logger.info "SimpleJob completed: #{message} after #{delay_seconds} seconds"
  end
end
