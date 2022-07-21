package domain

import "gorm.io/gorm"

type User struct {
	gorm.Model
	Name     string `gorm:"<-;type:varchar(255)" json:"name"`
	Email string `gorm:"<-uniqueIndex;type:varchar(255)" json:"email"`
	Password  string   `gorm:"->;<-;not null" json:"-"`
	IsAdmin  bool   `gorm:"not null;default:false" json:"is_admin"`
	Token    string  `gorm:"-" json:"token,omitempty"`
	Places    *[]Place `json:"places,omitempty"`

}