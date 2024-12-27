package utils

import (
	"fmt"
	"net/http"
	"reflect"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/mitchellh/mapstructure"
	"golang.org/x/crypto/bcrypt"
)

func CopyTimeHookFunc() mapstructure.DecodeHookFunc {
	return func(
		srcType reflect.Type,
		dstType reflect.Type,
		data interface{},
	) (interface{}, error) {

		if srcType == reflect.TypeOf(time.Time{}) && dstType == reflect.TypeOf(time.Time{}) {

			return data, nil
		}
		return data, nil
	}
}

func Decode(input interface{}, result interface{}) error {
	decoder, err := mapstructure.NewDecoder(&mapstructure.DecoderConfig{
		Metadata: nil,
		DecodeHook: mapstructure.ComposeDecodeHookFunc(
			CopyTimeHookFunc(),
		),
		Result: result,
	})
	if err != nil {
		return err
	}

	if err := decoder.Decode(input); err != nil {
		return err
	}
	return nil
}

func HashPassword(pwd string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(pwd), 14)
	if err != nil {
		return "", err
	}
	return string(bytes), nil
}

func HandleError(err error) error {
	if err != nil {
		return err
	}
	return nil
}

func HandleJsonError(c *gin.Context, err error) {
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	}
}

func VerifyPassword(userPassword string, providedPassword string) (bool, string) {
	err := bcrypt.CompareHashAndPassword([]byte(userPassword), []byte(providedPassword))
	if err != nil {

		return false, "Password does not match"
	}

	return true, "Password matched"
}

func GetDiscountedPrice(originalPrice float32, discount int64) (float64, error) {
	if discount < 0 || discount > 100 {
		return 0, fmt.Errorf("discount must be between 0 and 100")
	}
	discountedPrice := float64(originalPrice) * (1 - float64(discount)/100)
	return discountedPrice, nil
}
