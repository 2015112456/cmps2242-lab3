package main

import (
	"fmt"
	"math/rand"
	"net/http"
	"time"
)

// Task 1
func homeHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/plain; charset=utf-8")
	w.Write([]byte("Welcome to the Shapes API"))
}

func healthHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/plain; charset=utf-8")
	w.Write([]byte("Server is running"))
}

func aboutHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/plain; charset=utf-8")
	w.Write([]byte("Roger Zheng - CMPS2242 Lab #3"))
}

func timeHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/plain; charset=utf-8")
	w.Write([]byte(time.Now().Format("January 2, 2006 3:04:05 PM")))
}

// Task 2
func randomSumHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/plain; charset=utf-8")
	a := rand.Intn(1000)
	b := rand.Intn(1000)
	w.Write([]byte(fmt.Sprintf("Generated random numbers: %d and %d\n", a, b)))
	w.Write([]byte(fmt.Sprintf("Sum = %d", a+b)))
}

func main() {
	http.HandleFunc("/", homeHandler)
	http.HandleFunc("/health", healthHandler)
	http.HandleFunc("/about", aboutHandler)
	http.HandleFunc("/time", timeHandler)

	http.HandleFunc("/randomsum", randomSumHandler)

	fmt.Println("Server listening on port 4000...")
	http.ListenAndServe(":4000", nil)
}
