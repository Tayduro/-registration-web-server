package main

import (
	"fmt"
	"github.com/Tayduro/registration-web-server/pkg/config"
	"github.com/Tayduro/registration-web-server/pkg/handlers"
	"github.com/Tayduro/registration-web-server/pkg/service"
	"github.com/jmoiron/sqlx"
	"log"
	"net/http"
)

func indexFileHandler(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path == "/" {
		http.ServeFile(w, r, "../assets/index.html")
		return
	}
	http.ServeFile(w, r, "../assets" + r.URL.Path)

}
func main() {
	cfg, err := config.ReadConfig("../cmd/signup-server/config.yaml")
	if err != nil {
		log.Fatalf("cannot read config file, err:%s", err.Error())

	}

	connection, err := sqlx.Connect("postgres", cfg.DBURL())
	if err != nil {
		log.Fatalf("cannot connection in database, err:%s", err.Error())

	}

	key := cfg.Key

	svc := service.NewSignupService(connection, key)


	http.HandleFunc("/", indexFileHandler)
	http.HandleFunc("/registration", handlers.NewSignUpHandler(svc))
	http.HandleFunc("/entrance", handlers.NewSignInHandler(svc))
	http.HandleFunc("/send-form", handlers.NewGettingUserInformationHandler(svc))
	http.HandleFunc("/log-out", handlers.NewLogOutHandler(svc))

	fmt.Printf("Starting server for testing HTTP POST....PORT:8034...\n")
	if err := http.ListenAndServe("0.0.0.0:8034", nil); err != nil {
		log.Fatal(err)
	}
}
