package utils

import (
	"reflect"
	"time"

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
