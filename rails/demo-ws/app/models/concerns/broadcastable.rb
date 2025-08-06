module Broadcastable
  extend ActiveSupport::Concern

  private

  def broadcast_job_status(status, job_data)
    Rails.logger.info "Broadcasting job status: #{status} for job #{job_data[:job_id]}"
    Turbo::StreamsChannel.broadcast_update_to(
      "job_status",
      target: "job-#{job_data[:job_id]}",
      partial: "jobs/job_status",
      locals: { job_data: job_data.merge(status: status, timestamp: Time.current) }
    )
  rescue => e
    Rails.logger.error "Broadcast job status error: #{e.message}"
    Rails.logger.error e.backtrace.join("\n")
  end

  def broadcast_job_log(message, job_id)
    Rails.logger.info "Broadcasting job log: #{message} for job #{job_id}"
    Turbo::StreamsChannel.broadcast_append_to(
      "job_status",
      target: "job-logs",
      partial: "jobs/job_log",
      locals: {
        message: message,
        job_id: job_id,
        timestamp: Time.current
      }
    )
  rescue => e
    Rails.logger.error "Broadcast job log error: #{e.message}"
    Rails.logger.error e.backtrace.join("\n")
  end
end
