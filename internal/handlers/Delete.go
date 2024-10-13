package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/google/uuid"
	"github.com/gorilla/mux"
	"github.com/hackman01/todoApi/models"
	"github.com/hackman01/todoApi/pkg/utils"
)



func (apiCfg *Config) DeleteItem(w http.ResponseWriter, r *http.Request) {
	type parameters struct {
		ID uuid.UUID `json:"id"`
	}
    
    id,err := uuid.Parse(mux.Vars(r)["itemId"])

	if err!=nil {
		utils.RespondWithError(w,500,fmt.Sprintf("Cant parse the Item id : %v",err))
		return
	}

    params := parameters{}

	decoder := json.NewDecoder(r.Body)

    err = decoder.Decode(&params)

	if err!=nil {
		utils.RespondWithError(w,500,fmt.Sprintf("Cant resolve the Json : %v",err))
		return
	}

	list,err := apiCfg.DB.DeleteItem(r.Context(), id)

	if err!=nil {
		utils.RespondWithError(w,500,fmt.Sprintf("couldnt Delete item %v",err))
		return
	}

	utils.RespondWithJson(w,200,models.ChangeFormat(list))


}