require "redis"

class App
  def initialize
    @redis = Redis.new(host: "redis")
  end

  def call(env)
    key = env['REQUEST_PATH']
    result = @redis.incr key
    [200, { "Content-Type" => "text/html" }, ["key: #{key}, result: #{result}"]]
  end
end

run App.new
