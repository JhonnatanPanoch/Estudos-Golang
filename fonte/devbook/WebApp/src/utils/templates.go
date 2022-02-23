package utils

import (
	"html/template"
	"net/http"
)

var templates *template.Template

func CarregarTemplates() {
	templates = template.Must(template.ParseGlob("src/views/*.html"))
}

func ExecutarTemplate(rw http.ResponseWriter, template string, dados interface{}) {
	templates.ExecuteTemplate(rw, template, dados)
}
