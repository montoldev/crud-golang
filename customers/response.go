package customers

import (
	"crud-golang/database"
	"encoding/json"
	"net/http"
)

func respose(w http.ResponseWriter) {
	response := Response{}
	response.Message = "success"
	jsonResponse, err := json.Marshal(response)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusOK)
	w.Write(jsonResponse)
}

func resposeWithValue(res database.Customers, w http.ResponseWriter) {
	response := Response{}
	response.Message = "success"
	response.Data = &res
	jsonResponse, err := json.Marshal(response)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusOK)
	w.Write(jsonResponse)
}

func resposeError(w http.ResponseWriter, typeError int, messageErr string) {
	response := Response{}
	response.Message = "error"
	response.Error = &messageErr
	jsonResponse, err := json.Marshal(response)
	if err != nil {
		http.Error(w, err.Error(), typeError)
		return
	}
	w.WriteHeader(http.StatusOK)
	w.Write(jsonResponse)
}
