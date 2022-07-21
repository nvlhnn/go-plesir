package service

import (
	"github.com/nvlhnn/go-plesir/model/domain"
	"github.com/nvlhnn/go-plesir/model/dto"
	"github.com/nvlhnn/go-plesir/repository"
	"github.com/nvlhnn/go-plesir/schemas"
	"log"

	"golang.org/x/crypto/bcrypt"
)

type AuthService interface {
	VerifyCredential(email string, password string) interface{}
	CreateUser(user dto.RegisterDTO) (domain.User, schemas.SchemaError)
	FindByEmail(email string) domain.User
	IsDuplicateEmail(email string) bool
}

type authService struct {
	userRepository repository.UserRepository
}

//NewAuthService creates a new instance of AuthService
func NewAuthService(userRep repository.UserRepository) AuthService {
	return &authService{
		userRepository: userRep,
	}
}

func (service *authService) VerifyCredential(email string, password string) interface{} {
	res := service.userRepository.VerifyCredential(email, password)
	if v, ok := res.(domain.User); ok {
		comparedPassword := comparePassword(v.Password, []byte(password))
		if v.Email == email && comparedPassword {
			return res
		}
		return false
	}
	return false
}

func (service *authService) CreateUser(user dto.RegisterDTO) (domain.User, schemas.SchemaError) {
	userToCreate := domain.User{
		Name: user.Name,
		Email: user.Email,
		Password: user.Password,
	}
	// err := smapping.FillStruct(&userToCreate, smapping.MapFields(&user))
	// if err != nil {
	// 	log.Fatalf("Failed map %v", err)
	// }
	res, err := service.userRepository.InsertUser(userToCreate)
	return res, err
}

func (service *authService) FindByEmail(email string) domain.User {
	return service.userRepository.FindByEmail(email)
}

func (service *authService) IsDuplicateEmail(email string) bool {
	res := service.userRepository.IsDuplicateEmail(email)
	return !(res.Error == nil)
}

func comparePassword(hashedPwd string, plainPassword []byte) bool {
	byteHash := []byte(hashedPwd)
	err := bcrypt.CompareHashAndPassword(byteHash, plainPassword)
	if err != nil {
		log.Println(err)
		return false
	}
	return true
}