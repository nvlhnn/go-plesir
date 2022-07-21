package service

import (
	"github.com/nvlhnn/go-plesir/model/domain"
	"os"
	"strconv"

	"github.com/xendit/xendit-go"
	"github.com/xendit/xendit-go/invoice"
)

type XenditService interface {
	CreateInvoice(order domain.Order) (*xendit.Invoice, *xendit.Error)
	GetInvoice(id string) (*xendit.Invoice, *xendit.Error) 
	NotifCallback()
}

type xenditService struct {
	secret string
}

func NewXenditService() XenditService {

	secret := os.Getenv("XENDIT_SECRET")

	return &xenditService{secret}
}

func (s *xenditService) CreateInvoice(order domain.Order ) (*xendit.Invoice, *xendit.Error) {


	// err := schemas.SchemaError{}

	xendit.Opt.SecretKey = s.secret

	customer := xendit.InvoiceCustomer{
		GivenNames:   order.User.Name,
		Email:        order.User.Email,
		
		// MobileNumber: "+6289504767222",
	}
	  
	item := xendit.InvoiceItem{
		Name:          order.Place.Name,
		Quantity:       int(order.Quantity),
		Price:          float64(order.Place.Price),
		// Category:       "Electronic",
		Url:            "http://localhost:5000/places/"+order.Place.Slug,
	}
	  
	items := []xendit.InvoiceItem{item}
	  
	fee := xendit.InvoiceFee{
		Type:         "ADMIN",
		Value:        5000,
	}
	   
	fees := []xendit.InvoiceFee{fee}
	  
	NotificationType := []string{"email"}
	  
	customerNotificationPreference := xendit.InvoiceCustomerNotificationPreference{
		InvoiceCreated:     NotificationType,
		InvoiceReminder:    NotificationType,
		InvoicePaid:        NotificationType,
		InvoiceExpired:     NotificationType,
	}
	  
	data := invoice.CreateParams{
		ExternalID:         strconv.Itoa(int(order.ID)),
		Amount:             float64(order.TotalOrder),
		Description:        "Pembayaran tiket go plesir",
		InvoiceDuration:    86400,
		Customer:           customer, 
		CustomerNotificationPreference:   customerNotificationPreference, 
		SuccessRedirectURL: "http://localhost:5000/orders/",
		FailureRedirectURL: "http://localhost:5000/orders/",
		Currency:           "IDR",
		Items:              items,  
		Fees:               fees,
	}
	  
	resp, err := invoice.Create(&data)
	if err != nil {
		return &xendit.Invoice{}, err
	}

	  
	return resp, err
}

func (s *xenditService) GetInvoice(id string) (*xendit.Invoice, *xendit.Error) {

	xendit.Opt.SecretKey = s.secret

	data := invoice.GetParams{
		ID: id,
	}

	resp, err := invoice.Get(&data)
	if err != nil {
		return &xendit.Invoice{}, err
	}

	return resp, err

}

func (s *xenditService) NotifCallback() {

}
