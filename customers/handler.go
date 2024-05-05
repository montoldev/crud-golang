package customers

import (
	"crud-golang/database"
	"net/http"
)

type Customers struct {
	Database database.IDatabase
}

type ICustomers interface {
	CreateCustomer(w http.ResponseWriter, r *http.Request) error
}

func NewIUser(db database.IDatabase) *Customers {
	return &Customers{
		Database: db,
	}
}
