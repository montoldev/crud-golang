package customers

import (
	"crud-golang/database"
	"net/http"
)

type Customers struct {
	Database database.IDatabase
}

//go:generate mockgen -source=./handler.go -destination=./mocks/IHandler.go
type ICustomers interface {
	CreateCustomer(w http.ResponseWriter, r *http.Request) error
}

func NewICustomer(db database.IDatabase) *Customers {
	return &Customers{
		Database: db,
	}
}
