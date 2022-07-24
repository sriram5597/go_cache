package router

import (
	mux "github.com/gorilla/mux"
	"net/http"

	handler "github.com/sriram5597/go_cache/src/api/handlers"
)


func GetRouter() *mux.Router {
	router := mux.NewRouter().StrictSlash(true)

	router.HandleFunc("/get-key", handler.GetKey)
	router.HandleFunc("/set-key", handler.SetKey).Methods(http.MethodPost)

	return router
}
