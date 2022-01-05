package handlers

import (
	"encoding/json"
	"github.com/Tayduro/registration-web-server/pkg/models"
	"io/ioutil"
	"net/http"
)

func SignupHandler(w http.ResponseWriter, r *http.Request) {
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

	err = servise.Signup(&u)

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
	w.Write([]byte("[{}]"))
	w.WriteHeader(http.StatusOK)

}