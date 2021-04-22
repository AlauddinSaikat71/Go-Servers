package main

import (
	"fmt"
	"log"
	"net/http"
)

func helloHandler(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/hello" {
		http.Error(w, "404 not found", http.StatusNotFound)
		return
	}
	if r.Method != "GET" {
		http.Error(w, "Method is not supported", http.StatusNotFound)
		return
	}
	fmt.Fprint(w, "Hello!")
}

func main() {

	//HandleFunc function to add route handlers to the web server
	http.HandleFunc("/hello", helloHandler)

	fmt.Printf("Starting server at port 8080\n")
	//This method allows us to start the web server and specify the port to listen for incoming requests.
	//
	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatal(err)
	}
}
