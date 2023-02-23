package main

import (
	"fmt"
	"net/http"
)

func hello(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello world!")
}

func about(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "About us!")
}

func main() {

	http.HandleFunc("/", hello)
	http.HandleFunc("/about", about)

	if err := http.ListenAndServe(":8080", nil); err != nil {
		fmt.Println(err)
	}
}
