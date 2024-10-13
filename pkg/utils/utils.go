package utils

import (
	"encoding/json"
	"log"
	"net/http"
)


func RespondWithJson(w http.ResponseWriter,code int,payload interface{}){
	dat,err := json.Marshal(payload)

	if err!=nil {
		log.Println("Couldnt parse json")
		w.WriteHeader(500)
		return
	}

	w.WriteHeader(code)
	w.Header().Add("Content-Type", "application/json")
	w.Write(dat)
}

func RespondWithError(w http.ResponseWriter,code int,msg string){
	if code>499 {
		log.Println("Server Error",msg)
	}

	type errorResp struct {
		Error string `json:"error"`
	}

	RespondWithJson(w,code,errorResp{Error: msg})
}