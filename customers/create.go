package customers

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/go-playground/validator/v10"
)

func (user *Customers) Create(w http.ResponseWriter, r *http.Request) {

	ctx := r.Context()
	var body CreateCustomer
	err := json.NewDecoder(r.Body).Decode(&body)
	if err != nil {
		http.Error(w, "Failed to decode request body", http.StatusBadRequest)
		return
	}

	validate := validator.New()
	if err := validate.Struct(body); err != nil {
		resposeError(w, http.StatusBadRequest, err.Error())
		return
	}

	err = user.Database.Insert(ctx, "customers", body)
	if err != nil {
		msgErr := fmt.Sprintf("DB Error: %v", err.Error())
		http.Error(w, msgErr, http.StatusInternalServerError)
		return
	}

	respose(w)
}
