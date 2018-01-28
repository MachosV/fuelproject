package api

import (
	"net/http"
	"fmt"
	"storage"
)

func Drivers(w http.ResponseWriter, r *http.Request){
	switch r.Method{
	case "POST":
		fmt.Fprint(w, storage.CreateDriver(r))
	case "GET":
		fmt.Fprint(w,storage.GetDrivers())
	}
}

