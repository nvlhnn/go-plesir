package service

import (
	"github.com/nvlhnn/go-plesir/helper"
	"github.com/nvlhnn/go-plesir/model/domain"
	"github.com/nvlhnn/go-plesir/model/dto"
	"github.com/nvlhnn/go-plesir/repository"
	"github.com/nvlhnn/go-plesir/schemas"
	"log"
	"strconv"
)

type OrderService interface {
	Create(req dto.OrderCreate) (dto.OrderResponse, schemas.SchemaError)
	FindAllByUserId(id uint) ([]dto.OrderResponse, schemas.SchemaError)
	UpdateStatus(order dto.XenditRequest) (schemas.SchemaError)
	FindByInvoice(invoice string, id uint) (dto.OrderResponse, schemas.SchemaError)
}

type orderService struct{
	orderRepository repository.OrderRepository
	placeRepository repository.PlaceRepository
	xenditService XenditService
}

func NewOrderService(o repository.OrderRepository, p repository.PlaceRepository, xdt XenditService) OrderService{
	return &orderService{o, p, xdt}
}


func (s *orderService) Create(o dto.OrderCreate) (dto.OrderResponse, schemas.SchemaError){

	place, err := s.placeRepository.FindById(o.PlaceID)
	if err.Error != nil {
		return dto.OrderResponse{}, err
	}

	InvoiceNumber := s.orderRepository.CreateInvoiceNumber()

	order := domain.Order{
		InvoiceNumber: InvoiceNumber,
		Date: o.Date,
		PlaceID: o.PlaceID,
		UserID: o.UserID,
		Quantity: o.Quantity,
		Price: place.Price,
		AdminFee: 5000,
		TotalOrder: place.Price*float32(o.Quantity) + 5000,
	}

	
	res, err := s.orderRepository.Save(order)
	if err.Error != nil {
		return dto.OrderResponse{}, err
	}

	invoice, xdtErr := s.xenditService.CreateInvoice(res)
	if xdtErr != nil {
		return dto.OrderResponse{}, schemas.SchemaError{Error: xdtErr}
	}

	res.InvoiceNumber = invoice.ID
	res.PaymentURL = invoice.InvoiceURL

	res, err = s.orderRepository.Update(res)
	if err.Error != nil {
		return dto.OrderResponse{}, err
	}


	result := helper.ToOrderResponse(res)


	log.Println(invoice)
	return result, err
}


func (s *orderService) FindAllByUserId(userId uint)([]dto.OrderResponse, schemas.SchemaError){
	
	orders := []dto.OrderResponse{}
	results, err := s.orderRepository.FindAllByUserId(userId)

	for _, res := range results {
		order := helper.ToOrderResponse(res)
		orders = append(orders, order)
	}


	return orders, err
}


func (s *orderService) FindByInvoice(invoice string, id uint)(dto.OrderResponse, schemas.SchemaError){
	
	// res:= dto.PlaceResponseDTO{}
	// log.Panicln(res)


	result, err := s.orderRepository.FindByUserInvoice(invoice, id)
	res := helper.ToOrderResponse(result)
	
	// log.Panicln(result)

	return res, err
}


func (s *orderService) UpdateStatus(req dto.XenditRequest) ( schemas.SchemaError){

	id, _ := strconv.Atoi(req.ExsternalId)
	result, err := s.orderRepository.FindById(uint(id))
	if err.Error != nil{
		return  err
	}

	result.Status = req.Status
	
	_, err = s.orderRepository.Update(result)
	return err
	
}
