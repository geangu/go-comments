package config

import (
	"net/http"

	"../controller"

	"github.com/gorilla/mux"
)

// InitRoutes initialize application routes
func InitRoutes() *mux.Router {
	controller := new(controller.CommentController)
	router := mux.NewRouter().StrictSlash(true)

	router.
		Methods("POST").
		Path("/comments/").
		Name("Create").
		Handler(http.HandlerFunc(controller.Create))

	router.
		Methods("DELETE").
		Path("/purchase/{id}/comments/").
		Name("Delete").
		Handler(http.HandlerFunc(controller.Delete))

	router.
		Methods("GET").
		Path("/purchase/{id}/comments/").
		Name("FindByPurchase").
		Handler(http.HandlerFunc(controller.FindByPurchase))

	router.
		Methods("GET").
		Path("/establishment/{id}/comments/").
		Name("FindByEstablishment").
		Handler(http.HandlerFunc(controller.FindByEstablishment))

	return router
}
