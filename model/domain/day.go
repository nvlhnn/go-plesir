package domain

import "gorm.io/gorm"

type Day struct {
	gorm.Model
	Name     string   `gorm:"type:varchar(255);not null;unique" json:"name"`
	// Places   []*Place  `gorm:"many2many:place_days" json:"-"`
	// Time   PlaceDays  
}


