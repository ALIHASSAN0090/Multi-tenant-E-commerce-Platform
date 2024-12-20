package validation_service_impl

import (
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator"
)

type ValidationServiceImpl struct{}

func (vs *ValidationServiceImpl) ValidateReq(c *gin.Context, request interface{}) []string {
	validate := validator.New()
	var errorMsgs []string
	if err := validate.Struct(request); err != nil {
		ValidationErrors := err.(validator.ValidationErrors)
		for _, fieldError := range ValidationErrors {
			fieldName := fieldError.Field()
			tag := fieldError.Tag()
			errorMsg := fmt.Sprintf("%s is invalid. Error: %s", fieldName, tag)
			errorMsgs = append(errorMsgs, errorMsg)
		}
		return errorMsgs
	}
	return errorMsgs
}
