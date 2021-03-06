package main

import "net/http"
import (
	"views"
	"api"
	"middleware"
	"fmt"
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
	http.HandleFunc("/secret",middleware.WithMiddleware(views.SecretView,
		middleware.WithLogin()))
	http.HandleFunc("/login",views.Login)
	http.HandleFunc("/",views.IndexView)
	fmt.Printf("Server up and running on :8080")
	http.ListenAndServe(":8080", nil)
}
