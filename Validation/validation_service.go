package Validation

import (
	"ecommerce-platform/Validation/validation_service_impl"
	"ecommerce-platform/models"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator"
)

type ValidationService interface {
	ValidateReq(c *gin.Context, request interface{}) []string
	ValidateEmailPassword(fl validator.FieldLevel) bool
	ValidateOrder(orderData models.CreateOrder) error
	ValidateOauthCreds(req models.OauthUserInfo) bool
	ValidateFilter(filter string) error
}

func NewValidationService() ValidationService {
	return &validation_service_impl.ValidationServiceImpl{}
}
