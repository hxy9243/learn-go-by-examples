package ratelimit

import (
	"context"
	"encoding/json"
	"log"
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

type RateLimiterOptions struct {
	Limit  int64
	Window int64

	MasterName    string
	SentinelAddrs []string
	Password      string
}

type RateLimiterClient struct {
	redisConn   *redis.Client
	redisScript *redis.Script
	options     RateLimiterOptions
}

// NewRateLimiterClient returns a new rate limiter client that uses Redis as backend for rate limiting
func NewRateLimiterClient(options RateLimiterOptions) *RateLimiterClient {
	redisConn := redis.NewFailoverClient(&redis.FailoverOptions{
		MasterName:       options.MasterName,
		SentinelAddrs:    options.SentinelAddrs,
		SentinelPassword: options.Password,
		Password:         options.Password,
	})

	if options.Limit == 0 {
		options.Limit = 100
	}

	if options.Window == 0 {
		options.Window = 1
	}

	return &RateLimiterClient{
		redisConn:   redisConn,
		redisScript: redis.NewScript(luaScript),
		options:     options,
	}
}

// Allow checks if the request exceeds the limit
func (rl *RateLimiterClient) Allow(ctx context.Context, key string) (bool, error) {
	uniqueValue := uuid.New().String()

	val, err := rl.redisScript.Run(
		ctx, rl.redisConn, []string{key, uniqueValue}, rl.options.Limit, rl.options.Window, time.Now().UnixMilli(),
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
		clientId := r.Header.Get("X-Client-Id")
		if clientId == "" {
			log.Print("X-Client-Id header is required")

			w.WriteHeader(http.StatusBadRequest)
			json.NewEncoder(w).Encode(map[string]interface{}{"error": "X-Client-Id header is required", "status": http.StatusBadRequest})
			return
		}

		allow, err := rl.Allow(r.Context(), clientId)
		if err != nil {
			log.Printf("Error checking rate limiter: %s", err)

			w.WriteHeader(http.StatusInternalServerError)
			json.NewEncoder(w).Encode(map[string]interface{}{"error": "Internal error", "status": http.StatusInternalServerError})
			return
		}

		if !allow {
			log.Printf("Too many requests from: %s", clientId)

			w.WriteHeader(http.StatusTooManyRequests)
			json.NewEncoder(w).Encode(map[string]interface{}{"error": "Too many requests", "status": http.StatusTooManyRequests})
			return
		}

		next.ServeHTTP(w, r)
	})
}
