package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"regexp"
)

type User struct {
	FirstName string
	LastName  string
	Email     string
	Password  string
}

type ValidationErr struct {
	FieldValue string
	ErrMassage string
}

func (u *User) Validate() []ValidationErr {
	errors := make([]ValidationErr, 0, 0)

	if len(u.FirstName) < 2 {
		errors = append(errors, ValidationErr{
			FieldValue: "FirstName",
			ErrMassage: "ErrFirstName(2)",
		})
	}
	if len(u.FirstName) > 64 {
		errors = append(errors, ValidationErr{
			FieldValue: "FirstName",
			ErrMassage: "ErrFirstName(64)",
		})
	}

	if len(u.LastName) < 2 {
		errors = append(errors, ValidationErr{
			FieldValue: "LastName",
			ErrMassage: "ErrLastName(2)",
		})
	}

	if len(u.LastName) > 64 {
		errors = append(errors, ValidationErr{
			FieldValue: "LastName",
			ErrMassage: "ErrLastName(64)",
		})
	}
	if len(u.Password) < 8 {
		errors = append(errors, ValidationErr{
			FieldValue: "Password",
			ErrMassage: "ErrPasswordName(8)",
		})
	}

	if len(u.Password) > 64 {
		errors = append(errors, ValidationErr{
			FieldValue: "Password",
			ErrMassage: "ErrPasswordName(64)",
		})
	}

	var emailRegex = regexp.MustCompile("^[a-zA-Z0-9.!#$%&'+\\/=?^_`{|}~-]+@[a-zA-Z0-9](?:[a-zA-Z0-9-]{0,61}[a-zA-Z0-9])?(?:\\.[a-zA-Z0-9](?:[a-zA-Z0-9-]{0,61}[a-zA-Z0-9])?)*$")
	if emailRegex.MatchString(u.Email) != true {

		errors = append(errors, ValidationErr{
			FieldValue: "Email",
			ErrMassage: "ErrEmail",
		})
	}

	return errors
}

//func (u *User) registrationValidation(){
// fmt.Println(u.FirstName, ":", u.errorFirstName())
// fmt.Println(u.LastName, ":", u.errorLastName())
// fmt.Println(u.Email, ":", u.errorEmail())
// fmt.Println(u.Password, ":", u.passwordError())
//}

func login(w http.ResponseWriter, r *http.Request) {
	u := User{}
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	defer r.Body.Close()

	data := []byte(body)

	err = json.Unmarshal(data, &u)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	validationErrors := u.Validate()

	if len(validationErrors) > 0 {
		w.WriteHeader(http.StatusBadRequest)

		b, err := json.Marshal(&validationErrors)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			return
		}

		w.Write(b)
		return
	}

	w.WriteHeader(http.StatusOK)

	//	u.registrationValidation()
	// 	fmt.Printf("struct:\n\t%#v\n\n", u)
	// fmt.Printf("%s", u)

	// if u.errorFirstName() && u.errorLastName() && u.errorEmail() && u.passwordError()  {
	//  fmt.Println(true)
	// } else {
	// fmt.Println(false)
	// }

}

func hello(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" {
		http.Error(w, "404 not found.", http.StatusNotFound)
		return
	}
	http.ServeFile(w, r, "./index.html")
}
func main() {
	http.HandleFunc("/", hello)
	http.HandleFunc("/login", login)
	http.Handle("/index.js", http.FileServer(http.Dir("./")))
	http.Handle("/1.jpg", http.FileServer(http.Dir("./")))
	http.Handle("/style.css", http.FileServer(http.Dir("./")))

	fmt.Printf("Starting server for testing HTTP POST...\n")
	if err := http.ListenAndServe(":8034", nil); err != nil {
		log.Fatal(err)
	}
}
