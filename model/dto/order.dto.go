package dto

import (
	"time"
)

type OrderCreate struct {
	// Price    float32 `json:"price" binding:"required,numeric"`
	Date   time.Time `json:"date" binding:"required"` 
	Quantity uint `json:"quantity" binding:"required,numeric"`
	PlaceID  uint `json:"place_id,omitempty" binding:"required,numeric"`
	UserID   uint `json:"-"`
}

type OrderResponse struct {
	InvoiceNumber      string `json:"invoice_number"`
	Price      float32 `json:"price"`
	Quantity   uint    `json:"quantity"`
	Place      PlaceResponseDTO  `json:"place"`
	User       Manager `json:"user"`
	Date 		time.Time `json:"date"`
	TotalOrder float32 `json:"total_order"`
	Status     string  `json:"status"`
	PaymentURL string	`json:"payment_url"`
	AdminFee	float32 `json:"admin_fee"`
}


type XenditRequest struct {
	ExsternalId string `json:"external_id"`
	Status string	`json:"status"`
}

// type Date struct {
// 	// Date time.Time
// }