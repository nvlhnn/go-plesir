package routes

import (
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func InitRoutes(router *gin.Engine, db *gorm.DB) {

	api := router.Group("/api")
	
	auth := api.Group("/auth")
	RegisterAuthRoutes(auth, db)

	place := api.Group("/places")
	RegisterPlaceRoutes(place, db)

	order := api.Group("/orders")
	RegisterOrderRoutes(order, db)

}