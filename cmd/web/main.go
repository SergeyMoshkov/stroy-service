package main

import (
	"fmt"
	"log"
	"net/http"
	"stroy-service/pkg/config"
	"stroy-service/pkg/handlers"
	"stroy-service/pkg/render"
)

// main is the main function on our application
func main() {

	var app config.AppConfig

	tc, err := render.CreateTemplateCache()
	if err != nil {
		log.Fatal("cannot create template cache")
	}

	app.TemplateCache = tc

	http.FileServer(http.Dir("/static"))
	http.HandleFunc("/about", handlers.About)
	http.HandleFunc("/", handlers.Home)

	fmt.Println("Starting application on port: 8080")
	if err := http.ListenAndServe(":8080", nil); err != nil {
		fmt.Println(err)
	}
}
