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

	//w.WriteHeader(http.StatusOK)

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

	token := databace.Login(&u)

	b, err := json.Marshal(&token)

	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	w.Write(b)

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

	b, err := json.Marshal(&user)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	w.Write(b)
	return

}

func LogOutHandler(w http.ResponseWriter, r *http.Request) {
	head := r.Header.Get("Authorization")

	databace.DeleteToken(head)

	fmt.Println(head)

	databace.DeleteToken(head)

	w.Write([]byte("ok"))

}
