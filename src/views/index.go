package views

import (
	"html/template"
	"net/http"
	"data"
)

func IndexView(w http.ResponseWriter, r *http.Request) {
	var t *template.Template = data.GetTemplate("index.html")
	t.ExecuteTemplate(w, "base", nil)
}
