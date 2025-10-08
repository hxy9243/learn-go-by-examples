package main

import (
	"log"
	"net/http"
	"os"
	"strconv"

	"github.com/hxy9243/learn-go-by-examples/rate-limiter/ratelimit"
)

type ServerConfig struct {
	RedisAddr       string
	RateLimit       int64
	RateLimitWindow int64
}

// StartNewExampleServer creates an example server that uses rate limiter
// as a middleware for handling requests
func StartNewExampleServer(config ServerConfig) error {
	mux := http.NewServeMux()

	mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Hello World"))
	})

	// setup rate limiter
	ratelimiter := ratelimit.NewRateLimiterClient(
		config.RedisAddr,
		ratelimit.RateLimiterOptions{
			Limit:  config.RateLimit,
			Window: config.RateLimitWindow,
		},
	)

	handler := ratelimiter.Middleware(mux)
	return http.ListenAndServe(":8080", handler)
}

func getEnv(key, fallback string) string {
	if value, ok := os.LookupEnv(key); ok {
		return value
	}
	return fallback
}

func getEnvAsInt(name string, fallback int) int {
	valueStr := getEnv(name, "")
	if value, err := strconv.Atoi(valueStr); err == nil {
		return value
	}
	return fallback
}

func main() {
	redis_addr := getEnv("REDIS_ADDR", "localhost:6379")
	rate_limit := getEnvAsInt("RATE_LIMIT", 100)
	rate_limit_window := getEnvAsInt("RATE_LIMIT_WINDOW", 1)

	log.Fatal(StartNewExampleServer(ServerConfig{
		RedisAddr:       redis_addr,
		RateLimit:       int64(rate_limit),
		RateLimitWindow: int64(rate_limit_window),
	}))
}
