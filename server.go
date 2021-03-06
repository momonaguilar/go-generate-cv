package main

import (
	"bytes"
	"fmt"
	"html/template"
	"log"
	"net/http"
	"os"
	"strconv"
	"time"
)

var tpl = template.Must(template.ParseFiles("./static/index.html"))

func main() {
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	mux := http.NewServeMux()

	fileServer := http.FileServer(http.Dir("./static"))
	mux.Handle("/", fileServer)
	mux.HandleFunc("/hello", helloHandler)
	mux.HandleFunc("/form", formHandler)

	fmt.Printf("Starting server at port 8080\n")

	if err := http.ListenAndServe(":"+port, mux); err != nil {
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

	Skills := []Skills{
		{
			Skill:  "HTML",
			Rating: 100,
		},
	}
	Socials := []Socials{
		{
			Social: r.FormValue("social"),
			Link:   r.FormValue("link"),
		},
	}

	yearFrom, err := strconv.Atoi(r.FormValue("yearFrom"))
	if err != nil {
		//fmt.Fprintf(rw, "WARN: Unable to parse start year, defaulting to current year, error: %v", err)
		yearFrom = int(time.Now().Year())
	}

	yearTo, err := strconv.Atoi(r.FormValue("yearTo"))
	if err != nil {
		//fmt.Fprintf(rw, "WARN: Unable to parse end year, defaulting to current year, error: %v", err)
		yearTo = int(time.Now().Year())
	}

	Experiences := []Experiences{
		{
			YearFrom:    yearFrom,
			YearTo:      yearTo,
			Company:     r.FormValue("company"),
			Role:        r.FormValue("role"),
			Description: r.FormValue("description"),
		},
	}

	profile := Profile{
		Name:            r.FormValue("name"),
		Position:        r.FormValue("position"),
		Address:         r.FormValue("address"),
		Phone:           r.FormValue("phone"),
		EmailAddress:    r.FormValue("emailAddress"),
		PersonalWebsite: r.FormValue("personalWebsite"),
		URLToImage:      "",

		Skills:      Skills,
		Socials:     Socials,
		Experiences: Experiences,
	}

	cvObject := cvObject{
		Profile:       profile,
		ColorTemplate: r.FormValue("colorTemplate"),
		PaperDimesion: r.FormValue("paperDimension"),
	}

	buf := &bytes.Buffer{}
	err = tpl.Execute(buf, cvObject)
	if err != nil {
		http.Error(rw, err.Error(), http.StatusInternalServerError)
		return
	}
	buf.WriteTo(rw)

	//fmt.Fprintf(rw, "INFO: POST request successful\n")
	//fmt.Println("%v", r.Body)
	name := r.FormValue("name")
	address := r.FormValue("address")

	fmt.Fprintf(rw, "Name: %s\n", name)
	fmt.Fprintf(rw, "Address: %s\n", address)

}
