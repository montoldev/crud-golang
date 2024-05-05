package customers

import (
	"crud-golang/database"
)

type Customers struct {
	Database database.IDatabase
}

func NewICustomer(db database.IDatabase) *Customers {
	return &Customers{
		Database: db,
	}
}
