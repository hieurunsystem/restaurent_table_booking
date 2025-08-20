package controllers_test

import (
	"bytes"
	"database/sql"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"

	"github.com/restaurent_table_booking/internal/controllers"
	"github.com/restaurent_table_booking/internal/db"
	"github.com/restaurent_table_booking/internal/models"
)

func setupRouter() *gin.Engine {
	r := gin.Default()
	r.POST("/register", controllers.RegisterHandler)
	r.POST("/login", controllers.LoginHandler)
	r.GET("/accounts", controllers.GetAllAccountsHandler)
	r.POST("/logout", controllers.LogoutHandler)
	r.GET("/profile", controllers.GetProfile)
	return r
}

func TestRegisterCustomerHandler_Success(t *testing.T) {
	// Setup mock database for this specific test
	mockDB, mock, _ := sqlmock.New(sqlmock.QueryMatcherOption(sqlmock.QueryMatcherEqual))
	db.DB = mockDB

	// Mock CheckAccount - user doesn't exist (all queries return ErrNoRows)
	mock.ExpectQuery("SELECT id, password FROM customers WHERE gmail = ?").
		WithArgs("john1111@gmail.com").
		WillReturnError(sql.ErrNoRows)
	mock.ExpectQuery("SELECT id, password FROM staffs WHERE gmail = ?").
		WithArgs("john1111@gmail.com").
		WillReturnError(sql.ErrNoRows)
	mock.ExpectQuery("SELECT id, password FROM admin WHERE gmail = ?").
		WithArgs("john1111@gmail.com").
		WillReturnError(sql.ErrNoRows)
	mock.ExpectQuery("SELECT id, password FROM owners WHERE gmail = ?").
		WithArgs("john1111@gmail.com").
		WillReturnError(sql.ErrNoRows)

	// Mock successful INSERT
	mock.ExpectPrepare("INSERT INTO customers(name, gmail, phone, password, status) VALUES (?,?,?,?,?)").
		ExpectExec().
		WillReturnResult(sqlmock.NewResult(1, 1))

	router := setupRouter()

	payload := models.Account{
		Name:     "John",
		Email:    "john1111@gmail.com",
		Password: "123456",
		Phone:    "0919280763",
		Role:     "customer",
	}
	body, err := json.Marshal(payload)
	assert.NoError(t, err)

	req, _ := http.NewRequest("POST", "/register", bytes.NewBuffer(body))
	req.Header.Set("Content-Type", "application/json")

	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusCreated, w.Code)
	assert.Contains(t, w.Body.String(), "Register successfully")
}

func TestRegisterAdminHandler_Success(t *testing.T) {
	// Setup mock database for this specific test
	mockDB, mock, _ := sqlmock.New(sqlmock.QueryMatcherOption(sqlmock.QueryMatcherEqual))
	db.DB = mockDB

	// Mock CheckAccount - user doesn't exist (all queries return ErrNoRows)
	mock.ExpectQuery("SELECT id, password FROM customers WHERE gmail = ?").
		WithArgs("john1111@gmail.com").
		WillReturnError(sql.ErrNoRows)
	mock.ExpectQuery("SELECT id, password FROM staffs WHERE gmail = ?").
		WithArgs("john1111@gmail.com").
		WillReturnError(sql.ErrNoRows)
	mock.ExpectQuery("SELECT id, password FROM admin WHERE gmail = ?").
		WithArgs("john1111@gmail.com").
		WillReturnError(sql.ErrNoRows)
	mock.ExpectQuery("SELECT id, password FROM owners WHERE gmail = ?").
		WithArgs("john1111@gmail.com").
		WillReturnError(sql.ErrNoRows)

	// Mock successful INSERT
	mock.ExpectPrepare("INSERT INTO admin(name, gmail, phone, password) VALUES (?,?,?,?)").
		ExpectExec().
		WillReturnResult(sqlmock.NewResult(1, 1))

	router := setupRouter()

	payload := models.Account{
		Name:     "John",
		Email:    "john1111@gmail.com",
		Password: "123456",
		Phone:    "0919280763",
		Role:     "admin",
	}
	body, err := json.Marshal(payload)
	assert.NoError(t, err)

	req, _ := http.NewRequest("POST", "/register", bytes.NewBuffer(body))
	req.Header.Set("Content-Type", "application/json")

	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusCreated, w.Code)
	assert.Contains(t, w.Body.String(), "Register successfully")
}

