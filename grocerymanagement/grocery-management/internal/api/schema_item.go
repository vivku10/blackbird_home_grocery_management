package api

import (
	"fmt"
	"time"
)

type Item struct {
	// Expiration Date Of The Grocery Item
	ExpirationDate string `json:"expirationDate"`

	// Unique Identifier For The Grocery Item
	Id string `json:"id"`

	// Name Of The Grocery Item
	Name string `json:"name"`

	// Quantity Of The Grocery Item
	Quantity int `json:"quantity"`
}

// Item represents a grocery item in the database
type GroceryItem struct {
	id              int
	name            string
	category        string
	quantity        string
	expiration_date time.Time
}

// Checks if all of the required fields for Item are set
// and validates all of the constraints for the object.
func (obj *Item) Validate() error {
	if obj == nil {
		return nil
	}
	fields := map[string]interface{}{
		"id":             obj.Id,
		"name":           obj.Name,
		"quantity":       obj.Quantity,
		"expirationDate": obj.ExpirationDate,
	}

	for field, value := range fields {
		if isEmpty := IsValEmpty(value); isEmpty {
			return fmt.Errorf("required field '%s' for object 'Item' is empty or unset", field)
		}
	}

	return nil
}
