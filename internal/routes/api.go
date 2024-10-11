package routes

import (
	"github.com/gorilla/mux"
	"github.com/hackman01/todoApi/internal/handlers"
)

func SetUpRouter() *mux.Router {
	
	router := mux.NewRouter()
	router.Get("/api/health",handlers.health)
	return router
}