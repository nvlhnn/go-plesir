package routes

import (
	"github.com/nvlhnn/go-plesir/controller"
	"github.com/nvlhnn/go-plesir/middleware"
	"github.com/nvlhnn/go-plesir/repository"
	"github.com/nvlhnn/go-plesir/service"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)




func RegisterOrderRoutes(router *gin.RouterGroup, db *gorm.DB){

	placeRepository 	:= repository.NewPlaceRepository(db)
	jwtService     		:= service.NewJWTService()
	xenditService := service.NewXenditService()

	orderRepository := repository.NewOrderRepository(db)
	orderService	:= service.NewOrderService(orderRepository, placeRepository, xenditService)
	orderController := controller.NewOrderController(orderService)
	
	router.POST("/xendit_callback", orderController.XenditCallback)
	
	router.Use(middleware.AuthorizeJWT(jwtService))
	{
		router.POST("/", orderController.Create)
		router.GET("/user", orderController.FindAllByUserId)
		router.GET(":invoice_number", orderController.FindByInvoice)
		
	}
}

