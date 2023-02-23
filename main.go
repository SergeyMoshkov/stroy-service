package main

import (
	"errors"
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

// Divide is the divide page
func Divide(w http.ResponseWriter, _ *http.Request) {
	divide, err := divideValues(3, 0)
	if err != nil {
		_, _ = fmt.Fprintf(w, "На ноль делить нельзя! %s\n", err)
	}
	_, _ = fmt.Fprintf(w, fmt.Sprintf("devide x and y = %f", divide))
}

// divideValues is function divide arguments
func divideValues(x, y float32) (float32, error) {
	if y == 0 {
		err := errors.New("cannot divide by zero")
		return 0, err
	}
	result := x / y
	return result, nil
}

// addValues is the add page
func addValues(x, y int) int {
	return x + y
}

// main is the main function on our application
func main() {

	http.HandleFunc("/about", About)
	http.HandleFunc("/sum", Sum)
	http.HandleFunc("/divide", Divide)
	http.HandleFunc("/", Home)

	if err := http.ListenAndServe(":8080", nil); err != nil {
		fmt.Println(err)
	}
}
