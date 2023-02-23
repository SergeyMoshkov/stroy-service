package render

import (
	"fmt"
	"html/template"
	"net/http"
)

// RenderTemplate is function for render template
func RenderTemplate(w http.ResponseWriter, tmpl string) {
	parsedTemplate, _ := template.ParseFiles("./static/templates/" + tmpl)
	err := parsedTemplate.Execute(w, nil)
	if err != nil {
		fmt.Println("error parsing template:", err)
		return
	}
}
