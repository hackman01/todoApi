package handlers

import (
	"fmt"
	"net/http"

	"github.com/hackman01/todoApi/models"
	"github.com/hackman01/todoApi/pkg/utils"
)



func (apiCfg *Config) ReadItems(w http.ResponseWriter, r *http.Request) {


    items, err := apiCfg.DB.GetItems(r.Context())

	if err!=nil {
		utils.RespondWithError(w,500,fmt.Sprintf("Cant fetch from db : %v",err))
		return
	}

    
	utils.RespondWithJson(w,200,models.ChangeFormatList(items))


}