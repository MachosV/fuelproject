package data

import (
	"html/template"
	"io/ioutil"
	"log"
)

var templates map[string]*template.Template

func init() {
	templates = make(map[string]*template.Template)
	files, err := ioutil.ReadDir("./templates")
	if err != nil {
		log.Fatal(err)
	}
	for _, file := range files {
		filename := file.Name()
		t := template.New(filename)
		t, err = t.ParseFiles("./templates/"+filename, "./templates/base.html")
		templates[filename] = t
	}
	for key := range templates {
		log.Println("Found", key, "template")
	}
	log.Println("Total templates", len(templates))
}

func GetTemplate(templateName string) *template.Template {
	return templates[templateName]
}
