package views

import (
	"data"
	"html/template"
	"net/http"
)

func AboutView(w http.ResponseWriter, r *http.Request) {
	var t *template.Template = data.GetTemplate("index.html")
	t.ExecuteTemplate(w, "base", nil)
}
