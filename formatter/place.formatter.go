package formatter

import (
	"encoding/json"
	"github.com/nvlhnn/go-plesir/model/domain"
	"github.com/nvlhnn/go-plesir/model/dto"
	"log"
	"math/rand"
	"strconv"

	"github.com/gosimple/slug"
)

func ToPlaceResponse(place domain.Place) dto.PlaceResponseDTO {

	var urls []string
	err := json.Unmarshal(place.Images, &urls)
	if err != nil {
		log.Println(err)
	}
	log.Println(urls)

	workDays := []dto.Work{}
	for _, workDay := range place.PlaceDays {
		day := dto.Work{
			Day: workDay.Day.Name,
			Hour: dto.Hour{
				Open:  workDay.OpenTime,
				Close: workDay.CloseTime,
			},
		}
		workDays = append(workDays, day)
	}

	return dto.PlaceResponseDTO{
		ID: place.ID,
		Name:        place.Name,
		Description: place.Description,
		Price:       place.Price,
		Manager: dto.Manager{
			Name:  place.Manager.Name,
			Email: place.Manager.Email,
		},
		WorkDays: workDays,
		Images: urls,
		Slug: place.Slug,
	}
}

func ToPlaceModel(p dto.PlaceCreateDTO, urls []byte) domain.Place {


	workDays := []domain.PlaceDays{}
	rand := rand.Intn(99999-11111) + 11111
	slug := slug.Make(p.Name) + "-" + strconv.Itoa(rand)

	for _, workDay := range p.WorkDays {
		day := domain.PlaceDays{
			DayID: workDay.DayID,
			OpenTime: workDay.Hour.Open,
			CloseTime: workDay.Hour.Close,
		}
		workDays = append(workDays, day)
	}


	place := domain.Place{
		Name: p.Name,
		Description: p.Description,
		Price: p.Price,
		UserID: p.UserID,
		PlaceDays: workDays,
		Images: urls,
		Slug: slug,
	}

	return place

}