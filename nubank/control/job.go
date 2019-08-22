package control

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
	"nubank/database"
	"nubank/model"
	"nubank/validation"
)

func RegisterUser(w http.ResponseWriter, r *http.Request) {

	// retrieve body from request
	body := r.Body

	// declare use entity
	var user model.User

	// parsing io.ReadCLoser to slice of bytes []byte
	bytes, _ := ioutil.ReadAll(body)

	// parses json into user struct
	err := json.Unmarshal(bytes, &user)

	// checks if any error occurs in json parsing
	if err != nil {

		log.Printf("[WARN] problem parsing json body, because, %v\n", err)
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	// checks if struct is a valid one
	if err := validation.Validator.Struct(user); err != nil {

		log.Printf("[WARN] invalid user information, because, %v\n", err)
		w.WriteHeader(http.StatusPreconditionFailed)
		return
	}

	err = database.InsertJob(user)

	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.Write([]byte(`{"message": "success"}`))
}
