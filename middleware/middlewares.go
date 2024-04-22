package middleware

import (
	"log"
	"strings"

	"github.com/Jamlie/vodka"
	"github.com/Jamlie/vodka/cors"
)

func Logger(c vodka.Context) {
	log.Println(c.Request().Method, c.Request().URL.Path)
}

func CORS(c vodka.Context) {
	c.Response().Header().Set("Access-Control-Allow-Origin", "*")
	c.Response().Header().Set("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE, OPTIONS")
	c.Response().Header().Set("Access-Control-Allow-Headers", "Content-Type, Authorization")
}

type CORSOptions struct {
	AllowedOrigins []string
	AllowedMethods []cors.Method
	AllowedHeaders []cors.AllowedHeader
}

func CORSWithConfig(options CORSOptions) vodka.HandlerFunc {
	return func(ctx vodka.Context) {
		origin := ctx.Request().Header.Get("Origin")
		if originIsValid(origin, options.AllowedOrigins) {
			ctx.Response().Header().Set("Access-Control-Allow-Origin", origin)
		}

		ctx.Response().
			Header().
			Set("Access-Control-Allow-Methods", joinMethods(options.AllowedMethods))

		ctx.Response().
			Header().
			Set("Access-Control-Allow-Headers", joinHeaders(options.AllowedHeaders))
	}
}

func joinMethods(methods []cors.Method) string {
	return strings.Join(methods, ", ")
}

func joinHeaders(headers []cors.AllowedHeader) string {
	return strings.Join(headers, ", ")
}

func originIsValid(origin string, allowedOrigins []string) bool {
	if len(allowedOrigins) == 0 {
		return true
	}

	for _, o := range allowedOrigins {
		if o == "*" || o == origin {
			return true
		}
	}

	return false
}
