package utils

import (
	"fmt"
	"reflect"
	"strconv"
	"strings"
	"time"
)

func SetFieldValue(fieldValue reflect.Value, value, envTag string) error {
	switch fieldValue.Kind() {
	case reflect.String:
		fieldValue.SetString(value)
	case reflect.Int:
		intValue, err := strconv.Atoi(value)
		if err != nil {
			return fmt.Errorf("invalid integer value for %v: %v", envTag, err)
		}
		fieldValue.SetInt(int64(intValue))
	case reflect.Bool:
		boolValue, err := strconv.ParseBool(value)
		if err != nil {
			return fmt.Errorf("invalid boolean value for %v: %v", envTag, err)
		}
		fieldValue.SetBool(boolValue)
	case reflect.Slice:
		if fieldValue.Type().Elem().Kind() == reflect.String {
			sliceValue := strings.Split(value, ",")
			fieldValue.Set(reflect.ValueOf(sliceValue))
		}
	case reflect.Int64:
		if fieldValue.Type() == reflect.TypeOf(time.Duration(0)) {
			durationValue, err := time.ParseDuration(value)
			if err != nil {
				return fmt.Errorf("invalid duration value for %v: %v", envTag, err)
			}
			fieldValue.SetInt(int64(durationValue))
		}
	}
	return nil
}
