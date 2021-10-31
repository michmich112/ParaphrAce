package router

import (
	"server/context"
	"server/middleware"

	"github.com/gorilla/mux"
)

func Router(aCtx context.AppContext) *mux.Router {
	router := mux.NewRouter()

	router.HandleFunc("/api/user/create", middleware.CreateUser(aCtx)).Methods("POST", "OPTIONS")
	router.HandleFunc("/api/paraphrase/create", middleware.CreateParaphrase).Methods("POST", "OPTIONS")

	return router
}
