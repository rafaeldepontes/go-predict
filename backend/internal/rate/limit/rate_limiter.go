package limit

import (
	"log"
	"net/http"

	"golang.org/x/time/rate"
)

type Middleware interface {
	GlobalRateLimit(h http.Handler) http.Handler
}

type middleware struct {
	rl *rate.Limiter
}

func NewMiddleware() Middleware {
	return &middleware{
		rl: rate.NewLimiter(1, 2),
	}
}

func (m *middleware) GlobalRateLimit(h http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if !m.rl.Allow() {
			log.Println("[WARN] Request limit reached.")
			http.Error(w, "Too many requests.", http.StatusTooManyRequests)
			return
		}
		h.ServeHTTP(w, r)
	})
}
