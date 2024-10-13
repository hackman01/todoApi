package routes

import (
	"database/sql"
	"log"
	"os"
    
	"github.com/gorilla/mux"
	"github.com/hackman01/todoApi/internal/database"
	"github.com/hackman01/todoApi/internal/handlers"
	_ "github.com/lib/pq"
	
)

func SetUpRouter() *mux.Router {
	

    dbUrl := os.Getenv("DB_URL")
	conn, err := sql.Open("postgres",dbUrl)
    
	if err!=nil {
		log.Fatal("Cant connect to database : ",err)
	}

	
	apiCfg := &handlers.Config{
		DB : database.New(conn),
	}
	
	router := mux.NewRouter()
	router.HandleFunc("/api/v1/health",handlers.Health).Methods("GET")
	router.HandleFunc("/api/v1/addItem",apiCfg.AddItem).Methods("POST")
	router.HandleFunc("/api/v1/readItems",apiCfg.ReadItems).Methods("GET")
	router.HandleFunc("/api/v1/updateItem",apiCfg.UpdateItem).Methods("PUT")
	router.HandleFunc("/api/v1/deleteItem/{itemId}",apiCfg.DeleteItem).Methods("DELETE")
	
	return router
}