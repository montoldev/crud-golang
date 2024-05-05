package customers

import (
	"context"
	"crud-golang/database"
	mock_database "crud-golang/database/mocks"
	"encoding/json"
	"errors"
	"io"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/gorilla/mux"
	"github.com/stretchr/testify/suite"
)

type CustomersTestSuite struct {
	suite.Suite
	controller *gomock.Controller
	database   *mock_database.MockIDatabase
	customers  *Customers
}

func (s *CustomersTestSuite) SetupSuite() {
	s.controller = gomock.NewController(s.T())
	s.database = mock_database.NewMockIDatabase(s.controller)
	s.customers = &Customers{
		Database: s.database,
	}
}

func (s *CustomersTestSuite) TearDownAllSuite() {
	s.controller.Finish()
	s.database = nil
}

func TestCustomersTestSuite(t *testing.T) {
	suite.Run(t, new(CustomersTestSuite))
}

func (s *CustomersTestSuite) Test_Create() {
	s.Run("When_Create_Error_NewDecoder", func() {

		bodyReq := `invalid request`
		rec := httptest.NewRecorder()
		req := httptest.NewRequest(http.MethodPost, "/customers", strings.NewReader(bodyReq))

		router := mux.NewRouter()
		router.HandleFunc("/customers", s.customers.Create)
		router.ServeHTTP(rec, req)

		resp := rec.Result()
		bodyRes, _ := io.ReadAll(resp.Body)
		s.Contains(string(bodyRes), `{"message":"error","error":"Failed to decode request body"}`)
	})

	s.Run("When_Create_Error_Validator", func() {
		bodyReq := `{}`
		rec := httptest.NewRecorder()
		req := httptest.NewRequest(http.MethodPost, "/customers", strings.NewReader(bodyReq))

		router := mux.NewRouter()
		router.HandleFunc("/customers", s.customers.Create)
		router.ServeHTTP(rec, req)

		resp := rec.Result()
		bodyRes, _ := io.ReadAll(resp.Body)
		s.Contains(string(bodyRes), `{"message":"error","error":"Key: 'CreateCustomer.Name' Error:Field validation for 'Name' failed on the 'required' tag"}`)
	})

	s.Run("When_Create_Error_Database", func() {
		bodyReq := `{
			"name": "ronaldo",
			"age": 30
		}`
		rec := httptest.NewRecorder()
		req := httptest.NewRequest(http.MethodPost, "/customers", strings.NewReader(bodyReq))

		customers := CreateCustomer{}

		json.Unmarshal([]byte(bodyReq), &customers)

		s.database.EXPECT().Insert(gomock.Any(), "customers", customers).Return(errors.New("Insert errors"))

		router := mux.NewRouter()
		router.HandleFunc("/customers", s.customers.Create)
		router.ServeHTTP(rec, req)

		resp := rec.Result()
		bodyRes, _ := io.ReadAll(resp.Body)
		s.Contains(string(bodyRes), `{"message":"error","error":"DB Error: Insert errors"}`)

	})

	s.Run("When_Create_Success", func() {
		bodyReq := `{
			"name": "ronaldo",
			"age": 30
		}`
		rec := httptest.NewRecorder()
		req := httptest.NewRequest(http.MethodPost, "/customers", strings.NewReader(bodyReq))

		customers := CreateCustomer{}

		json.Unmarshal([]byte(bodyReq), &customers)

		s.database.EXPECT().Insert(gomock.Any(), "customers", customers).Return(nil)

		router := mux.NewRouter()
		router.HandleFunc("/customers", s.customers.Create)
		router.ServeHTTP(rec, req)

		resp := rec.Result()
		bodyRes, _ := io.ReadAll(resp.Body)
		s.Contains(string(bodyRes), `{"message":"success"}`)

	})
}

func (s *CustomersTestSuite) Test_update() {

	s.Run("When_Update_Error_NewDecoder", func() {

		bodyReq := `invalid request`
		rec := httptest.NewRecorder()
		req := httptest.NewRequest(http.MethodPut, "/customers", strings.NewReader(bodyReq))

		router := mux.NewRouter()
		router.HandleFunc("/customers", s.customers.Update)
		router.ServeHTTP(rec, req)

		resp := rec.Result()
		bodyRes, _ := io.ReadAll(resp.Body)
		s.Contains(string(bodyRes), `{"message":"error","error":"Failed to decode request body"}`)
	})

	s.Run("When_Update_Error_Validator", func() {
		bodyReq := `{}`
		rec := httptest.NewRecorder()
		req := httptest.NewRequest(http.MethodPut, "/customers", strings.NewReader(bodyReq))

		router := mux.NewRouter()
		router.HandleFunc("/customers", s.customers.Update)
		router.ServeHTTP(rec, req)

		resp := rec.Result()
		bodyRes, _ := io.ReadAll(resp.Body)
		s.Contains(string(bodyRes), `{"message":"error","error":"Key: 'Customers.ID' Error:Field validation for 'ID' failed on the 'required' tag\nKey: 'Customers.Name' Error:Field validation for 'Name' failed on the 'required' tag"}`)
	})

	s.Run("When_Update_Error_Database", func() {
		condition := "id = ?"
		args := []interface{}{1}
		bodyReq := `{
			"id": 1,
			"name": "ronaldo",
			"age": 30
		}`

		rec := httptest.NewRecorder()
		req := httptest.NewRequest(http.MethodPut, "/customers", strings.NewReader(bodyReq))

		customers := database.Customers{}

		json.Unmarshal([]byte(bodyReq), &customers)

		s.database.EXPECT().Update(gomock.Any(), "customers", customers, condition, args...).Return(errors.New("Update errors"))

		router := mux.NewRouter()
		router.HandleFunc("/customers", s.customers.Update)
		router.ServeHTTP(rec, req)

		resp := rec.Result()
		bodyRes, _ := io.ReadAll(resp.Body)
		s.Contains(string(bodyRes), `{"message":"error","error":"DB Error: Update errors"}`)

	})

	s.Run("When_Update_Success", func() {
		condition := "id = ?"
		args := []interface{}{1}
		bodyReq := `{
			"id": 1,
			"name": "ronaldo",
			"age": 30
		}`
		rec := httptest.NewRecorder()
		req := httptest.NewRequest(http.MethodPut, "/customers", strings.NewReader(bodyReq))

		customers := database.Customers{}

		json.Unmarshal([]byte(bodyReq), &customers)

		s.database.EXPECT().Update(gomock.Any(), "customers", customers, condition, args...).Return(nil)

		router := mux.NewRouter()
		router.HandleFunc("/customers", s.customers.Update)
		router.ServeHTTP(rec, req)

		resp := rec.Result()
		bodyRes, _ := io.ReadAll(resp.Body)
		s.Contains(string(bodyRes), `{"message":"success"}`)

	})
}

