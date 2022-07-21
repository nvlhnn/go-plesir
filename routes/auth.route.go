package routes

import (
	"github.com/nvlhnn/go-plesir/controller"
	"github.com/nvlhnn/go-plesir/repository"
	"github.com/nvlhnn/go-plesir/service"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)




func RegisterAuthRoutes(router *gin.RouterGroup, db *gorm.DB){

	userRepository := repository.NewUserRepository(db)	
	jwtService     := service.NewJWTService()
	authService    := service.NewAuthService(userRepository)
	authController := controller .NewAuthController(authService, jwtService)

	router.POST("/login", authController.Login)
	router.POST("/register", authController.Register)
	
}