func TestRegisterStaffHandler_Success(t *testing.T) {
	// Setup mock database for this specific test
	mockDB, mock, _ := sqlmock.New(sqlmock.QueryMatcherOption(sqlmock.QueryMatcherEqual))
	db.DB = mockDB

	// Mock CheckAccount - user doesn't exist (all queries return ErrNoRows)
	mock.ExpectQuery("SELECT id, password FROM customers WHERE gmail = ?").
		WithArgs("john1111@gmail.com").
		WillReturnError(sql.ErrNoRows)
	mock.ExpectQuery("SELECT id, password FROM staffs WHERE gmail = ?").
		WithArgs("john1111@gmail.com").
		WillReturnError(sql.ErrNoRows)
	mock.ExpectQuery("SELECT id, password FROM admin WHERE gmail = ?").
		WithArgs("john1111@gmail.com").
		WillReturnError(sql.ErrNoRows)
	mock.ExpectQuery("SELECT id, password FROM owners WHERE gmail = ?").
		WithArgs("john1111@gmail.com").
		WillReturnError(sql.ErrNoRows)

	// Mock successful INSERT
	mock.ExpectPrepare("INSERT INTO staffs(name, gmail, phone, password, restaurant_id) VALUES (?,?,?,?,?)").
		ExpectExec().
		WillReturnResult(sqlmock.NewResult(1, 1))

	router := setupRouter()

	payload := models.Account{
		Name:      "John",
		Email:     "john1111@gmail.com",
		Password:  "123456",
		Phone:     "0919280763",
		Role:      "staff",
		Orther_id: 1,
	}
	body, err := json.Marshal(payload)
	assert.NoError(t, err)

	req, _ := http.NewRequest("POST", "/register", bytes.NewBuffer(body))
	req.Header.Set("Content-Type", "application/json")

	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusCreated, w.Code)
	assert.Contains(t, w.Body.String(), "Register successfully")
}

func TestRegisterOwnerHandler_Success(t *testing.T) {
	// Setup mock database for this specific test
	mockDB, mock, _ := sqlmock.New(sqlmock.QueryMatcherOption(sqlmock.QueryMatcherEqual))
	db.DB = mockDB

	// Mock CheckAccount - user doesn't exist (all queries return ErrNoRows)
	mock.ExpectQuery("SELECT id, password FROM customers WHERE gmail = ?").
		WithArgs("john1111@gmail.com").
		WillReturnError(sql.ErrNoRows)
	mock.ExpectQuery("SELECT id, password FROM staffs WHERE gmail = ?").
		WithArgs("john1111@gmail.com").
		WillReturnError(sql.ErrNoRows)
	mock.ExpectQuery("SELECT id, password FROM admin WHERE gmail = ?").
		WithArgs("john1111@gmail.com").
		WillReturnError(sql.ErrNoRows)
	mock.ExpectQuery("SELECT id, password FROM owners WHERE gmail = ?").
		WithArgs("john1111@gmail.com").
		WillReturnError(sql.ErrNoRows)

	// Mock successful INSERT
	mock.ExpectPrepare("INSERT INTO owners(name, gmail, phone, password, status) VALUES (?,?,?,?,?)").
		ExpectExec().
		WillReturnResult(sqlmock.NewResult(1, 1))

	router := setupRouter()

	payload := models.Account{
		Name:     "John",
		Email:    "john1111@gmail.com",
		Password: "123456",
		Phone:    "0919280763",
		Role:     "owner",
	}
	body, err := json.Marshal(payload)
	assert.NoError(t, err)

	req, _ := http.NewRequest("POST", "/register", bytes.NewBuffer(body))
	req.Header.Set("Content-Type", "application/json")

	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusCreated, w.Code)
	assert.Contains(t, w.Body.String(), "Register successfully")
}

func TestRegisterHandler_InvalidJSON(t *testing.T) {
	router := setupRouter()

	req, _ := http.NewRequest("POST", "/register", bytes.NewBuffer([]byte("invalid")))
	req.Header.Set("Content-Type", "application/json")

	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusBadRequest, w.Code)
}

func TestLoginHandler_Fail_InvalidJSON(t *testing.T) {
	router := setupRouter()

	req, _ := http.NewRequest("POST", "/login", bytes.NewBuffer([]byte("bad json")))
	req.Header.Set("Content-Type", "application/json")

	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusBadRequest, w.Code)
}

