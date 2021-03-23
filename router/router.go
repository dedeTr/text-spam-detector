package router

import (
	"github.com/gorilla/mux"

	"github.com/dedeTr/text-spam-detector.git/middleware"
)

func Router() *mux.Router {
	router := mux.NewRouter()

	router.HandleFunc("/add", middleware.AddText).Methods("POST", "OPTIONS")

	return router
}
