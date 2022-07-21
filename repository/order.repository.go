package repository

import (
	"errors"
	"github.com/nvlhnn/go-plesir/model/domain"
	"github.com/nvlhnn/go-plesir/schemas"
	"log"
	"net/http"
	"strconv"
	"time"

	"gorm.io/gorm"
)

type OrderRepository interface {
	Save(order domain.Order) (domain.Order, schemas.SchemaError)
	Update(order domain.Order) (domain.Order, schemas.SchemaError)
	FindAllByUserId(userId uint ) ([]domain.Order, schemas.SchemaError)
	CreateInvoiceNumber() string
	FindByUserInvoice(invoice string, userId uint) (domain.Order, schemas.SchemaError)
	FindById(id uint) (domain.Order, schemas.SchemaError)

}


type orderRepository struct{
	db *gorm.DB
}

//NewUserRepository is creates a new instance of UserRepository
func NewOrderRepository(db *gorm.DB) OrderRepository {
	return &orderRepository{db}
}



func (r *orderRepository) Save (order domain.Order) (domain.Order, schemas.SchemaError){

	errorResponse := schemas.SchemaError{}

	// saving data
	saveErr := r.db.Save(&order).Error
	if saveErr != nil {
		errorResponse.Error = saveErr
		return order, errorResponse
	}

	// get data
	getErr := r.db.Preload("User").Preload("Place").Find(&order).Error
	if saveErr != nil {
		errorResponse.Error = getErr
		return order, errorResponse
	}

	return order, errorResponse
} 


func (r *orderRepository) Update (order domain.Order) (domain.Order, schemas.SchemaError){

	errorResponse := schemas.SchemaError{}

	// saving data
	saveErr := r.db.Updates(&order).Error
	if saveErr != nil {
		errorResponse.Error = saveErr
		return order, errorResponse
	}

	return order, errorResponse
} 

func (r *orderRepository) FindAllByUserId (userId uint) ([]domain.Order, schemas.SchemaError){
	results := []domain.Order{}
	errorResponse := schemas.SchemaError{}

	err := r.db.Preload("User").Preload("Place").Where("user_id = ?", userId).Order("created_at desc").Find(&results).Error
	if err != nil {
		errorResponse.Error = err
		return results, errorResponse
	}
	
	return results, errorResponse

}

func (r *orderRepository) CreateInvoiceNumber () string {

	var counter int
	r.db.Raw("SELECT COUNT(id) FROM orders WHERE MONTH(created_at) = MONTH(CURRENT_DATE()) AND YEAR(created_at) = YEAR(CURRENT_DATE())").Scan(&counter)
	date := time.Now().Format("060102")

	return "INV-"+date+"-"+strconv.Itoa(counter+1)

}

func (r *orderRepository) FindByUserInvoice (invoice string, userId uint) (domain.Order, schemas.SchemaError){

	errorResponse := schemas.SchemaError{}
	order := domain.Order{}

	log.Println(invoice, userId)
	res := r.db.Debug().Preload("Place").Where("invoice_number = ? AND user_id = ?", invoice, userId).First(&order)
	if(res.Error != nil){
		errorResponse.Error = res.Error
		return order, errorResponse
	}

	if res.RowsAffected == 0 {
		errorResponse.Error = errors.New("Order not found")
		errorResponse.Code = http.StatusNotFound
		return order, errorResponse
	}

	return order, errorResponse
}


func (r *orderRepository) FindById (id uint) (domain.Order, schemas.SchemaError){

	errorResponse := schemas.SchemaError{}
	order := domain.Order{}

	res := r.db.Preload("User").Preload("Place").Find(&order, id)
	if(res.Error != nil){
		errorResponse.Error = res.Error
		return order, errorResponse
	}

	if res.RowsAffected == 0 {
		errorResponse.Error = errors.New("Order not found")
		errorResponse.Code = http.StatusNotFound
		return order, errorResponse
	}

	return order, errorResponse
}