package service

import (
	"net/url"

	"github.com/nvlhnn/go-plesir/formatter"
	"github.com/nvlhnn/go-plesir/model/domain"
	"github.com/nvlhnn/go-plesir/model/dto"
	"github.com/nvlhnn/go-plesir/repository"
	"github.com/nvlhnn/go-plesir/schemas"

	"github.com/gin-gonic/gin"
)

type PlaceService interface {
	Create(ctx *gin.Context, req dto.PlaceCreateDTO) (dto.PlaceResponseDTO, schemas.SchemaError)
	FindAll(query url.Values) ([]dto.PlaceResponseDTO, schemas.SchemaError)
	FindById(ctx *gin.Context, id uint) (dto.PlaceResponseDTO, schemas.SchemaError)
	FindBySlug(ctx *gin.Context, slug string) (dto.PlaceResponseDTO, schemas.SchemaError)
	Update(id uint, req dto.PlaceUpdateDTO) (domain.Place, schemas.SchemaError)
	Delete(id uint) (schemas.SchemaError)
}

type placeService struct{
	placeRepository repository.PlaceRepository
	cloudinaryService CloudinaryService
}

func NewPlaceService(placeRepo repository.PlaceRepository, cldService CloudinaryService) PlaceService{
	return &placeService{placeRepo, cldService}
}


func (s *placeService) Create(ctx *gin.Context, p dto.PlaceCreateDTO) (dto.PlaceResponseDTO, schemas.SchemaError){

	err := schemas.SchemaError{}
	result := dto.PlaceResponseDTO{}

	err = s.placeRepository.CheckPlaceExist(p.Name)
	if err.Error != nil {
		return result, err
	}

	err = s.placeRepository.CheckUserExist(p.UserID)
	if err.Error != nil {
		return result, err
	}

	urls := s.cloudinaryService.UploadImages(ctx, p.Images)

	// urlJSON, _ := json.Marshal(urls)
	
	place := formatter.ToPlaceModel(p, urls)

	res, err := s.placeRepository.Save(place)
	if err.Error != nil {
		return result, err
	}

	result = formatter.ToPlaceResponse(res)

	return result, err
}


func (s *placeService) FindAll(query url.Values)([]dto.PlaceResponseDTO, schemas.SchemaError){
	
	var res []dto.PlaceResponseDTO

	results, err := s.placeRepository.FindAll(query)
	if err.Error != nil {
		return res, err
	}

	for _, result := range results {
		item := formatter.ToPlaceResponse(result)
		res = append(res, item)
	}
	return res, err
}


func (s *placeService) FindById(ctx *gin.Context, id uint)(dto.PlaceResponseDTO, schemas.SchemaError){
	
	res := dto.PlaceResponseDTO{}

	result, err := s.placeRepository.FindById(id)
	if err.Error != nil {
		return res, err
	}

	res = formatter.ToPlaceResponse(result)
	
	return res, err
}

func (s *placeService) FindBySlug(ctx *gin.Context, slug string)(dto.PlaceResponseDTO, schemas.SchemaError){
	
	res := dto.PlaceResponseDTO{}

	result, err := s.placeRepository.FindBySlug(slug)
	if err.Error != nil {
		return res, err
	}

	res = formatter.ToPlaceResponse(result)
	
	return res, err
}

func (s *placeService) Update(id uint, p dto.PlaceUpdateDTO)(domain.Place, schemas.SchemaError){
	
	result, err := s.placeRepository.FindById(id)
	if err.Error != nil{
		return result, err
	}

	// log.Println(p.WorkDays)
	if p.Name != nil {
		result.Name = *p.Name
	}

	if p.Price != nil {
		result.Price = *p.Price
	}

	if p.Description != nil {
		result.Description = *p.Description
	}
	if p.UserID != nil {
		result.UserID = *p.UserID
	}

	if len(p.WorkDays) > 0{
		workDays := []domain.PlaceDays{}
		for _, workDay := range p.WorkDays {
			day := domain.PlaceDays{
				DayID: workDay.DayID,
				OpenTime: workDay.Hour.Open,
				CloseTime: workDay.Hour.Close,
			}
			workDays = append(workDays, day)
		}

		result.PlaceDays = workDays
	
	}

	result, err = s.placeRepository.Update(result)
	return result, err
}

func (s *placeService) Delete(id uint)(schemas.SchemaError){
	
	result, err := s.placeRepository.FindById(id)
	if err.Error != nil{
		return err
	}

	err = s.placeRepository.Delete(result)
	return err
}