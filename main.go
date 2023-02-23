package main

import (
	"fmt"
	"net/http"
	"stroy-service/handlers"
)

// main is the main function on our application
func main() {

	http.HandleFunc("/about", handlers.About)
	http.HandleFunc("/", handlers.Home)

	if err := http.ListenAndServe(":8080", nil); err != nil {
		fmt.Println(err)
	}
}
