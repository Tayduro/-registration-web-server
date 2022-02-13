package handlers

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/Tayduro/registration-web-server/pkg/databace"
	"github.com/Tayduro/registration-web-server/pkg/models"
	"github.com/Tayduro/registration-web-server/pkg/service"
)

func RegistrationHandler(w http.ResponseWriter, r *http.Request) {
	u := models.User{}
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

	validationErrors := service.Signup(&u)

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

	me := &u
       databace.DataBaseRegistration(me)
	 w.Write([]byte("[]"))
<<<<<<< HEAD
	//w.WriteHeader(http.StatusOK)
=======
//	w.WriteHeader(http.StatusOK)
>>>>>>> bcf1b17145918f375e87bb4df57158a9ea8e60a3
}

func SignInHandler(w http.ResponseWriter, r *http.Request) {
	u := models.User{}
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

<<<<<<< HEAD
	token := databace.Login(&u)


	b, err := json.Marshal(&token)
=======
	tocken := service.Login(&u)


	b, err := json.Marshal(&tocken)
>>>>>>> bcf1b17145918f375e87bb4df57158a9ea8e60a3
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	w.Write(b)
<<<<<<< HEAD
	//w.WriteHeader(http.StatusOK)
}

func CheckInformation(w http.ResponseWriter, r *http.Request) {
	head := r.Header.Get("Authorization")

	user := databace.GettingUserData(head)
	if len(user) == 0 {
       tokenError := "Token is expired"
		b, err := json.Marshal(&tokenError)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
		w.Write(b)
		return
	}
=======

	//me := &u
	//service.Login(me)

	//	w.WriteHeader(http.StatusOK)
}

func CheckInformation(w http.ResponseWriter, r *http.Request) {
	//u := models.User{}
	//body, err := ioutil.ReadAll(r.Body)
	//if err != nil {
	//	w.WriteHeader(http.StatusInternalServerError)
	//	return
	//}
	//defer r.Body.Close()

	head := r.Header.Get("Authorization")

	user := databace.GettingUserData(head)

	fmt.Println(head)
>>>>>>> bcf1b17145918f375e87bb4df57158a9ea8e60a3

	b, err := json.Marshal(&user)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	w.Write(b)

<<<<<<< HEAD
=======
	//data := []byte(body)

	//err = json.Unmarshal(data, &u)
	//if err != nil {
	//	w.WriteHeader(http.StatusInternalServerError)
	//	return
	//}
	//me := &u

	//databace.DataBaseRegistration(me)
	//tocken := service.Login(&u)


	//b, err := json.Marshal(&tocken)
	//if err != nil {
	//	w.WriteHeader(http.StatusInternalServerError)
	//	return
	//}
	//w.Write(b)

	//w.Write([]byte("[]"))

	//me := &u
	//service.Login(me)

	//	w.WriteHeader(http.StatusOK)
>>>>>>> bcf1b17145918f375e87bb4df57158a9ea8e60a3
}


func LogOutHandler (w http.ResponseWriter, r *http.Request){
	head := r.Header.Get("Authorization")

<<<<<<< HEAD
	databace.DeleteToken(head)

=======
	//user := databace.GettingUserData(head)

	fmt.Println(head)

	databace.DeleteToken(head)

	w.Write([]byte("ok"))

	//b, err := json.Marshal(&user)
	//if err != nil {
	//	w.WriteHeader(http.StatusInternalServerError)
	//	return
	//}
	//w.Write(b)
>>>>>>> bcf1b17145918f375e87bb4df57158a9ea8e60a3
}