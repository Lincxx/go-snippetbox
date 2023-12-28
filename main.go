package main

import (
	"log"
	"net/http"
)

// Define a home handler function which writes a byte slice containing
// Home from Snipperbox as the response body
func home(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Hello from Snipptebox"))
}

func main() {
	// Use the http.NewServeMux() function to initialize a new servemux, then
	// register the home function as the handler for the "/" URL pattern.
	mux := http.NewServeMux()
	mux.HandleFunc("/", home)

	// Print a log a message to say that the server is starting.
	log.Print("starting server on :4000")

	// Use the http.ListenAndServe() function to start a new web server. We pass in
	// two parameters: the TCP network address to listen on (in this case ":4000")
	// and the servemux we just created. If http.ListenAndServe() returns an error
	// we use the log.Fatal() function to log the error message and exit. Note
	// that any error returned by http.ListenAndServe() is always non-nil.
	err := http:http.ListenAndServe(":4000", mux)
	log.Fatal(err)

}