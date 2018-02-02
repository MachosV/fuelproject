package views

import (
	"data"
	"html/template"
	"net/http"
)

func SignupView(w http.ResponseWriter, r *http.Request) {
	var t *template.Template = data.GetTemplate("signup.html")
	t.ExecuteTemplate(w, "base", nil)
}

