package helper

import (
	"github.com/nvlhnn/go-plesir/schemas"
	"log"

	"github.com/gin-gonic/gin"
)

//Response is used for static shape json return

//EmptyObj object is used when data doesnt want to be null on json
type EmptyObj struct{}

//BuildResponse method is to inject data value to dynamic success response
// func BuildResponse(status bool, message string, data interface{}) Response {
// 	res := Response{
// 		Status:  status,
// 		Message: message,
// 		Errors:  nil,
// 		Data:    data,
// 	}
// 	return res
// }

// //BuildErrorResponse method is to inject data value to dynamic failed response
// func BuildErrorResponse(message string, err string, data interface{}) Response {
// 	splittedError := strings.Split(err, "\n")
// 	res := Response{
// 		Status:  false,
// 		Message: message,
// 		Errors:  splittedError,
// 		Data:    data,
// 	}
// 	return res
// }


//BuildServerErrorResponse method is to create static failed response
// func BuildServerErrorResponse(err string) Response {
// 	emptyObj := EmptyObj{}
// 	res := Response{
// 		Status:  false,
// 		Message: "Internal server error",
// 		Errors:  emptyObj,
// 		Data:    emptyObj,
// 	}

// 	log.Println("error", err)

// 	return res
// }


// BuildResponse method is to inject data value to dynamic success response
func SuccessResponse(ctx *gin.Context, Code int, message string, data interface{}) {
	res := schemas.Response{
		Status:  true,
		Message: message,
		Data:    data,
	}

	ctx.JSON(Code, res)
}

//BuildErrorResponse method is to inject data value to dynamic failed response
func FailedResponse(ctx *gin.Context, Code int, message string, err error ) {
	jsonResponse := schemas.Response{}
	
	if len(message)==0 {
		message = CodeToMessage(Code)
	}

	if Code == 0 {
		Code = 500
	}

	jsonResponse.Status = false
	jsonResponse.Data = nil
	jsonResponse.Message = message

	errors := err.Error()
	log.Println("[error] :",errors)

	ctx.AbortWithStatusJSON(int(Code), jsonResponse)
}


func FailedValidationResponse(ctx *gin.Context, Code int, message interface{}, err error ) {
	jsonResponse := schemas.ErrorValidationResponse{}

	jsonResponse.Status = false
	jsonResponse.Data = nil
	jsonResponse.Message = message

	errors := err.Error()
	log.Println("[error] :",errors)

	ctx.AbortWithStatusJSON(int(Code), jsonResponse)
}





// func APIResponse(ctx *gin.Context, Code uint, Error schemas.SchemaError, Message string, Data interface{}) {
	
// 	emptyObj := EmptyObj{}
// 	jsonResponse := Response{}

// 	if Error.Error == nil{
// 		jsonResponse.Code = Code
// 		jsonResponse.Message = Message
// 		jsonResponse.Data = Data
	
// 		ctx.JSON(int(jsonResponse.Code), jsonResponse)
	

// 	}else{
// 		if Error.Code == 0 {
// 			jsonResponse.Code = 500
// 			jsonResponse.Message = "Internal server error"
// 			jsonResponse.Data = emptyObj
// 		}else{
// 			jsonResponse.Code = Error.Code
// 			jsonResponse.Message = Error.Message
// 			jsonResponse.Data = emptyObj
// 		}

// 		ctx.AbortWithStatusJSON(int(jsonResponse.Code), jsonResponse.Code)

// 	}

// 	// if StatusCode >= 400 {
// 	// 	ctx.AbortWithStatusJSON(StatusCode, jsonResponse)
// 	// } else {
// 	// 	ctx.JSON(StatusCode, jsonResponse)
// 	// }
// }