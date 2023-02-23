package main

import (
	"fmt"
	"net/http"
)

// Home is the home page
func Home(w http.ResponseWriter, _ *http.Request) {
	n, _ := fmt.Fprintf(w, "Hello world and people!")
	res := fmt.Sprintf("Types of bytes: %d", n)
	fmt.Println(res)
}

// About is the about page
func About(w http.ResponseWriter, _ *http.Request) {
	n, _ := fmt.Fprintf(w, "About us!")
	res := fmt.Sprintf("Types of bytes: %d", n)
	fmt.Println(res)
}

// Sum is the sum page
func Sum(w http.ResponseWriter, _ *http.Request) {
	sum := addValues(2, 4)
	_, _ = fmt.Fprintf(w, fmt.Sprintf("Add x and y = %d", sum))
}

// addValues is the add page
func addValues(x, y int) int {
	return x + y
}

// main is the main function on our application
func main() {

	http.HandleFunc("/about", About)
	http.HandleFunc("/sum", Sum)
	http.HandleFunc("/", Home)

	if err := http.ListenAndServe(":8080", nil); err != nil {
		fmt.Println(err)
	}
}
