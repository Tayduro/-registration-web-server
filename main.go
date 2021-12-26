package main

import (
	"fmt"
	"log"
	"net/http"
	"io/ioutil"
	"encoding/json"
	"regexp"
)

type User struct {
	FirstName  string
	LastName string
	Email   string
	Password  string
}

func login(w http.ResponseWriter, r *http.Request) {
 b, err := ioutil.ReadAll(r.Body)
 if err != nil {
     w.WriteHeader(500)
     return
    }
   defer r.Body.Close()

   //fmt.Printf("%s", string(b))

   	data := []byte(b)

   	u := &User{}
   	json.Unmarshal(data, u)
    // 	fmt.Printf("struct:\n\t%#v\n\n", u)
  	// fmt.Printf("%s", u)

  fmt.Println(u.FirstName, ":", u.errorFirstName())
  fmt.Println(u.LastName, ":", u.errorLastName())
  fmt.Println(u.Email, ":", u.errorEmail())
  fmt.Println(u.Password, ":", u.passwordError())


   if u.errorFirstName() && u.errorLastName() && u.errorEmail() && u.passwordError()  {
    fmt.Println(true)
    } else {
    fmt.Println(false)
    }

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

func (u *User) errorFirstName() bool {
if len(u.FirstName) < 255 {
return true
}
return false

}

func (u *User) errorLastName()  bool {
 if len(u.LastName) < 255 {
  return true
 }
return false
}

func (u *User) passwordError() bool {
 if len(u.Password) < 8 || len(u.Password) > 64 {
 return false
 }
return true
}

func (u User) errorEmail() bool {
    var emailRegex = regexp.MustCompile("^[a-zA-Z0-9.!#$%&'+\\/=?^_`{|}~-]+@[a-zA-Z0-9](?:[a-zA-Z0-9-]{0,61}[a-zA-Z0-9])?(?:\\.[a-zA-Z0-9](?:[a-zA-Z0-9-]{0,61}[a-zA-Z0-9])?)*$")
    if len(u.Email) < 6 && len(u.Email) > 30 {
        return false
    }
    return emailRegex.MatchString(u.Email)
}
