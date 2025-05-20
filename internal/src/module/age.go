package module

import "errors"

type Age struct {
	Age int
}

// Validation performs validations regarding the business rule from Age
func (age *Age) Validation() error {
	if age.Age <= 18 || age.Age > 120 {
		return errors.New("the age is not valid")
	}
	return nil
}
