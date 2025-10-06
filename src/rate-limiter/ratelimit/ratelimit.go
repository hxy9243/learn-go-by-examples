package ratelimit

import (
	"context"
	"time"

	"github.com/go-redis/redis/v8"
	"github.com/google/uuid"
)

// Lua script to execute on the Redis server
const luaScript = `
	local key = KEYS[1]
	local value = KEYS[2]
	local limit = tonumber(ARGV[1])
	local window = tonumber(ARGV[2])
	local now = tonumber(ARGV[3])

	-- remove keys that are out of window
	redis.call('ZREMRANGEBYSCORE', key, 0, now - 1000 * window)

	--
	local current = redis.call('ZCARD', key)

	if current < limit then
		redis.call('ZADD', key, now, value)
		redis.call('EXPIRE', key, 2 * window)
		return 1
	else
		return 0
	end
`

type RateLimiterClient struct {
	redisConn   *redis.Client
	redisScript *redis.Script

	limit  int64
	window int64
}

func NewRateLimiterClient(redisConn *redis.Client, limit, window int64) *RateLimiterClient {
	return &RateLimiterClient{
		redisConn:   redisConn,
		redisScript: redis.NewScript(luaScript),
		limit:       limit,
		window:      window,
	}
}

func (r *RateLimiterClient) Allow(ctx context.Context, key string) (bool, error) {
	uniqueValue := uuid.New().String()

	val, err := r.redisScript.Run(ctx, r.redisConn, []string{key, uniqueValue}, r.limit, r.window, time.Now().UnixMilli()).Result()
	if err != nil {
		return false, err
	}

	if val.(int64) > 0 {
		return true, nil
	}
	return false, nil
}
