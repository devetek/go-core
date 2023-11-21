package gqlclient

import (
	"context"
	"net/http"
)

// Middleware for common HTTP
func Middleware(client *Client) func(http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			ctx := context.WithValue(r.Context(), "gql", client)

			next.ServeHTTP(w, r.WithContext(ctx))
		})
	}
}
