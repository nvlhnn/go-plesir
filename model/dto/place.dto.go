package dto

import "mime/multipart"

type PlaceCreateDTO struct {
	Name        string  `form:"name" json:"name" binding:"required"`
	Description string  `form:"description" json:"description" binding:"required"`
	Price       float32 `form:"price" json:"price" binding:"required,numeric"`
	UserID      uint    `form:"user_id" json:"user_id,omitempty" binding:"required,numeric"`
	WorkDays    []Work  `form:"work_days" json:"work_days"`
	Images      []multipart.FileHeader `form:"images" binding:"required"`
}

type PlaceUpdateDTO struct {
	Name        *string  `json:"name,omitempty"`
	Description *string  `json:"description,omitempty"`
	Price       *float32 `json:"price,omitempty"`
	UserID      *uint    `json:"user_id,omitempty"`
	WorkDays    []Work   `json:"work_days,omitempty"`
}

type PlaceResponseDTO struct {
	ID          uint    `json:"id"`
	Name        string  `json:"name"`
	Description string  `json:"description"`
	Price       float32 `json:"price"`
	Manager     Manager `json:"manager"`
	WorkDays    []Work  `json:"work_days"`
	Images		[]string `json:"images"`
	Slug		string	 `json:"slug"`
}

type Manager struct {
	Name  string `json:"name"`
	Email string `json:"email"`
}

type Work struct {
	Day   string `json:"day,omitempty"`
	DayID uint   `json:"day_id,omitempty"`
	Hour  Hour   `json:"hour"`
}

type Hour struct {
	Open  string `json:"open"`
	Close string `json:"close"`
}
