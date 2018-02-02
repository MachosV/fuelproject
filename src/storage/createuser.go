package storage

import (
	"golang.org/x/crypto/bcrypt"
	"fmt"
	"net/http"
)

func CreateUser(r *http.Request)(int64, error){
	var db = GetDb();
	var username = r.PostFormValue("username")
	var password = r.PostFormValue("password1")
	var email = r.PostFormValue("email")
	stmt, _ := db.Prepare("INSERT INTO accounts VALUES(?,?,?);")
	hash, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	res, err:= stmt.Exec(email,hash,username)
	_ = res
	if err != nil{
		fmt.Print(err)
		return 0,err
	}
	return 1,nil
}
