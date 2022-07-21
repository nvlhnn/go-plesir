package helper

import (
	"github.com/nvlhnn/go-plesir/formatter"
	"github.com/nvlhnn/go-plesir/model/domain"
	"github.com/nvlhnn/go-plesir/model/dto"
)

func ToOrderResponse(order domain.Order) dto.OrderResponse {

	// var status string

	place := formatter.ToPlaceResponse(order.Place)

	// switch order.Status {
	// case 0:
	// 	status = "PENDING"
	// case 1:
	// 	status = "PAID"
	// case 2:
	// 	status = "SETTLED"
	// case 3:
	// 	status = "EXPIRED"
	// }

	return dto.OrderResponse{
		InvoiceNumber: order.InvoiceNumber,
		Price: order.Price,
		Quantity: order.Quantity,
		TotalOrder: order.TotalOrder,
		User: dto.Manager{
			Name:  order.User.Name,
			Email: order.User.Email,
		},
		Date:order.Date,
		Place: place,
		Status: order.Status,
		PaymentURL: order.PaymentURL,
		AdminFee: order.AdminFee,
	}
}