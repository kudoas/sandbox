class ServerLoggingMiddleware
  include Sidekiq::ServerMiddleware

  def call(worker, job, queue)
    Rails.logger.info "[SERVER MIDDLEWARE] Processing job #{job['jid']} in queue '#{queue}'"

    yield

    Rails.logger.info "[SERVER MIDDLEWARE] Completed job #{job['jid']} in queue '#{queue}'"
  end
end
