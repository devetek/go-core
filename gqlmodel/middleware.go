package gqlmodel

import (
	"context"
	"net/http"
)

// Middleware for common HTTP
func Middleware(model *Graphql) func(http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			ctx := context.WithValue(r.Context(), "model", model)

			next.ServeHTTP(w, r.WithContext(ctx))
		})
	}
}
