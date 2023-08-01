package dto

type BookRequest struct {
	Title       string `json:"title" binding:"required"`
	Description string `json:"description" binding:"required,min=10"`
	Quantity    int    `json:"quantity" binding:"required,gt=0,numeric"`
	AuthorId    int    `json:"authorId" binding:"required"`
}

type BookResponse struct {
	ID          int             `json:"id"`
	Title       string          `json:"title"`
	Description string          `json:"description"`
	Quantity    int             `json:"quantity"`
	Cover       *string         `json:"cover,omitempty"`
	AuthorId    int             `json:"authorId,omitempty"`
	Author      *AuthorResponse `json:"author,omitempty"`
}
