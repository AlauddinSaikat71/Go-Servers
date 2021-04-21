package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {

	//HandleFunc function to add route handlers to the web server
	http.HandleFunc("/hello", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprint(w, "Hello! Saikat")
	})

	fmt.Printf("Starting server at port 8080\n")
	//This method allows us to start the web server and specify the port to listen for incoming requests.
	//
	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatal(err)
	}
}
