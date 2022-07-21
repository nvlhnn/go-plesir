package routes

import (
	"github.com/nvlhnn/go-plesir/controller"
	"github.com/nvlhnn/go-plesir/middleware"
	"github.com/nvlhnn/go-plesir/repository"
	"github.com/nvlhnn/go-plesir/service"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)




func RegisterPlaceRoutes(router *gin.RouterGroup, db *gorm.DB){

	jwtService     		:= service.NewJWTService()
	placeRepository 	:= repository.NewPlaceRepository(db)
	cloudinaryService 		:= service.NewCloudinaryService()
	placeService 		:= service.NewPlaceService(placeRepository, cloudinaryService)
	placeController		:= controller.NewPLaceController(placeService)

	router.GET("/", placeController.FindAll)
	router.GET("/:slug", placeController.FindBySlug)
	
	router.Use(middleware.AuthorizeJWT(jwtService))
	router.Use(middleware.IsAdmin())
	{
		router.POST("/", placeController.Create)
		router.PUT("/:id", placeController.Update)
		router.DELETE("/:id", placeController.Delete)
	}
}

