package main

import "net/http"
import (
	"views"
	"api"
	"middleware"
)

func main() {
	http.HandleFunc("/api/payments",api.Payments)
	http.HandleFunc("/api/vehicles",api.Vehicles)
	http.HandleFunc("/api/drivers",api.Drivers)
	http.HandleFunc("/api/makepayment",views.MakePayment)
	http.HandleFunc("/dashboard",views.DashboardView)
	http.HandleFunc("/about",views.AboutView)
	http.HandleFunc("/createuser",views.CreateUserView)
	http.HandleFunc("/signup",views.SignupView)
	http.Handle("/secret",middleware.WithMiddleware(views.SecretView,
	middleware.WithLogin()))
	http.HandleFunc("/login",views.Login)
	http.HandleFunc("/",views.IndexView)
	http.ListenAndServe(":8080", nil)
}
