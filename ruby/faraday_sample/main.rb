require "faraday"

# issue: https://github.com/lostisland/faraday/issues/1346
Faraday.get("google.com")
# => [x] returns NoMethodError
# => [o] Failed to open TCP connection to :80 (Connection refused - connect(2) for nil port 80) (Faraday::ConnectionFailed)

# Faraday.get("http://google.com") # => this works
