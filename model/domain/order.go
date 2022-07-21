package domain

import (
	"time"

	"gorm.io/gorm"
)

type Order struct {
	gorm.Model
	InvoiceNumber		string `gorm:"not null" json:"invoice_number"`
	XenditID	string  `gorm:"" json:"xendit_id"`
	Price      float32	`gorm:"type:decimal(18,2);not null" json:"price"`
	Quantity   uint 	`gorm:"not null" json:"quantity"`
	AdminFee	float32	`gorm:"type:decimal(18,2)" json:"admin_fee"`
	TotalOrder float32 	`gorm:"type:decimal(18,2);not null" json:"total_order"`
	Status     string 		`gorm:"default:PENDING" json:"status"`	
	Date    time.Time 	`gorm:"not null" json:"date"`
	PaymentURL string	`gorm:"" json:"payment_url"`

	UserID     uint 	`gorm:"not null" json:"-"`
	User       User 	`gorm:"foreignKey:UserID;constraint:onUpdate:CASCADE,onDelete:CASCADE" json:"user"`

	PlaceID     uint 	`gorm:"not null" json:"-"`
	Place       Place 	`gorm:"foreignKey:PlaceID;constraint:onUpdate:CASCADE,onDelete:CASCADE" json:"place"`
}




