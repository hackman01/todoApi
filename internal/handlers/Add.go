package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"github.com/google/uuid"
	"github.com/hackman01/todoApi/internal/database"
	"github.com/hackman01/todoApi/models"

	"github.com/hackman01/todoApi/pkg/utils"
)



func (apiCfg *Config) AddItem(w http.ResponseWriter, r *http.Request) {
	type parameters struct {
		Item string `json:"item"`
	}

    params := parameters{}

	decoder := json.NewDecoder(r.Body)

    err := decoder.Decode(&params)

	if err!=nil {
		utils.RespondWithError(w,500,fmt.Sprintf("Cant resolve the Json : %v",err))
		return
	}

	list,err := apiCfg.DB.AddItem(r.Context(),database.AddItemParams{ 
		ID: uuid.New(),
		CreatedAt: time.Now().Local().UTC(),
		UpdatedAt: time.Now().Local().UTC(),
		Item: params.Item,
	})

	if err!=nil {
		utils.RespondWithError(w,500,fmt.Sprintf("couldnt add the item %v",err))
		return
	}

	utils.RespondWithJson(w,201,models.ChangeFormat(list))


}