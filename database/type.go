package database

type Customers struct {
	ID   int    `json:"id" validate:"required"`
	Name string `json:"name" validate:"required"`
	Age  int    `json:"age" validate:"gte=0"`
}
