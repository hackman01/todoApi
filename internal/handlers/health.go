package handlers

import (
	"net/http"
	"github.com/hackman01/todoApi/pkg/utils"
)


func health(w http.ResponseWriter, r *http.Request){
	type status struct {
        Message string `json:"message"`
	}
    
	utils.respondWithJson(w,200,status{Message : "Server is Up!"})

}