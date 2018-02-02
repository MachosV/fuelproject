package views

import (
	"net/http"
	"fmt"
	"storage"
)

func CreateUserView(w http.ResponseWriter, r *http.Request) {
	if r.PostFormValue("password1") != r.PostFormValue("password2"){
		fmt.Fprintf(w,"Password mismatch")
		return
	}
	_, err := storage.CreateUser(r)
	if err !=  nil{
		fmt.Print(err)
		fmt.Fprintf(w,"User creation error")
		return
	}
	fmt.Fprint(w,"User created")
}
