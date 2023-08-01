package entity

import (
	"github.com/ericjovian/gin-template/dto"
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
