package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/Tayduro/registration-web-server/pkg/handlers"
)

func indexFileHandler(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" {
		http.Error(w, "404 not found.", http.StatusNotFound)
		return
	}
	http.ServeFile(w, r, "./assets/index.html")
}
func main() {
	http.HandleFunc("/", indexFileHandler)
	http.HandleFunc("/registration", handlers.RegistrationHandler)
<<<<<<< HEAD
	http.HandleFunc("/entrance", handlers.SignInHandler)
=======
	http.HandleFunc("/login", handlers.SignInHandler)
>>>>>>> bcf1b17145918f375e87bb4df57158a9ea8e60a3
	http.HandleFunc("/send-form", handlers.CheckInformation)
	http.HandleFunc("/log-out", handlers.LogOutHandler)
	http.Handle("/index.js", http.FileServer(http.Dir("./assets")))
	http.Handle("/entrance.js", http.FileServer(http.Dir("./assets")))
	http.Handle("/profile.js", http.FileServer(http.Dir("./assets")))
	http.Handle("/entrance.html", http.FileServer(http.Dir("./assets")))
	http.Handle("/profile.html", http.FileServer(http.Dir("./assets")))
	http.Handle("/confirmedRegistration.html", http.FileServer(http.Dir("./assets")))
	http.Handle("/background.jpg", http.FileServer(http.Dir("./assets")))
	http.Handle("/style.css", http.FileServer(http.Dir("./assets")))

	fmt.Printf("Starting server for testing HTTP POST....PORT:8034...\n")
	if err := http.ListenAndServe("127.0.0.1:8034", nil); err != nil {
		log.Fatal(err)
	}
}
