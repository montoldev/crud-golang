package customers

import (
	"crud-golang/database"
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

func (c *Customers) GetById(w http.ResponseWriter, r *http.Request) {

	ctx := r.Context()
	vars := mux.Vars(r)
	condition := "id = ?"
	var res database.Customers
	args := []interface{}{vars["userId"]}

	err := c.Database.GetOne(ctx, "customers", &res, condition, args...)
	if err != nil {
		msgErr := fmt.Sprintf("DB Error: %v", err.Error())
		resposeError(w, http.StatusInternalServerError, msgErr)
		return
	}

	resposeWithValue(res, w)
}
