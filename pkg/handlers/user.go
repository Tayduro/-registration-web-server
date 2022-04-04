package handlers

import (
	"encoding/json"
	"io/ioutil"
	"net/http"

	"github.com/Tayduro/registration-web-server/pkg/models"
	"github.com/Tayduro/registration-web-server/pkg/service"
)

func NewSignUpHandler(signupService *service.SignupService) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		u := &models.User{}
		body, err := ioutil.ReadAll(r.Body)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
		defer r.Body.Close()

		data := []byte(body)

		err = json.Unmarshal(data, u)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			return
		}

		validationErrors, err := signupService.CheckError(u)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			return
		}

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

		err = signupService.SignUp(u)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			return
		}

		w.Write([]byte("[]"))
	}
}

func NewSignInHandler(signupService *service.SignupService) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		u := &models.User{}
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

		token, err := signupService.SignIn(u)

		b, err := json.Marshal(&token)

		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
		w.Write(b)
	}
}


func NewGettingUserInformationHandler(signupService *service.SignupService) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
			token := r.Header.Get("Authorization")

			user,err := signupService.GettingUserInformationHandler(token)
			if err != nil {
				error := signupService.DeleteToken(token)
				if error != nil {
					w.WriteHeader(http.StatusInternalServerError)
					return
				}
				textErr := err.Error()
				b, err := json.Marshal(&textErr)
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
}

func NewLogOutHandler(signupService *service.SignupService) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		token := r.Header.Get("Authorization")
		err := signupService.DeleteToken(token)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			return
		}

		w.Write([]byte("ok"))

	}
}