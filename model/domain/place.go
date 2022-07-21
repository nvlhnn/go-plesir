package domain

import "gorm.io/gorm"

type Place struct {
	gorm.Model
	Name        string `gorm:"type:varchar(255);not null" json:"name"`
	Description string `gorm:"type:text;not null" json:"description"`
	Slug		string `gorm:"type:varchar(255);not null" json:"slug"`	 
	Price       float32 `gorm:"type:decimal(18,2);not null" json:"price"`
	Images      []byte `gorm:"type:JSON;not null" json:"images"`
	UserID      uint `gorm:"not null" json:"-"`
	Manager     User   `gorm:"foreignkey:UserID;constraint:onUpdate:CASCADE,onDelete:CASCADE" json:"manager"`
	// Days   		[]Day  `gorm:"many2many:place_days" `
	PlaceDays []PlaceDays `gorm:"foreignKey:PlaceID"`

}


// type User struct {
// 	gorm.Model
// 	Profiles []Profile `gorm:"many2many:user_profiles;foreignKey:Refer;joinForeignKey:UserReferID;References:UserRefer;joinReferences:ProfileRefer"`
// 	Refer    uint      `gorm:"index:,unique"`
//   }
  
//   type Profile struct {
// 	gorm.Model
// 	Name      string
// 	UserRefer uint `gorm:"index:,unique"`
//   }


// {

// // 	days: [
// // 		"name": "asasa",
// // 		"time": {
// // 			open:"asasas"
// // 		},
// // 		"name": "asasa",
// // 		"time": {
// // 			open:"asasas"
// // 		},
// // 		"name": "asasa",
// // 		"time": {
// // 			open:"asasas"
// // 		}

// // 	]
	
// // }