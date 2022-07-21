package controller

import (
	"github.com/nvlhnn/go-plesir/helper"
	"github.com/nvlhnn/go-plesir/model/dto"
	"github.com/nvlhnn/go-plesir/service"
	"log"
	"net/http"
	"reflect"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
)

type OrderController interface {
	Create(ctx *gin.Context)
	FindAllByUserId(ctx *gin.Context)
	FindByInvoice(ctx *gin.Context)
	XenditCallback(ctx *gin.Context)
}

type orderController struct {
	orderService service.OrderService
}

func NewOrderController(s service.OrderService) OrderController {
	return &orderController{s}
}



func (c *orderController) Create(ctx *gin.Context){
	orderReq := dto.OrderCreate{}
	userIdString := ctx.MustGet("userId").(float64)
	// log.Println("ok")
	// log.Println(userIdString)
	// userId, _ := strconv.Atoi(userIdString)
	userId := userIdString
	log.Println(userId)

	orderReq.UserID = uint(userId)

	errDTO := ctx.ShouldBindJSON(&orderReq)

	if errDTO != nil {
		errData := helper.GetErrorData(errDTO)
		helper.FailedValidationResponse(ctx, http.StatusBadRequest, errData, errDTO)
	}else {
		res, err := c.orderService.Create(orderReq)

		if err.Error != nil {
			helper.FailedResponse(ctx, int(err.Code), err.Message, err.Error)
			return
		}

		helper.SuccessResponse(ctx, http.StatusCreated ,"Order created", res)
	}

}

func (c *orderController) FindAllByUserId(ctx *gin.Context){

	var userId float64 = ctx.MustGet("userId").(float64)

	claim := ctx.MustGet("user").(jwt.MapClaims)

	// user := claim.(service.JWTClaim)

	log.Println(reflect.TypeOf(claim["UserID"]))
	
	results, err := c.orderService.FindAllByUserId(uint(userId))
	if err.Error != nil {
		helper.FailedResponse(ctx, int(err.Code), err.Message, err.Error)
	}else{
		helper.SuccessResponse(ctx, http.StatusOK, "get all orders", results)
	}

}


func (c *orderController) FindByInvoice(ctx *gin.Context){

	var userId float64 = ctx.MustGet("userId").(float64)
	invoice := ctx.Param("invoice_number")

	result, err := c.orderService.FindByInvoice(invoice, uint(userId))
	if err.Error != nil {
		helper.FailedResponse(ctx, int(err.Code), err.Message, err.Error)
	}else{
		helper.SuccessResponse(ctx, http.StatusOK, "OK", result)
	}

}


func (c *orderController) XenditCallback(ctx *gin.Context) {
	xenditReq := dto.XenditRequest{}
	errDTO := ctx.ShouldBindJSON(&xenditReq)
	if errDTO != nil {
		errData := helper.GetErrorData(errDTO)
		helper.FailedValidationResponse(ctx, http.StatusBadRequest, errData, errDTO)
	}else {
		err := c.orderService.UpdateStatus(xenditReq)
		if err.Error != nil {
			helper.FailedResponse(ctx, int(err.Code), err.Message, err.Error)
		}else{
			helper.SuccessResponse(ctx, http.StatusOK, "OK", nil)
		}
	}
}
