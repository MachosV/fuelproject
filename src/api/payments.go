package api

import (
	"net/http"
	"fmt"
	"storage"
)

func Payments(w http.ResponseWriter, r *http.Request){
	switch r.Method{
	case "POST":
		fmt.Fprint(w, storage.CreatePayment(r))
	case "GET":
		fmt.Fprint(w,storage.GetPayments())
	}
}