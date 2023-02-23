package main

import (
	"fmt"
	"net/http"
)

func Home(w http.ResponseWriter, _ *http.Request) {
	n, _ := fmt.Fprintf(w, "Hello world!")
	res := fmt.Sprintf("Types of bytes: %d", n)
	fmt.Println(res)
}

func About(w http.ResponseWriter, _ *http.Request) {
	n, _ := fmt.Fprintf(w, "About us!")
	res := fmt.Sprintf("Types of bytes: %d", n)
	fmt.Println(res)
}

func main() {

	http.HandleFunc("/", Home)
	http.HandleFunc("/about", About)

	if err := http.ListenAndServe(":8080", nil); err != nil {
		fmt.Println(err)
	}
}
