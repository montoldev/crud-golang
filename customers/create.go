package customers

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/go-playground/validator/v10"
)

func (c *Customers) Create(w http.ResponseWriter, r *http.Request) {

	ctx := r.Context()
	var body CreateCustomer
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

	err = c.Database.Insert(ctx, "customers", body)
	if err != nil {
		msgErr := fmt.Sprintf("DB Error: %v", err.Error())
		resposeError(w, http.StatusInternalServerError, msgErr)
		return
	}

	respose(w)
}
