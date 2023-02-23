package handlers

import (
	"net/http"
	"stroy-service/render"
)

// Home is the home page
func Home(w http.ResponseWriter, _ *http.Request) {
	render.RenderTemplate(w, "home.page.tmpl")
}

// About is the about page
func About(w http.ResponseWriter, _ *http.Request) {
	render.RenderTemplate(w, "about.page.tmpl")
}
