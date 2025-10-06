package ratelimit

import (
	"context"
	"net/http"
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

// NewRateLimiterClient returns a new rate limiter client that uses Redis as backend for rate limiting
func NewRateLimiterClient(redisConn *redis.Client, limit, window int64) *RateLimiterClient {
	return &RateLimiterClient{
		redisConn:   redisConn,
		redisScript: redis.NewScript(luaScript),
		limit:       limit,
		window:      window,
	}
}

// Allow checks if the request exceeds the limit
func (rl *RateLimiterClient) Allow(ctx context.Context, key string) (bool, error) {
	uniqueValue := uuid.New().String()

	val, err := rl.redisScript.Run(
		ctx, rl.redisConn, []string{key, uniqueValue}, rl.limit, rl.window, time.Now().UnixMilli(),
	).Result()
	if err != nil {
		return false, err
	}

	if val.(int64) > 0 {
		return true, nil
	}
	return false, nil
}

// Middleware returns an HTTP middleware for HTTP servers to check rate limits
func (rl *RateLimiterClient) Middleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// get request client identifier from request header

		allow, err := rl.Allow(r.Context(), r.RemoteAddr)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			return
		}

		if !allow {
			w.WriteHeader(http.StatusTooManyRequests)
			return
		}

		next.ServeHTTP(w, r)
	})
}
