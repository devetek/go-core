package render

import (
	"context"
	"net/http"
)

// Middleware for common HTTP
func Middleware(engine *Engine) func(http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			ctx := context.WithValue(r.Context(), CTX_KEY, engine.HTML(w))

			next.ServeHTTP(w, r.WithContext(ctx))
		})
	}
}
