package api

import (
	"fmt"
)

type ErrorMessage struct {
	Error	string	`json:"error"`
}

// Checks if all of the required fields for ErrorMessage are set
// and validates all of the constraints for the object.
func (obj *ErrorMessage) Validate() error {
	if obj == nil {
		return nil
	}
	fields := map[string]interface{}{
		"error": obj.Error,
	}

	for field, value := range fields {
		if isEmpty := IsValEmpty(value); isEmpty{
			return fmt.Errorf("required field '%s' for object 'ErrorMessage' is empty or unset", field)
		}
	}

	return nil
}

