package controller

import (
	"errors"
	"net/http"

	"github.com/nvlhnn/go-plesir/helper"
	"github.com/nvlhnn/go-plesir/model/domain"
	"github.com/nvlhnn/go-plesir/model/dto"
	"github.com/nvlhnn/go-plesir/service"

	"github.com/gin-gonic/gin"
)

type AuthController interface {
	Login(ctx *gin.Context)
	Register(ctx *gin.Context)
}

type authController struct {
	authService service.AuthService
	jwtService  service.JWTService
}

//NewAuthController creates a new instance of AuthController
func NewAuthController(authService service.AuthService, jwtService service.JWTService) AuthController {
	return &authController{
		authService: authService,
		jwtService:  jwtService,
	}
}

func (c *authController) Login(ctx *gin.Context) {
	var loginDTO dto.LoginDTO
	errDTO := ctx.ShouldBind(&loginDTO)
	if errDTO != nil {
		helper.FailedResponse(ctx, http.StatusBadRequest, errDTO.Error(), errDTO)
		return
	}
	authResult := c.authService.VerifyCredential(loginDTO.Email, loginDTO.Password)
	if v, ok := authResult.(domain.User); ok {
		generatedToken := c.jwtService.GenerateToken(v)
		v.Token = generatedToken
		helper.SuccessResponse(ctx, http.StatusOK, "login success", v)
		return
	}

	helper.FailedResponse(ctx, http.StatusUnauthorized, "Invalid credential", errors.New("credential failed") )
}

func (c *authController) Register(ctx *gin.Context) {
	var registerDTO dto.RegisterDTO
	errDTO := ctx.ShouldBind(&registerDTO)
	if errDTO != nil {
		errData := helper.GetErrorData(errDTO)
		helper.FailedValidationResponse(ctx, http.StatusBadRequest, errData, errDTO)
		return
	}

	if !c.authService.IsDuplicateEmail(registerDTO.Email) {
		helper.FailedResponse(ctx, http.StatusConflict, "Duplicate email", errors.New("Duplicate email"))
		return
	} 

	createdUser, err := c.authService.CreateUser(registerDTO)
	if err.Error != nil {
		helper.FailedResponse(ctx, int(err.Code), "", err.Error)
		return
	}

	token := c.jwtService.GenerateToken(createdUser)
	createdUser.Token = token

	helper.SuccessResponse(ctx, http.StatusOK ,"User created", createdUser)

	
	
}