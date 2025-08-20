package tests

import (
	"database/sql"
	"database/sql/driver"
	"fmt"
	"os"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/restaurent_table_booking/internal/db"
	"github.com/restaurent_table_booking/internal/models"
)

// TestDB holds the mock database connection and mock for testing
type TestDB struct {
	DB   *sql.DB
	Mock sqlmock.Sqlmock
}

// SetupTestDB creates a mock database for testing
func SetupTestDB(t *testing.T) *TestDB {
	mockDB, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("Failed to create mock database: %v", err)
	}

	// Store original DB
	originalDB := db.DB
	
	// Replace with mock
	db.DB = mockDB

	// Cleanup function
	t.Cleanup(func() {
		mockDB.Close()
		db.DB = originalDB
	})

	return &TestDB{
		DB:   mockDB,
		Mock: mock,
	}
}

// CreateTestAccount creates a test account with default values
func CreateTestAccount(role string) *models.Account {
	baseAccount := &models.Account{
		Name:     "Test User",
		Email:    fmt.Sprintf("test_%s@example.com", role),
		Phone:    "1234567890",
		Password: "testpassword123",
		Role:     role,
	}

	if role == "staff" {
		baseAccount.Orther_id = 1
	}

	return baseAccount
}

// CreateTestAccounts creates multiple test accounts for different roles
func CreateTestAccounts() map[string]*models.Account {
	roles := []string{"customer", "admin", "owner", "staff"}
	accounts := make(map[string]*models.Account)

	for _, role := range roles {
		accounts[role] = CreateTestAccount(role)
	}

	return accounts
}

// MockCheckAccountExists mocks the CheckAccount function for existing account
func MockCheckAccountExists(mock sqlmock.Sqlmock, email, role, password string, id int64) {
	query := fmt.Sprintf("SELECT id, password FROM %ss WHERE gmail = ?", role)
	if role == "admin" {
		query = "SELECT id, password FROM admin WHERE gmail = ?"
	}
	
	mock.ExpectQuery(query).
		WithArgs(email).
		WillReturnRows(sqlmock.NewRows([]string{"id", "password"}).
			AddRow(id, password))
}

// MockCheckAccountNotExists mocks the CheckAccount function for non-existing account
func MockCheckAccountNotExists(mock sqlmock.Sqlmock, email string) {
	roles := []string{"customer", "staff", "admin", "owner"}
	
	for _, role := range roles {
		query := fmt.Sprintf("SELECT id, password FROM %ss WHERE gmail = ?", role)
		if role == "admin" {
			query = "SELECT id, password FROM admin WHERE gmail = ?"
		}
		
		mock.ExpectQuery(query).
			WithArgs(email).
			WillReturnError(sql.ErrNoRows)
	}
}

// MockInsertAccount mocks successful account insertion
func MockInsertAccount(mock sqlmock.Sqlmock, table string, id int64, args ...interface{}) {
	placeholders := ""
	for i := range args {
		if i > 0 {
			placeholders += ","
		}
		placeholders += "?"
	}
	
	// Convert []interface{} to []driver.Value
	driverArgs := make([]driver.Value, len(args))
	for i, arg := range args {
		driverArgs[i] = arg
	}
	
	query := fmt.Sprintf("INSERT INTO %s VALUES \\(%s\\)", table, placeholders)
	mock.ExpectPrepare(query).
		ExpectExec().
		WithArgs(driverArgs...).
		WillReturnResult(sqlmock.NewResult(id, 1))
}

// MockGetAllAccounts mocks the GetAllAccounts function
func MockGetAllAccounts(mock sqlmock.Sqlmock, accounts []models.Account) {
	rows := sqlmock.NewRows([]string{"id", "gmail", "name", "phone", "password"})
	
	for _, account := range accounts {
		rows.AddRow(account.Id, account.Email, account.Name, account.Phone, account.Password)
	}
	
	mock.ExpectQuery("SELECT \\* FROM users UNION SELECT \\* FROM admin").
		WillReturnRows(rows)
}

// SetupTestEnvironment sets up the test environment
func SetupTestEnvironment() {
	os.Setenv("GIN_MODE", "test")
}

// CleanupTestEnvironment cleans up the test environment
func CleanupTestEnvironment() {
	os.Unsetenv("GIN_MODE")
}

// AssertNoMockExpectationsError checks that all mock expectations were met
func AssertNoMockExpectationsError(t *testing.T, mock sqlmock.Sqlmock) {
	if err := mock.ExpectationsWereMet(); err != nil {
		t.Errorf("Mock expectations were not met: %v", err)
	}
}

// TestAccountData provides common test data for account testing
type TestAccountData struct {
	ValidCustomer *models.Account
	ValidAdmin    *models.Account
	ValidOwner    *models.Account
	ValidStaff    *models.Account
	InvalidRole   *models.Account
	EmptyEmail    *models.Account
}

// GetTestAccountData returns a set of test account data
func GetTestAccountData() *TestAccountData {
	return &TestAccountData{
		ValidCustomer: &models.Account{
			Name:     "John Customer",
			Email:    "customer@test.com",
			Phone:    "1234567890",
			Password: "password123",
			Role:     "customer",
		},
		ValidAdmin: &models.Account{
			Name:     "Admin User",
			Email:    "admin@test.com",
			Phone:    "0987654321",
			Password: "adminpass123",
			Role:     "admin",
		},
		ValidOwner: &models.Account{
			Name:     "Restaurant Owner",
			Email:    "owner@test.com",
			Phone:    "1122334455",
			Password: "ownerpass123",
			Role:     "owner",
		},
		ValidStaff: &models.Account{
			Name:      "Staff Member",
			Email:     "staff@test.com",
			Phone:     "5566778899",
			Password:  "staffpass123",
			Role:      "staff",
			Orther_id: 1,
		},
		InvalidRole: &models.Account{
			Name:     "Invalid User",
			Email:    "invalid@test.com",
			Phone:    "1111111111",
			Password: "password123",
			Role:     "invalid_role",
		},
		EmptyEmail: &models.Account{
			Name:     "No Email User",
			Email:    "",
			Phone:    "2222222222",
			Password: "password123",
			Role:     "customer",
		},
	}
}
