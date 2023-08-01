package entity

import (
	"git.garena.com/sea-labs-id/batch-05/gin-template/dto"
	"gorm.io/gorm"
)

type Author struct {
	ID         int    `json:"id,omitempty"`
	Name       string `json:"name,omitempty"`
	gorm.Model `json:"-"`
}

func (a *Author) ToDTO() *dto.AuthorResponse {
	return &dto.AuthorResponse{
		Id:   a.ID,
		Name: a.Name,
	}
}
