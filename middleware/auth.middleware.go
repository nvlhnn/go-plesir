package middleware

import (
	"errors"
	"github.com/nvlhnn/go-plesir/helper"
	"github.com/nvlhnn/go-plesir/service"
	"log"
	"net/http"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
)

//AuthorizeJWT validates the token user given, return 401 if not valid
func AuthorizeJWT(jwtService service.JWTService) gin.HandlerFunc {
	return func(c *gin.Context) {
		authHeader := c.GetHeader("Authorization")
		if authHeader == "" {

			helper.FailedResponse(c, 401, "", errors.New("no token provided"))
			return
		}
		token, err := jwtService.ValidateToken(authHeader)
		if token.Valid {
			claims := token.Claims.(jwt.MapClaims)
			c.Set("userId", claims["user_id"])
			c.Set("isAdmin", claims["is_admin"])
			c.Set("user", claims)
			log.Print(claims)
			// log.Println("Claim[user_id]: ", claims["user_id"])
			// log.Println("Claim[issuer] :", claims["Issuer"])
		} else {

			helper.FailedResponse(c, 401, "", err)
		}
	}
}

func IsAdmin() gin.HandlerFunc {
	return func(c *gin.Context){
		var isAdmin bool = c.MustGet("isAdmin").(bool)
		if !isAdmin {
			helper.FailedResponse(c, http.StatusForbidden,"", errors.New("forbidden request") )
		}
	}
}