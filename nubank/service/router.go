package service

import (
	"net/http"
	"nubank/configs"
	"nubank/control"

	"github.com/gorilla/mux"
	"github.com/rs/cors"
)

func createHandler() (handler *mux.Router) {

	// creats router
	handler = mux.NewRouter()

	// associate register user route
	handler.HandleFunc(configs.JOB_PATH, control.RegisterUser).Methods(http.MethodPost)

	// login user
	// handler.HandleFunc(configs.LOGIN_PATH, control.LoginUser).Methods(http.MethodPost)

	//
	//handler.HandleFunc(configs.USER_PATH, control.UpdateUser).Methods(http.MethodPut)

	// returns handle
	return
}

func createCORS() *cors.Cors {

	return cors.New(cors.Options{
		AllowedOrigins: []string{"*"},
		AllowedMethods: []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowedHeaders: []string{"Content-Type", "Accept", "Content-Length", "Accept-Encoding", "Authorization", "X-CSRF-Token"},
	})
}
