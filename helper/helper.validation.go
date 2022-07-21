package helper

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/nvlhnn/go-plesir/schemas"
	"io"
	"strings"

	"github.com/go-playground/validator/v10"
)

func getErrorMsg(fe validator.FieldError) string {
	switch fe.Tag() {
	case "required":
		return "This field is required"
	case "numeric":
		return "This field should be a number"
	// case "gte":
	// 	return "Should be greater than " + fe.Param()
	}
	return "Unknown error"
}



func GetErrorData(err error) []schemas.ValidationEror{
	// var ve validator.ValidationErrors
	// if errors.As(err, &ve) {
	// 	out := make([]schemas.ValidationEror, len(ve))
	// 	for i, fe := range ve {
	// 		out[i] = schemas.ValidationEror{fe.Field(), getErrorMsg(fe)}
	// 	}
	// 	return out
	// }

	// return nil



	var collectedErrors []schemas.ValidationEror

	// Other, none data validation errors
	var syntaxError *json.SyntaxError
	var unmarshalTypeError *json.UnmarshalTypeError
	var validationErrors validator.ValidationErrors

	switch {

	// Syntax errors in json
	case errors.As(err, &syntaxError):
		collectedErrors = append(
			collectedErrors,
			schemas.ValidationEror{
				Field:   "general",
				Message: fmt.Sprintf("Request body contains invalid formed Json at position %d", syntaxError.Offset),
			})

	// In some circumstances Decode() may also return an
	// io.ErrUnexpectedEOF error for syntax errors in the JSON.
	case errors.Is(err, io.ErrUnexpectedEOF):
		collectedErrors = append(
			collectedErrors,
			schemas.ValidationEror{
				Field:   "general",
				Message: "Request body contains invalid formed Json",
			})

	// The case when trying to assign not valid type into struct.
	case errors.As(err, &unmarshalTypeError):
		collectedErrors = append(
			collectedErrors,
			schemas.ValidationEror{
				Field:   unmarshalTypeError.Field,
				Message: fmt.Sprintf("Invalid type specified for field %s at position %d", unmarshalTypeError.Field, unmarshalTypeError.Offset),
			})

	// The case when detected extra unexpected field in the request body.
	case strings.HasPrefix(err.Error(), "json: unknown field "):
		fieldName := strings.TrimPrefix(err.Error(), "json: unknown field ")
		collectedErrors = append(
			collectedErrors,
			schemas.ValidationEror{
				Field:   fieldName,
				Message: fmt.Sprintf("Request body contains unknown field %s", fieldName),
			})

	// An io.EOF error is returned by Decode() if the request body is empty.
	case errors.Is(err, io.EOF):
		collectedErrors = append(
			collectedErrors,
			schemas.ValidationEror{
				Field:   "general",
				Message: "Request body must not be empty",
			})

	// The case the request body is too large.
	case err.Error() == "http: request body too large":
		collectedErrors = append(
			collectedErrors,
			schemas.ValidationEror{
				Field:   "general",
				Message: "Request body must not be larger than 1MB",
			})

	// The case the request body is too large.
	case err.Error() == "multipart: NextPart: EOF":
		collectedErrors = append(
			collectedErrors,
			schemas.ValidationEror{
				Field:   "images",
				Message: "Invalid file",
			})

	case errors.As(err, &validationErrors):
		for _, f := range validationErrors {

			msg := "Unknown error"

			switch f.ActualTag() {
			case "required":
				msg = fmt.Sprintf("The field %s Required", f.Field())
			case "email":
				msg = "Should be a valid email address"
			case "lte":
				msg = fmt.Sprintf("Should be less than %s" + f.Param())
			case "gte":
				msg = fmt.Sprintf("Should be greater than %s", f.Param())
			case "alpha":
				msg = "Should be alpha characters only"
			case "numeric":
				msg = "Should be numbers only"
			case "oneof":
				msg = fmt.Sprintf("Should contain one of values %s", f.Param())
			case "url":
				msg = "Should be valid web address starting with http(s)://..."
			case "min":
				msg = fmt.Sprintf("Should be minimum %s characters long", f.Param())
			case "max":
				msg = fmt.Sprintf("Should be maximum %s characters long", f.Param())
			case "e164":
				msg = "Should be valid phone number"
			case "datetime":
				msg = fmt.Sprintf("Should be valid Date/Time with %s format", f.Param())
			}

			//err := f.ActualTag()
			//
			//if f.Param() != "" {
			//	err = fmt.Sprintf("%s=%s", err, f.Param())
			//}
			collectedErrors = append(collectedErrors, schemas.ValidationEror{Field: f.Field(), Message: msg})
		}

	// Server Error response.
	default:
		collectedErrors = append(collectedErrors, schemas.ValidationEror{Field: "general", Message: err.Error()})

	}

	return collectedErrors
}