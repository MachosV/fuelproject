package api

import (
	"net/http"
	"fmt"
	"storage"
)

func Vehicles(w http.ResponseWriter, r *http.Request){
	switch r.Method{
	case "POST":
		fmt.Fprint(w, storage.CreateVehicle(r))
	case "GET":
		fmt.Fprint(w,storage.GetVehicles())
	}
}
