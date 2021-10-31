package router

import (
	"server/middleware"

	"github.com/gorilla/mux"
)

func Router() *mux.Router {
	router := mux.NewRouter()

	router.HandleFunc("/api/paraphrase/create", middleware.CreateParaphrase).Methods("POST", "OPTIONS")

	return router
}
