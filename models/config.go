package models

import (
	"time"

	"github.com/google/uuid"
	"github.com/hackman01/todoApi/internal/database"
)

type List struct {
	ID        uuid.UUID `json:"id"`
	Item      string    `json:"item"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

func ChangeFormat(list database.List) List {
	return List{
		ID : list.ID,
		Item : list.Item,
		CreatedAt: list.CreatedAt,
		UpdatedAt: list.UpdatedAt,
	}
}

func ChangeFormatList(list []database.List) []List {

	var res []List

	for _,lst := range list{
		res = append(res, ChangeFormat(lst))
	}
	return res
	
}
