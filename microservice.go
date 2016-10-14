package main

import (
	"fmt"
	"net/http"
	"encoding/json"
)

const (
	AUTHOR = "hocklo"
)

type Message struct {
	Text	string
	Version	string
	Date	string
	Author	string
}

/**
 * Take the first url parameter and say Welcome, "value of argument".
 * e.g : Welcome, hocklo!
 */
func handler(w http.ResponseWriter, r *http.Request) {
	// localhost:8080/{argument} 
	fmt.Fprintf(w, "Welcome, %s!", r.URL.Path[1:])
}
/**
 * Output a message about this go program.
 */
func about(w http.ResponseWriter, r *http.Request) {
	// Save the message inside m
	m:= Message{"Welcome to the Hocklo API","0.1,","2016/10/12", AUTHOR}
	// Marshal "m" inside "b" if some problem occurs the value are writed inside "err".
	b, err := json.Marshal(m)
	// Check the value of "err" if it isn't "nil" output a panic error.
	if err != nil {
		panic(err)
	}
	// Write the value of "b" at ResponseWriter
	w.Write(b)
}

/**
 * Main function to handle all mappings
 */
func main() {
	http.HandleFunc("/", handler)
	http.HandleFunc("/about", about)
	http.ListenAndServe(":8080", nil)
}
