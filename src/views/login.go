package views

import (
	"net/http"
	"data"
	"html/template"
	"auth"
)


func GetLogin(w http.ResponseWriter, r *http.Request){
	var t *template.Template = data.GetTemplate("login.html")
	t.ExecuteTemplate(w, "base", nil)
}

func PostLogin(w http.ResponseWriter, r *http.Request){
	var cookie *http.Cookie = auth.DoLogin(r.PostFormValue("username"), r.PostFormValue("password"))
	if cookie == nil{
		http.Redirect(w, r, "/login", http.StatusMovedPermanently)
	}
	http.SetCookie(w,cookie)
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