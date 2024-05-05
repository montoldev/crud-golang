package customers

import (
	"crud-golang/database"
	"encoding/json"
	"net/http"
)

func respose(w http.ResponseWriter) {
	response := Response{}
	response.Message = "success"
	jsonResponse, _ := json.Marshal(response)
	w.WriteHeader(http.StatusOK)
	w.Write(jsonResponse)
}

func resposeWithValue(res database.Customers, w http.ResponseWriter) {
	response := Response{}
	response.Message = "success"
	response.Data = &res
	jsonResponse, _ := json.Marshal(response)

	w.WriteHeader(http.StatusOK)
	w.Write(jsonResponse)
}

func resposeError(w http.ResponseWriter, typeError int, messageErr string) {
	response := Response{}
	response.Message = "error"
	response.Error = &messageErr
	jsonResponse, _ := json.Marshal(response)
	w.WriteHeader(typeError)
	w.Write(jsonResponse)
}
