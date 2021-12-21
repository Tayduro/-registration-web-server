package main

import (
	"fmt"
	"log"
	"net/http"
)
func hello(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" {
		http.Error(w, "404 not found.", http.StatusNotFound)
		return
	}

	switch r.Method {
	case "GET":
		 http.ServeFile(w, r, "index.html")
	case "POST":
		// Call ParseForm() to parse the raw query and update r.PostForm and r.Form.
		if err := r.ParseForm(); err != nil {
			fmt.Fprintf(w, "ParseForm() err: %v", err)
			return
		}
		// fmt.Fprintf(w, "Post from website! r.PostFrom = %v\n", r.PostForm)
		        Fname := r.FormValue("FirstName")
        		Lname := r.FormValue("LastName")
        		email := r.FormValue("Email")
        		password := r.FormValue("Password")
        		fmt.Printf("First name = %s\n", Fname)
        		fmt.Printf("Last Name = %s\n", Lname)
        		fmt.Printf("Email = %s\n", email)
        		fmt.Printf("password = %s\n", password)
		 http.ServeFile(w, r, "./confirmedRegistration.html")
	default:
		fmt.Fprintf(w, "Sorry, only GET and POST methods are supported.")
	}
}

func main() {
	http.HandleFunc("/", hello)
	http.Handle("/index.js", http.FileServer(http.Dir("./")))
    http.Handle("/1.jpg", http.FileServer(http.Dir("./")))
    http.Handle("/style.css", http.FileServer(http.Dir("./")))


	fmt.Printf("Starting server for testing HTTP POST...\n")
	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatal(err)
	}
}