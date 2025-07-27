class AlwaysErrorMiddleware
  include Sidekiq::ClientMiddleware

  def call(job_class_or_string, job, queue, redis_pool)
    Rails.logger.info "AlwaysErrorMiddleware: #{job_class_or_string} is being processed with job: #{job.inspect} in queue: #{queue}"

    raise StandardError, "This is a test error from AlwaysErrorMiddleware"
    yield
  end
end
