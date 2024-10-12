package routes

import (
	"github.com/gorilla/mux"
	"github.com/hackman01/todoApi/internal/handlers"
)

func SetUpRouter() *mux.Router {
	
	router := mux.NewRouter()
	router.HandleFunc("/api/health",handlers.health).Methods("GET")
	return router
}