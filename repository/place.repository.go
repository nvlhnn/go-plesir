package repository

import (
	"errors"
	"log"
	"net/http"
	"net/url"
	"strconv"
	"strings"

	"github.com/nvlhnn/go-plesir/model/domain"
	"github.com/nvlhnn/go-plesir/schemas"

	"gorm.io/gorm"
)

type PlaceRepository interface {
	Save(place domain.Place) (domain.Place, schemas.SchemaError)
	FindAll(query url.Values) ([]domain.Place, schemas.SchemaError)
	FindById(id uint) (domain.Place, schemas.SchemaError)
	FindBySlug(slug string) (domain.Place, schemas.SchemaError)
	Update(place domain.Place) (domain.Place, schemas.SchemaError)
	Delete(place domain.Place) (schemas.SchemaError)
	CheckPlaceExist(name string) (schemas.SchemaError)
	CheckUserExist(id uint) (schemas.SchemaError)
}


type placeRepository struct{
	db *gorm.DB
}

//NewUserRepository is creates a new instance of UserRepository
func NewPlaceRepository(db *gorm.DB) PlaceRepository {
	return &placeRepository{db}
}



func (r *placeRepository) Save (place domain.Place) (domain.Place, schemas.SchemaError){

	errorResponse := schemas.SchemaError{}

	// saving data
	saveErr := r.db.Save(&place).Error
	if saveErr != nil {
		errorResponse.Error = saveErr
		return place, errorResponse
	}

	getErr := r.db.Preload("Manager").Preload("PlaceDays.Day").Find(&place).Error
	if saveErr != nil {
		errorResponse.Error = getErr
		return place, errorResponse
	}

	return place, errorResponse
} 

func (r *placeRepository) FindAll (query url.Values) ([]domain.Place, schemas.SchemaError){
	results := []domain.Place{}
	errorResponse := schemas.SchemaError{}

	tx := r.db.Debug().Preload("Manager")

	// filter name
	if val, ok := query["search"]; ok {
		tx.Where("name LIKE ?", "%"+val[0]+"%")
	}
	
	// sorting
	if val, ok := query["sort"]; ok {
		switch val[0] {
		case "latest":
			tx.Order("created_at desc")		
		case "oldest":
			tx.Order("created_at asc")
		case "cheapest":
			tx.Order("price asc")
		case "expensive":
			tx.Order("price desc")
		default:
			tx.Order("created_at desc")		
		}		
	}

	// filter days
	if val, ok := query["days"]; ok && val[0] != "null" {
		day_ids := strings.Split(val[0], ",")

		var day_uint []uint
		for _, v := range day_ids {
			numb, _ := strconv.Atoi(v)
			day_uint = append(day_uint, uint(numb))
		}
		
		tx.Where("id IN (?)", 
			r.db.Table("place_days").
			Select("place_id").
			Where("day_id IN ?", day_uint),
		)
	}else{
		tx.Where("id IN (?)", 
			r.db.Table("place_days").Select("place_id"),
		)
	}


	// pagination 
	if val, ok := query["page"]; ok {
		page, _ := strconv.Atoi(val[0])
		offset := (page - 1) * 12
		tx.Limit(12).Offset(offset)
	}
		
	err := tx.Find(&results).Error

	// log.Println(results)
	// err := r.db.Preload("Manager").Find(&results).Error
	if err != nil {
		errorResponse.Error = err
		return results, errorResponse
	}
	
	return results, errorResponse
}

func (r *placeRepository) FindById (id uint) (domain.Place, schemas.SchemaError){

	errorResponse := schemas.SchemaError{}
	place := domain.Place{}

	res := r.db.Preload("Manager").Preload("PlaceDays.Day").Find(&place, id)
	if(res.Error != nil){
		errorResponse.Error = res.Error
		return place, errorResponse
	}

	if res.RowsAffected == 0 {
		errorResponse.Error = errors.New("Place not found")
		errorResponse.Code = http.StatusNotFound
		return place, errorResponse
	}

	return place, errorResponse
}

func (r *placeRepository) FindBySlug (slug string) (domain.Place, schemas.SchemaError){

	errorResponse := schemas.SchemaError{}
	place := domain.Place{}

	res := r.db.Preload("Manager").Preload("PlaceDays.Day").Where("slug = ?", slug).First(&place)
	if(res.Error != nil){
		errorResponse.Error = res.Error
		return place, errorResponse
	}

	if res.RowsAffected == 0 {
		errorResponse.Error = errors.New("Place not found")
		errorResponse.Code = http.StatusNotFound
		return place, errorResponse
	}

	return place, errorResponse
}

func (r *placeRepository) Update (place domain.Place) (domain.Place, schemas.SchemaError){

	errorResponse := schemas.SchemaError{}

	// update data
	err := r.db.Session(&gorm.Session{FullSaveAssociations: true}).Updates(&place).Error
	if(err!= nil){
		errorResponse.Error = err
	}

	// update relation
	return place, errorResponse
}

func (r *placeRepository) Delete (place domain.Place) (schemas.SchemaError){

	errorResponse := schemas.SchemaError{}
	err := r.db.Delete(&place).Error
	log.Println(err)
	if(err!= nil){
		errorResponse.Error = err
	}

	return errorResponse
}

func (r *placeRepository) CheckPlaceExist (name string) (schemas.SchemaError){
	errorResponse := schemas.SchemaError{}
	existErr := r.db.Where("name = ?", name).First(&domain.Place{})
	if existErr.RowsAffected > 0 {
		errorResponse.Error = errors.New("Place already exist")
		errorResponse.Code = http.StatusConflict
		errorResponse.Message = "Place already exist"
	}

	return errorResponse
}

func (r *placeRepository) CheckUserExist (id uint) (schemas.SchemaError){
	errorResponse := schemas.SchemaError{}
	existErr := r.db.First(&domain.User{}, id)
	if existErr.RowsAffected == 0 {
		errorResponse.Error = errors.New("User doesnt exist")
		errorResponse.Code = http.StatusNotFound
		errorResponse.Message = "User doesnt exist"
	}

	return errorResponse
}
