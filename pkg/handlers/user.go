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
//	w.WriteHeader(http.StatusOK)
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

	tocken := service.Login(&u)


	b, err := json.Marshal(&tocken)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	w.Write(b)

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

	b, err := json.Marshal(&user)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	w.Write(b)

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
}


func LogOutHandler (w http.ResponseWriter, r *http.Request){
	head := r.Header.Get("Authorization")

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
}