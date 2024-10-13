package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/google/uuid"
	"github.com/hackman01/todoApi/internal/database"
	"github.com/hackman01/todoApi/models"
	"github.com/hackman01/todoApi/pkg/utils"
)



func (apiCfg *Config) UpdateItem(w http.ResponseWriter, r *http.Request) {
	type parameters struct {
		Item string `json:"item"`
		ID   uuid.UUID `json:"id"`
	}

    params := parameters{}

	decoder := json.NewDecoder(r.Body)

    err := decoder.Decode(&params)

	if err!=nil {
		utils.RespondWithError(w,500,fmt.Sprintf("Cant resolve the Json : %v",err))
		return
	}

	list,err := apiCfg.DB.UpdateItem(r.Context(),database.UpdateItemParams{
		Item: params.Item,
		ID: params.ID,
	})

	if err!=nil {
		utils.RespondWithError(w,500,fmt.Sprintf("couldnt updtae the item %v",err))
		return
	}

	utils.RespondWithJson(w,200,models.ChangeFormat(list))


}