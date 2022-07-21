package seeders

import (
	"errors"
	"github.com/nvlhnn/go-plesir/model/domain"
	"log"

	"gorm.io/gorm"
)

func SeedDay(db *gorm.DB){

	if db.Migrator().HasTable(&domain.Day{}){
		if err := db.First(&domain.Day{}).Error; errors.Is(err, gorm.ErrRecordNotFound) {
			days := []domain.Day{
				{Name:"senin"},
				{Name:"selasa"},
				{Name:"rabu"},
				{Name:"kamis"},
				{Name:"jumat"},
				{Name:"sabtu"},
				{Name:"minggu"},
			}
		
			err := db.Create(&days).Error
			if err != nil {
				log.Println("[seeding error] : ",err )
			}
		}
	}

	return

}