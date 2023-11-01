package validator

import (
	"github.com/asaskevich/govalidator"
)

func ValidateStruct(obj interface{}) error {
	_, err := govalidator.ValidateStruct(obj)
	if err != nil {
		return err
	}
	return nil
}
