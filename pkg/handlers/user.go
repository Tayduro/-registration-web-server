package handlers

import (
	"encoding/json"
	"io/ioutil"
	"net/http"

	"github.com/Tayduro/registration-web-server/pkg/databace"
	"github.com/Tayduro/registration-web-server/pkg/models"
	"github.com/Tayduro/registration-web-server/pkg/service"
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
       databace.DataB(me)
	 w.Write([]byte("[]"))
//	w.WriteHeader(http.StatusOK)
}