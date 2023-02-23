package main

import (
	"fmt"
	"net/http"
)

func hello(w http.ResponseWriter, _ *http.Request) {
	n, _ := fmt.Fprintf(w, "Hello world!")
	_ = fmt.Sprintf("Types of bytes: %d", n)
}

func about(w http.ResponseWriter, _ *http.Request) {
	n, _ := fmt.Fprintf(w, "About us!")
	_ = fmt.Sprintf("Types of bytes: %d", n)
}

func main() {

	http.HandleFunc("/", hello)
	http.HandleFunc("/about", about)

	if err := http.ListenAndServe(":8080", nil); err != nil {
		fmt.Println(err)
	}
}
