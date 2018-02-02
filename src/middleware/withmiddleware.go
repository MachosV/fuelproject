package middleware

import "net/http"

type Middleware func(http.HandlerFunc) http.HandlerFunc


func WithMiddleware(h http.HandlerFunc, middlewares ...Middleware) http.HandlerFunc{
	for _, middleware := range middlewares{
		h = middleware(h)
	}
	return h
}