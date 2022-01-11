package main

import (
	"fmt"
	"log"
	"net/http"
	// "log"
	// "net/http"
)

func main() {
	fileServer := http.FileServer(http.Dir("./static"))
	http.Handle("/", fileServer)
	http.HandleFunc("/hello", helloHandler)
	http.HandleFunc("/form", formHandler)

	fmt.Printf("Starting server at port 8080\n")

	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatal(err)
	}
}

func helloHandler(rw http.ResponseWriter, r *http.Request) {
	fmt.Printf("%v", r.URL.Path)
	if r.URL.Path != "/hello" {
		http.Error(rw, "ERROR: 404 not found.", http.StatusNotFound)
		return
	}

	if r.Method != "GET" {
		http.Error(rw, "ERROR: Method is not supported", http.StatusNotAcceptable)
		return
	}

	fmt.Fprintf(rw, "INFO: Hello handler activated\n")
}

func formHandler(rw http.ResponseWriter, r *http.Request) {
	if err := r.ParseForm(); err != nil {
		fmt.Fprintf(rw, "ERROR: ParseForm() error: %v", err)
		return
	}

	fmt.Fprintf(rw, "INFO: POST request successful\n")

	name := r.FormValue("name")
	address := r.FormValue("address")

	fmt.Fprintf(rw, "Name: %s\n", name)
	fmt.Fprintf(rw, "Address: %s\n", address)

}