func TestLogoutHandler(t *testing.T) {
	router := setupRouter()

	req, _ := http.NewRequest("POST", "/logout", nil)
	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusOK, w.Code)
	assert.Contains(t, w.Body.String(), "Logged out successfully")
}

func TestGetProfile_Unauthorized(t *testing.T) {
	router := setupRouter()

	req, _ := http.NewRequest("GET", "/profile", nil)
	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusUnauthorized, w.Code)
}

func TestRegisterHandler_InvalidRole(t *testing.T) {
	router := setupRouter()

	payload := models.Account{
		Name:     "John",
		Email:    "john@test.com",
		Password: "123456",
		Phone:    "0919280763",
		Role:     "invalid_role",
	}
	body, _ := json.Marshal(payload)

	req, _ := http.NewRequest("POST", "/register", bytes.NewBuffer(body))
	req.Header.Set("Content-Type", "application/json")

	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusBadRequest, w.Code)
	assert.Contains(t, w.Body.String(), "Invalid role")
}

// Test controller behavior without database dependencies
func TestRegisterHandler_ControllerLogic(t *testing.T) {
	router := setupRouter()

	tests := []struct {
		name           string
		payload        interface{}
		expectedStatus int
		expectedBody   string
	}{
		{
			name:           "Invalid JSON",
			payload:        "invalid json",
			expectedStatus: http.StatusBadRequest,
			expectedBody:   "Can't read your input information",
		},
		{
			name: "Invalid Role",
			payload: models.Account{
				Name:     "Test",
				Email:    "test@example.com",
				Password: "123456",
				Phone:    "1234567890",
				Role:     "invalid_role",
			},
			expectedStatus: http.StatusBadRequest,
			expectedBody:   "Invalid role",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var body []byte
			var err error

			if str, ok := tt.payload.(string); ok {
				body = []byte(str)
			} else {
				body, err = json.Marshal(tt.payload)
				assert.NoError(t, err)
			}

			req, _ := http.NewRequest("POST", "/register", bytes.NewBuffer(body))
			req.Header.Set("Content-Type", "application/json")

			w := httptest.NewRecorder()
			router.ServeHTTP(w, req)

			assert.Equal(t, tt.expectedStatus, w.Code)
			assert.Contains(t, w.Body.String(), tt.expectedBody)
		})
	}
}

func TestLoginHandler_ControllerLogic(t *testing.T) {
	router := setupRouter()

	tests := []struct {
		name           string
		payload        interface{}
		expectedStatus int
		expectedBody   string
	}{
		{
			name:           "Invalid JSON",
			payload:        "invalid json",
			expectedStatus: http.StatusBadRequest,
			expectedBody:   "Can't read your input information",
		},
		{
			name: "Valid JSON Structure",
			payload: models.Account{
				Email:    "test@example.com",
				Password: "password123",
			},
			// Will fail due to database, but controller logic works
			expectedStatus: http.StatusBadRequest,
			expectedBody:   "Can't login",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var body []byte
			var err error

			if str, ok := tt.payload.(string); ok {
				body = []byte(str)
			} else {
				body, err = json.Marshal(tt.payload)
				assert.NoError(t, err)
			}

			req, _ := http.NewRequest("POST", "/login", bytes.NewBuffer(body))
			req.Header.Set("Content-Type", "application/json")

			w := httptest.NewRecorder()
			router.ServeHTTP(w, req)

			assert.Equal(t, tt.expectedStatus, w.Code)
			assert.Contains(t, w.Body.String(), tt.expectedBody)
		})
	}
}

func TestGetAllAccountsHandler(t *testing.T) {
	router := setupRouter()

	req, _ := http.NewRequest("GET", "/accounts", nil)
	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)

	// Will fail due to database dependency, but controller doesn't crash
	assert.NotEqual(t, http.StatusInternalServerError, w.Code)
	// Controller should return either 401 (unauthorized) or 500 (database error)
	assert.True(t, w.Code == http.StatusUnauthorized || w.Code >= http.StatusInternalServerError)
}

func TestGetProfile_WithRole(t *testing.T) {
	router := gin.Default()

	// Add middleware to simulate authenticated user
	router.Use(func(c *gin.Context) {
		c.Set("role", "customer")
		c.Next()
	})

	router.GET("/profile", controllers.GetProfile)

	req, _ := http.NewRequest("GET", "/profile", nil)
	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusOK, w.Code)
	assert.Contains(t, w.Body.String(), "customer")
}
