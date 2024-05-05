package customers

import (
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

func (c *Customers) DeleteById(w http.ResponseWriter, r *http.Request) {

	ctx := r.Context()
	vars := mux.Vars(r)
	condition := "id = ?"
	args := []interface{}{vars["userId"]}

	err := c.Database.Delete(ctx, "customers", condition, args...)
	if err != nil {
		msgErr := fmt.Sprintf("DB Error: %v", err.Error())
		resposeError(w, http.StatusInternalServerError, msgErr)
		return
	}

	respose(w)
}
