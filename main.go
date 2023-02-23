package main

import (
	"fmt"
	"html/template"
	"net/http"
)

// Home is the home page
func Home(w http.ResponseWriter, _ *http.Request) {
	renderTemplate(w, "home.page.tmpl")
}

// About is the about page
func About(w http.ResponseWriter, _ *http.Request) {
	renderTemplate(w, "about.page.tmpl")
}

func renderTemplate(w http.ResponseWriter, tmpl string) {
	parsedTemplete, _ := template.ParseFiles("./templates/" + tmpl)
	err := parsedTemplete.Execute(w, nil)
	if err != nil {
		fmt.Println("error parsing template:", err)
		return
	}
}

// main is the main function on our application
func main() {

	http.HandleFunc("/about", About)
	http.HandleFunc("/", Home)

	if err := http.ListenAndServe(":8080", nil); err != nil {
		fmt.Println(err)
	}
}
