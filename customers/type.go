package customers

import (
	"crud-golang/database"
)

type CreateCustomer struct {
	Name string `json:"name" validate:"required"`
	Age  int    `json:"age" validate:"gte=0"`
}

type Response struct {
	Message string              `json:"message"`
	Error   *string             `json:"error,omitempty"`
	Data    *database.Customers `json:"data,omitempty"`
}
