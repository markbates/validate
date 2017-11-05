package validators

import (
	"fmt"

	"github.com/markbates/validate"
)

// IntIsGreaterThan represents the name and the value of the field to compare with another value.
type IntIsGreaterThan struct {
	Name     string
	Field    int
	Compared int
}

// IsValid checks if the field's value is greater than the desired one.
func (v *IntIsGreaterThan) IsValid(errors *validate.Errors) {
	if !(v.Field > v.Compared) {
		errors.Add(GenerateKey(v.Name), fmt.Sprintf("%s: %d is not greater than %d.", v.Name, v.Field, v.Compared))
	}
}
