class ServerLoggingMiddleware
  include Sidekiq::ServerMiddleware

  def call(worker, job, queue)
    custom_text = job["aaa"]
    Rails.logger.info "[SERVER MIDDLEWARE] Processing job #{job['jid']} in queue '#{queue}' with custom text: #{custom_text}"

    yield

    Rails.logger.info "[SERVER MIDDLEWARE] Completed job #{job['jid']} with custom text: #{custom_text}"
  end
end
