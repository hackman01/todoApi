package main

import (
	"log"
	"net/http"
	"os"
	"github.com/hackman01/todoApi/internal/routes"
	"github.com/joho/godotenv"
)


func main(){

	godotenv.Load(".env")
	port := os.Getenv("PORT")
	
    router := routes.SetUpRouter()
    
	


	srv:=&http.Server{
		Handler: router,
		Addr: ":" + port,
	}

	log.Printf("Server starting on port %s",port)

    err := srv.ListenAndServe()

	if err!=nil {
		log.Fatal("Cant create the server",err)
	}

}