func (s *CustomersTestSuite) Test_Get() {
	s.Run("When_Get_Error_Database", func() {
		args := []interface{}{"1"}
		rec := httptest.NewRecorder()
		req := httptest.NewRequest(http.MethodGet, "/customers/1", strings.NewReader(``))

		customers := database.Customers{}

		s.database.EXPECT().GetOne(gomock.Any(), "customers", &customers, "id = ?", args...).Return(errors.New("Get errors"))

		router := mux.NewRouter()
		router.HandleFunc("/customers/{userId}", s.customers.GetById)
		router.ServeHTTP(rec, req)

		resp := rec.Result()
		bodyRes, _ := io.ReadAll(resp.Body)
		s.Contains(string(bodyRes), `{"message":"error","error":"DB Error: Get errors"}`)

	})

	s.Run("When_Get_Success", func() {
		args := []interface{}{"1"}

		rec := httptest.NewRecorder()
		req := httptest.NewRequest(http.MethodPut, "/customers/1", strings.NewReader(``))

		customers := database.Customers{}

		s.database.EXPECT().GetOne(gomock.Any(), "customers", &customers, "id = ?", args...).DoAndReturn(func(ctx context.Context, collection string, result interface{}, condition string, args ...interface{}) error {
			// Mock behavior here, for example:
			customer := database.Customers{ID: 1, Name: "John", Age: 30}
			*result.(*database.Customers) = customer // Cast to the correct type and assign mock data
			return nil
		})

		router := mux.NewRouter()
		router.HandleFunc("/customers/{userId}", s.customers.GetById)
		router.ServeHTTP(rec, req)

		resp := rec.Result()
		bodyRes, _ := io.ReadAll(resp.Body)
		s.Contains(string(bodyRes), `{"message":"success","data":{"id":1,"name":"John","age":30}}`)

	})
}

func (s *CustomersTestSuite) Test_Delete() {
	s.Run("When_Delete_Error_Database", func() {
		args := []interface{}{"1"}
		rec := httptest.NewRecorder()
		req := httptest.NewRequest(http.MethodDelete, "/customers/1", strings.NewReader(``))

		s.database.EXPECT().Delete(gomock.Any(), "customers", "id = ?", args...).Return(errors.New("Delete errors"))

		router := mux.NewRouter()
		router.HandleFunc("/customers/{userId}", s.customers.DeleteById)
		router.ServeHTTP(rec, req)

		resp := rec.Result()
		bodyRes, _ := io.ReadAll(resp.Body)
		s.Contains(string(bodyRes), `{"message":"error","error":"DB Error: Delete errors"}`)

	})

	s.Run("When_Delete_Success", func() {
		args := []interface{}{"1"}

		rec := httptest.NewRecorder()
		req := httptest.NewRequest(http.MethodDelete, "/customers/1", strings.NewReader(``))

		s.database.EXPECT().Delete(gomock.Any(), "customers", "id = ?", args...).Return(nil)
		// Mock behavior here, for example:

		router := mux.NewRouter()
		router.HandleFunc("/customers/{userId}", s.customers.DeleteById)
		router.ServeHTTP(rec, req)

		resp := rec.Result()
		bodyRes, _ := io.ReadAll(resp.Body)
		s.Contains(string(bodyRes), `{"message":"success"}`)

	})
}

func TestNewICustomer(t *testing.T) {
	// Create a new instance of the gomock controller
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	// Create a mock database
	mockDB := mock_database.NewMockIDatabase(ctrl)

	// Call the NewICustomer function
	icustomer := NewICustomer(mockDB)

	// Assert that icustomer is not nil
	if icustomer == nil {
		t.Error("Expected non-nil value for icustomer, got nil")
	}

	// Assert that the Database field of icustomer is set to the mock database
	if icustomer.Database != mockDB {
		t.Error("Expected icustomer.Database to be set to mockDB")
	}
}
