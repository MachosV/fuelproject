package views

import (
	"fmt"
	"net/http"
)

func SecretView(w http.ResponseWriter, r *http.Request){
	fmt.Fprintf(w,"Successfuly got inside")
}

