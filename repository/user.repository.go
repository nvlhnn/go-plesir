package repository

import (
	"github.com/nvlhnn/go-plesir/model/domain"
	"github.com/nvlhnn/go-plesir/schemas"
	"log"

	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

//UserRepository is contract what userRepository can do to db
type UserRepository interface {
	InsertUser(user domain.User) (domain.User, schemas.SchemaError)
	UpdateUser(user domain.User) domain.User
	VerifyCredential(email string, password string) interface{}
	IsDuplicateEmail(email string) (tx *gorm.DB)
	FindByEmail(email string) domain.User
	ProfileUser(userID string) domain.User
}

type userConnection struct {
	connection *gorm.DB
}

//NewUserRepository is creates a new instance of UserRepository
func NewUserRepository(db *gorm.DB) UserRepository {
	return &userConnection{db}
}

func (db *userConnection) InsertUser(user domain.User) (domain.User, schemas.SchemaError) {
	user.Password = hashAndSalt([]byte(user.Password))
	errorResponse := schemas.SchemaError{}
	err := db.connection.Save(&user).Error

	if err != nil {
		errorResponse.Error = err
	}
	return user, errorResponse
}

func (db *userConnection) UpdateUser(user domain.User) domain.User {
	if user.Password != "" {
		user.Password = hashAndSalt([]byte(user.Password))
	} else {
		var tempUser domain.User
		db.connection.Find(&tempUser, user.ID)
		user.Password = tempUser.Password
	}

	db.connection.Save(&user)
	return user
}

func (db *userConnection) VerifyCredential(email string, password string) interface{} {
	var user domain.User
	res := db.connection.Where("email = ?", email).Take(&user)
	if res.Error == nil {
		return user
	}
	return nil
}

func (db *userConnection) IsDuplicateEmail(email string) (tx *gorm.DB) {
	var user domain.User
	return db.connection.Where("email = ?", email).Take(&user)
}

func (db *userConnection) FindByEmail(email string) domain.User {
	var user domain.User
	db.connection.Where("email = ?", email).Take(&user)
	return user
}

func (db *userConnection) ProfileUser(userID string) domain.User {
	var user domain.User
	db.connection.Preload("Places").Preload("Places.User").Find(&user, userID)
	return user
}

func hashAndSalt(pwd []byte) string {
	hash, err := bcrypt.GenerateFromPassword(pwd, bcrypt.MinCost)
	if err != nil {
		log.Println(err)
		panic("Failed to hash a password")
	}
	return string(hash)
}