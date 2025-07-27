require Rails.root.join("app", "middleware", "server_logging_middleware")
require Rails.root.join("app", "middleware", "custom_text_middleware")

Sidekiq.configure_server do |config|
  config.redis = { url: ENV.fetch("REDIS_URL", "redis://localhost:6379/0") }
  config.server_middleware do |chain|
    chain.add ServerLoggingMiddleware
  end
  config.client_middleware do |chain|
    chain.add AlwaysErrorMiddleware
  end
end

Sidekiq.configure_client do |config|
  config.redis = { url: ENV.fetch("REDIS_URL", "redis://localhost:6379/0") }
end
