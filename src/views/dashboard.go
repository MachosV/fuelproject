package views

import (
	"html/template"
	"net/http"
	"data"
)

func DashboardView(w http.ResponseWriter, r *http.Request) {
	var t *template.Template = data.GetTemplate("dashboard.html")
	t.ExecuteTemplate(w, "base", nil)
}
