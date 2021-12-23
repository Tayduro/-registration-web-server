package main

import (
	"fmt"
	"log"
	"net/http"
	"io/ioutil"
)

func login(w http.ResponseWriter, r *http.Request) {
 b, err := ioutil.ReadAll(r.Body)
 if err != nil {
     w.WriteHeader(500)
     return
    }
   defer r.Body.Close()

   fmt.Printf("%s", string(b))

   http.ServeFile(w, r, "./index.html")
}

func hello(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" {
		http.Error(w, "404 not found.", http.StatusNotFound)
		return
	}
}
func main() {
	http.HandleFunc("/", hello)
	http.HandleFunc("/login", login)
	http.Handle("/index.js", http.FileServer(http.Dir("./")))
    http.Handle("/1.jpg", http.FileServer(http.Dir("./")))
    http.Handle("/style.css", http.FileServer(http.Dir("./")))

	fmt.Printf("Starting server for testing HTTP POST...\n")
	if err := http.ListenAndServe(":8033", nil); err != nil {
		log.Fatal(err)
	}
}