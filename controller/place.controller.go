package controller

import (
	"github.com/nvlhnn/go-plesir/helper"
	"github.com/nvlhnn/go-plesir/model/dto"
	"github.com/nvlhnn/go-plesir/service"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type PlaceController interface {
	Create(ctx *gin.Context)
	FindAll(ctx *gin.Context)
	FindBySlug(ctx *gin.Context)
	Update(ctx *gin.Context)
	Delete(ctx *gin.Context)
}

type placeController struct {
	placeService service.PlaceService
}

func NewPLaceController(ps service.PlaceService) PlaceController {
	return &placeController{ps}
}




func (c *placeController) Create(ctx *gin.Context){
	placeDTO := dto.PlaceCreateDTO{}
	// userId := ctx.MustGet("userId").(string)
	// placeDTO.UserID, _ = strconv.ParseUint(userId , 10, 64)
	
	errDTO := ctx.Bind(&placeDTO)
	
	if errDTO != nil {
		errData := helper.GetErrorData(errDTO)
		helper.FailedValidationResponse(ctx, http.StatusBadRequest, errData, errDTO)
	}else {
		createdPlace, err := c.placeService.Create(ctx, placeDTO)

		if err.Error != nil {
			helper.FailedResponse(ctx, int(err.Code), err.Message, err.Error)
			return
		}

		helper.SuccessResponse(ctx, http.StatusCreated ,"Place created", createdPlace)
	}

}

func (c *placeController) FindAll(ctx *gin.Context){

	// log.Println(ctx.Request.URL.Query())
	query := ctx.Request.URL.Query()
	results, err := c.placeService.FindAll(query)
	if err.Error != nil {
		helper.FailedResponse(ctx, int(err.Code), err.Message, err.Error)
	}else{
		helper.SuccessResponse(ctx, http.StatusOK, "get all places", results)
	}

}


func (c *placeController) FindBySlug(ctx *gin.Context){

	slug := ctx.Param("slug")
	// slug, _ := strconv.Atoi(idString)
	result, err := c.placeService.FindBySlug(ctx, slug)
	if err.Error != nil {
		helper.FailedResponse(ctx, int(err.Code), err.Message, err.Error)
	}else{
		helper.SuccessResponse(ctx, http.StatusOK, "OK", result)
	}

}


func (c *placeController) Update(ctx *gin.Context){

	placeDTO := dto.PlaceUpdateDTO{}

	errDTO := ctx.ShouldBindJSON(&placeDTO)

	if errDTO != nil {
		helper.FailedResponse(ctx, http.StatusBadRequest, errDTO.Error(), errDTO)
	}else {
		idString := ctx.Param("id")
		id, _ := strconv.Atoi(idString)
	
		result, err := c.placeService.Update(uint(id), placeDTO)
		if err.Error != nil {
			helper.FailedResponse(ctx, int(err.Code), err.Message, err.Error)
		}else{
			helper.SuccessResponse(ctx, http.StatusOK, "OK", result)
		}
	}


}


func (c *placeController) Delete(ctx *gin.Context){

	idString := ctx.Param("id")
	id, _ := strconv.Atoi(idString)

	err := c.placeService.Delete(uint(id))
	if err.Error != nil {
		helper.FailedResponse(ctx, int(err.Code), err.Message, err.Error)
	}else{
		helper.SuccessResponse(ctx, http.StatusOK, "OK", nil)
	}

}