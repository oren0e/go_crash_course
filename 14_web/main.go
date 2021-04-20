package main

import (
	"fmt"
	"net/http"
)

func index(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "<h1>Hello World</h1>")
}

func about(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "<h1>About</h1>")
}

func main() {
	http.HandleFunc("/", index) // a route and a function to deal with that route (index) (like in Python flask)
	http.HandleFunc("/about", about)
	fmt.Println("Server starting...")
	http.ListenAndServe(":3000", nil) // localhost address with port 3000 and no handlers, it should just run a simple web server
}
