package views

import (
	"fmt"
	"net/http"
	"log"
)

func SecretView(w http.ResponseWriter, r *http.Request){
	log.Println("SecretView")
	fmt.Fprintf(w,"Successfuly got inside")
}

