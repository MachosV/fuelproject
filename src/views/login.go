package views

import (
	"net/http"
	"data"
	"html/template"
	"auth"
	"time"
)


func GetLogin(w http.ResponseWriter, r *http.Request){
	var t *template.Template = data.GetTemplate("login.html")
	t.ExecuteTemplate(w, "base", nil)
}

func PostLogin(w http.ResponseWriter, r *http.Request){
	var sessionid string = auth.DoLogin(r.PostFormValue("username"), r.PostFormValue("password"))
	if sessionid == ""{
		http.Redirect(w, r, "/login", http.StatusSeeOther)
	}
	cookie := http.Cookie{Name:"gses", Value:sessionid, Expires:time.Now().Add(2*time.Minute)}
	http.SetCookie(w,&cookie)
	http.Redirect(w,r,"/secret",http.StatusMovedPermanently)
}

func Login(w http.ResponseWriter, r *http.Request){
	switch r.Method {
		case "GET":
			GetLogin(w,r)
		case "POST":
			PostLogin(w,r)
	default:
		GetLogin(w,r)
	}
}