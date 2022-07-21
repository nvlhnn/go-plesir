package domain

import (
	"strconv"

	"gorm.io/gorm"
)

type PlaceDays struct {
	ID        string `gorm:"primaryKey;unique;not null"`

	PlaceID   uint    `gorm:"primaryKey" json:"-"`
	DayID     uint    `gorm:"primaryKey" json:"-"`
	
	Place     Place  `gorm:"foreignKey:PlaceID"`
	Day       Day    `gorm:"foreignKey:DayID"`

	OpenTime  string `gorm:"varchar(255);not null" json:"open_time"`
	CloseTime string `gorm:"varchar(255);not null" json:"close_time"`
}

func (placeDay *PlaceDays) BeforeSave(tx *gorm.DB) (err error) {
	placeDay.ID = strconv.Itoa(int(placeDay.PlaceID)) + strconv.Itoa(int(placeDay.DayID))
	return
}
