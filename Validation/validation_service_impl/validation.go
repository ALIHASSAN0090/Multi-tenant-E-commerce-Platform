package validation_service_impl

import (
	"ecommerce-platform/models"
	"errors"
	"fmt"
	"regexp"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator"
)

type ValidationServiceImpl struct{}

func (vs *ValidationServiceImpl) ValidateOauthCreds(req models.OauthUserInfo) bool {
	if req.Email == "" {
		return false
	}
	if !req.VerifiedEmail {
		return false
	}
	return true
}

func (vs *ValidationServiceImpl) ValidateEmailPassword(fl validator.FieldLevel) bool {
	value := fl.Field().String()
	if value == "" {
		return false
	}
	match, _ := regexp.MatchString("^[a-zA-Z0-9]+$", value)
	return match
}

func (vs *ValidationServiceImpl) ValidateReq(c *gin.Context, request interface{}) []string {
	validate := validator.New()

	validate.RegisterValidation("alphanum", func(fl validator.FieldLevel) bool {
		return vs.ValidateEmailPassword(fl)
	})

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

func (vs *ValidationServiceImpl) ValidateOrder(orderData models.CreateOrder) error {

	if orderData.Order.StoreID == 0 {
		return errors.New("store_id is required")
	}

	if len(orderData.OrderItems) == 0 {
		return errors.New("at least one order item is required")
	}

	itemIDMap := make(map[int64]bool)

	for _, item := range orderData.OrderItems {
		if item.ID == 0 {
			return errors.New("order item id is required")
		}
		if item.Quantity <= 0 {
			return errors.New("order item quantity must be greater than zero")
		}
		if _, exists := itemIDMap[item.ID]; exists {
			return errors.New("duplicate order item id found")
		}
		itemIDMap[item.ID] = true
	}

	return nil
}
