package views

import (
	"data"
	"html/template"
	"net/http"
)

func AboutView(w http.ResponseWriter, r *http.Request) {
	var t *template.Template = data.GetTemplate("dashboard.html")
	t.ExecuteTemplate(w, "base", nil)
}
