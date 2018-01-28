package main

import "net/http"
import (
	"views"
	"api"
)

func main() {
	http.HandleFunc("/api/drivers",api.Drivers)
	http.HandleFunc("/api/makepayment",views.MakePayment)
	http.ListenAndServe(":8080", nil)
}
