module gihub.com/hackman01/todoApi

replace github.com/hackman01/todoApi/handlers v0.0.0 => ./handlers

go 1.23.1

require (
	github.com/go-chi/chi v1.5.5
	github.com/joho/godotenv v1.5.1
)

require github.com/gorilla/mux v1.8.1 // indirect
