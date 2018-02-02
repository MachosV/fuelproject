package middleware

import (
	"net/http"
	"log"
)

func Log() Middleware{
	return func(h http.HandlerFunc) http.HandlerFunc {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request){
			log.Println("Before")
			defer log.Println("After")
			h.ServeHTTP(w,r)
		})
	}
}
