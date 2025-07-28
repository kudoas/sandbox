class ErrorSuppressorMiddleware
  include Sidekiq::ClientMiddleware

  def call(job_class_or_string, job, queue, redis_pool)
    Rails.logger.info "[ERROR SUPPRESSOR] Processing #{job_class_or_string} in queue #{queue}"

    if queue == "scheduled"
      begin
        Rails.logger.info "[ERROR SUPPRESSOR] Suppressing error for scheduled job: #{job_class_or_string}"
        yield
      rescue => e
        Rails.logger.warn "[ERROR SUPPRESSOR] Caught and suppressed error in scheduled queue: #{e.message}"
        job
      end
    else
      yield
    end
  end
end
