package entity

import (
	"github.com/ericjovian/gin-template/dto"
	"gorm.io/gorm"
)

type Book struct {
	ID          int     `json:"id"`
	Title       string  `json:"title"`
	Description string  `json:"description"`
	Quantity    int     `json:"quantity"`
	Cover       *string `json:"cover,omitempty"`
	AuthorId    int     `json:"authorId"`
	Author      *Author `json:"author,omitempty"`
	gorm.Model  `json:"-"`
}

func (b *Book) ToDTO() *dto.BookResponse {
	var author *dto.AuthorResponse
	if b.Author != nil {
		author = b.Author.ToDTO()
	}
	return &dto.BookResponse{
		ID:          b.ID,
		Title:       b.Title,
		Description: b.Description,
		Quantity:    b.Quantity,
		Cover:       b.Cover,
		AuthorId:    b.AuthorId,
		Author:      author,
	}
}

func (b *Book) ToStruct(req dto.BookRequest) {
	b.Title = req.Title
	b.Description = req.Description
	b.Quantity = req.Quantity
	b.AuthorId = req.AuthorId
}
