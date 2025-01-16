package utils

import (
	"crypto/rand"
	"encoding/base64"
	"fmt"
	"net/http"
	"os/exec"
	"reflect"
	"runtime"
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

func GetContextId(c *gin.Context) (int64, error) {
	idInterface, exists := c.Get("Id")
	if !exists {
		c.JSON(http.StatusBadRequest, gin.H{"error": "ID not found in context"})
		return 0, fmt.Errorf("ID not found in context")
	}

	idUint, ok := idInterface.(uint)
	if !ok {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "ID is not a uint"})
		return 0, fmt.Errorf("ID is not a uint")
	}

	return int64(idUint), nil
}

func GenerateRandomString(n int) string {
	b := make([]byte, n)
	_, err := rand.Read(b)
	if err != nil {
		return ""
	}
	return base64.URLEncoding.EncodeToString(b)
}

func OpenURLInBrowser(url string) error {
	var err error
	switch runtime.GOOS {
	case "linux":
		err = exec.Command("xdg-open", url).Start()
	case "windows":
		err = exec.Command("rundll32", "url.dll,FileProtocolHandler", url).Start()
	case "darwin":
		err = exec.Command("open", url).Start()
	default:
		err = fmt.Errorf("unsupported platform")
	}
	if err != nil {
		return fmt.Errorf("could not open URL in the browser: %w", err)
	}
	return nil
}
