package main

import (
	"fmt"
	"github.com/Tayduro/registration-web-server/pkg/handlers"
	"log"
	"net/http"
)

func indexFileHendler(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" {
		http.Error(w, "404 not found.", http.StatusNotFound)
		return
	}
	http.ServeFile(w, r, "./index.html")
}
func main() {
	http.HandleFunc("/", indexFileHendler)
	http.HandleFunc("/login", handlers.SignupHandler)
	http.Handle("/index.js", http.FileServer(http.Dir("./")))
	http.Handle("/confirmedRegistration.html", http.FileServer(http.Dir("./")))
	http.Handle("/1.jpg", http.FileServer(http.Dir("./")))
	http.Handle("/style.css", http.FileServer(http.Dir("./")))
	http.Handle("/style2.css", http.FileServer(http.Dir("./")))

	fmt.Printf("Starting server for testing HTTP POST....PORT:8034...\n")
	if err := http.ListenAndServe(":8034", nil); err != nil {
		log.Fatal(err)
	}
}
