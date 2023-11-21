package render

import (
	"context"
	"net/http"

	"github.com/labstack/echo/v4"
)

// Middleware for common HTTP
func Middleware(engine *Engine) func(http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			ctx := context.WithValue(r.Context(), "renderer", engine.HTML(w))

			next.ServeHTTP(w, r.WithContext(ctx))
		})
	}
}

// Middleware for echo Framework
func EchoMiddleware(engine *Engine) echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			c.Set("renderer", engine.HTML(c.Response().Writer))

			return next(c)

		}
	}
}
