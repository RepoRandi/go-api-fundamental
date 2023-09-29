package book

import "time"

type BookRequest struct {
	Title       string    `json:"title" binding:"required"`
	Description string    `json:"description" binding:"required"`
	Price       int       `json:"price" binding:"required,number"`
	Rating      int       `json:"rating" binding:"required,number"`
	Discount    int       `json:"discount" binding:"required,number"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
}
