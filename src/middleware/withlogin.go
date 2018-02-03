package middleware

import (
	"net/http"
	"auth"
)

func WithLogin() Middleware{
	return func(h http.HandlerFunc) http.HandlerFunc {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request){
			cookie, err := r.Cookie("gses")
			if err != nil{
				http.Redirect(w, r, "/login",http.StatusMovedPermanently)
				return
			}
			if auth.CheckAuth(cookie.Value){
				h.ServeHTTP(w, r)
			}else {
				http.Redirect(w, r, "/login",http.StatusMovedPermanently)
			}
		})
	}
}

