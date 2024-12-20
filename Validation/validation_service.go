package Validation

import (
	"ecommerce-platform/Validation/validation_service_impl"

	"github.com/gin-gonic/gin"
)

type ValidationService interface {
	ValidateReq(c *gin.Context, request interface{}) []string
}

func NewValidationService() ValidationService {
	return &validation_service_impl.ValidationServiceImpl{}
}
