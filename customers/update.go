package customers

import (
	"crud-golang/database"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/go-playground/validator/v10"
)

func (user *Customers) Update(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	var body database.Customers
	err := json.NewDecoder(r.Body).Decode(&body)
	if err != nil {
		resposeError(w, http.StatusBadRequest, "Failed to decode request body")
		return
	}
	validate := validator.New()
	if err := validate.Struct(body); err != nil {
		resposeError(w, http.StatusBadRequest, err.Error())
		return
	}
	condition := "ID = ?"
	args := []interface{}{body.ID}
	err = user.Database.Update(ctx, "customers", body, condition, args...)
	if err != nil {
		msgErr := fmt.Sprintf("DB Error: %v", err.Error())
		resposeError(w, http.StatusInternalServerError, msgErr)
		return
	}
	respose(w)
}
