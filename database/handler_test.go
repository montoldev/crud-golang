package database

import (
	"context"

	"encoding/json"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
)

type MockCustomers struct {
	ID   int    `json:"id" validate:"required"`
	Name string `json:"name" validate:"required"`
	Age  int    `json:"age" validate:"gte=0"`
}

type MockCreateCustomers struct {
	Name string `json:"name" validate:"required"`
	Age  int    `json:"age" validate:"gte=0"`
}
type DatabaseTestSuite struct {
	suite.Suite // เป็นการเปิดฟังก์ชันทดสอบที่ได้รับการให้ความช่วยเหลือจาก Testify suite
	database    *Database
}

func (s *DatabaseTestSuite) SetupSuite() {
	// Create a new instance of the database
	db, err := NewDatabase()
	assert.NoError(s.T(), err)

	// Create the "customers" table if it doesn't exist
	err = db.DB.AutoMigrate(&MockCustomers{}) // Replace YourModelStruct with your actual model struct
	assert.NoError(s.T(), err)

	// Store the database instance in the test suite for later use
	s.database = db
}

func (s *DatabaseTestSuite) TearDownAllSuite() {
	s.database.DB.Migrator().DropTable(&MockCustomers{})
}
func TestCustomersTestSuite(t *testing.T) {
	suite.Run(t, new(DatabaseTestSuite))
}

func (s *DatabaseTestSuite) Test_Insert() {
	bodyReq := `{
		"name": "ronaldo",
		"age": 30
	}`
	customers := MockCreateCustomers{}
	json.Unmarshal([]byte(bodyReq), &customers)
	s.database.Insert(context.Background(), "mock_customers", customers)
}

func (s *DatabaseTestSuite) Test_Inser_errt() {
	bodyReq := `{}`
	customers := MockCreateCustomers{}
	json.Unmarshal([]byte(bodyReq), &customers)
	err := s.database.Insert(context.Background(), "mock", customers)
	s.NotNil(err)
}

func (s *DatabaseTestSuite) Test_Update() {
	bodyReq := `{
		"id": "1"
		"name": "ronaldo",
		"age": 30
	}`
	customers := MockCreateCustomers{}
	json.Unmarshal([]byte(bodyReq), &customers)
	args := []interface{}{1}
	s.database.Update(context.Background(), "mock_customers", customers, "id = ?", args)
}

func (s *DatabaseTestSuite) Test_GetOne() {
	customers := MockCustomers{}
	args := []interface{}{1}
	s.database.GetOne(context.Background(), "mock_customers", &customers, "id = ?", args...)
}

func (s *DatabaseTestSuite) Test_GetOne_Error() {
	customers := MockCustomers{}
	args := []interface{}{1000}
	err := s.database.GetOne(context.Background(), "mock_customers", &customers, "id = ?", args...)
	s.NotNil(err)
}
func (s *DatabaseTestSuite) Test_Delete() {
	args := []interface{}{5}
	s.database.Delete(context.Background(), "mock_customers", "id = ?", args...)
}

func (s *DatabaseTestSuite) Test_Close() {
	s.database.Close()
}
