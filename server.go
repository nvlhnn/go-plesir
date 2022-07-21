package main

import (
	"github.com/nvlhnn/go-plesir/config"
	"github.com/nvlhnn/go-plesir/middleware"
	"github.com/nvlhnn/go-plesir/routes"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

// test
func main() {
	
	var db  *gorm.DB   = config.OpenConnection()
	
	defer config.CloseConnection(db)



	r := gin.Default()

	r.Use(middleware.CORSMiddleware())
	routes.InitRoutes(r, db)


	// authRoutes := r.Group("api/auth")
	// {
	// 	authRoutes.POST("/login", authController.Login)
	// 	authRoutes.POST("/register", authController.Register)
	// }

	// placeRoutes := r.Group(("api/places"), middleware.AuthorizeJWT(jwtService))
	// {
	// 	placeRoutes.POST("/", placeController.Create)
	// 	placeRoutes.GET("/", placeController.FindAll)
	// 	placeRoutes.GET("/:id", placeController.FindById)
	// 	placeRoutes.PUT("/:id", placeController.Update)
	// 	placeRoutes.DELETE("/:id", placeController.Delete)

	// }

	// orderRoutes := r.Group(("api/orders"), middleware.AuthorizeJWT(jwtService))
	// {
	// 	orderRoutes.POST("/", orderController.Create)
	// }


	
	r.GET("/api/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "ping",
		})
	})
	r.Run(":3000") // listen and serve on 0.0.0.0:8080
}