package repositories

import (
	"database/sql"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/restaurent_table_booking/internal/db"
	"github.com/restaurent_table_booking/internal/models"
	"github.com/stretchr/testify/assert"
)

func TestCheckAccount(t *testing.T) {
	// Create a mock database
	mockDB, mock, err := sqlmock.New()
	assert.NoError(t, err)
	defer mockDB.Close()

	// Replace the global DB with our mock
	originalDB := db.DB
	db.DB = mockDB
	defer func() { db.DB = originalDB }()

	tests := []struct {
		name           string
		account        *models.Account
		mockSetup      func()
		expectedPass   string
		expectedExists bool
		expectedError  bool
	}{
		{
			name: "Customer account exists",
			account: &models.Account{
				Email: "customer@test.com",
			},
			mockSetup: func() {
				mock.ExpectQuery("SELECT id, password FROM customers WHERE gmail = ?").
					WithArgs("customer@test.com").
					WillReturnRows(sqlmock.NewRows([]string{"id", "password"}).
						AddRow(1, "hashedpassword"))
			},
			expectedPass:   "hashedpassword",
			expectedExists: false,
			expectedError:  false,
		},
		{
			name: "Staff account exists",
			account: &models.Account{
				Email: "staff@test.com",
			},
			mockSetup: func() {
				mock.ExpectQuery("SELECT id, password FROM customers WHERE gmail = ?").
					WithArgs("staff@test.com").
					WillReturnError(sql.ErrNoRows)
				mock.ExpectQuery("SELECT id, password FROM staffs WHERE gmail = ?").
					WithArgs("staff@test.com").
					WillReturnRows(sqlmock.NewRows([]string{"id", "password"}).
						AddRow(2, "staffpassword"))
			},
			expectedPass:   "staffpassword",
			expectedExists: false,
			expectedError:  false,
		},
		{
			name: "Account does not exist",
			account: &models.Account{
				Email: "nonexistent@test.com",
			},
			mockSetup: func() {
				mock.ExpectQuery("SELECT id, password FROM customers WHERE gmail = ?").
					WithArgs("nonexistent@test.com").
					WillReturnError(sql.ErrNoRows)
				mock.ExpectQuery("SELECT id, password FROM staffs WHERE gmail = ?").
					WithArgs("nonexistent@test.com").
					WillReturnError(sql.ErrNoRows)
				mock.ExpectQuery("SELECT id, password FROM admin WHERE gmail = ?").
					WithArgs("nonexistent@test.com").
					WillReturnError(sql.ErrNoRows)
				mock.ExpectQuery("SELECT id, password FROM owners WHERE gmail = ?").
					WithArgs("nonexistent@test.com").
					WillReturnError(sql.ErrNoRows)
			},
			expectedPass:   "",
			expectedExists: true,
			expectedError:  false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.mockSetup()

			password, exists, err := CheckAccount(tt.account)

			if tt.expectedError {
				assert.Error(t, err)
			} else {
				assert.NoError(t, err)
				assert.Equal(t, tt.expectedPass, password)
				assert.Equal(t, tt.expectedExists, exists)
			}

			assert.NoError(t, mock.ExpectationsWereMet())
		})
	}
}

func TestInsertAccount(t *testing.T) {
	mockDB, mock, err := sqlmock.New()
	assert.NoError(t, err)
	defer mockDB.Close()

	originalDB := db.DB
	db.DB = mockDB
	defer func() { db.DB = originalDB }()

	tests := []struct {
		name          string
		table         string
		values        []interface{}
		mockSetup     func()
		expectedID    int64
		expectedError bool
	}{
		{
			name:   "Successful customer insert",
			table:  "customers(name, gmail, phone, password, status)",
			values: []interface{}{"John Doe", "john@test.com", "1234567890", "hashedpass", "active"},
			mockSetup: func() {
				mock.ExpectPrepare("INSERT INTO customers\\(name, gmail, phone, password, status\\) VALUES \\(\\?,\\?,\\?,\\?,\\?\\)").
					ExpectExec().
					WithArgs("John Doe", "john@test.com", "1234567890", "hashedpass", "active").
					WillReturnResult(sqlmock.NewResult(123, 1))
			},
			expectedID:    123,
			expectedError: false,
		},
		{
			name:   "Database error",
			table:  "customers(name, gmail, phone, password, status)",
			values: []interface{}{"John Doe", "john@test.com", "1234567890", "hashedpass", "active"},
			mockSetup: func() {
				mock.ExpectPrepare("INSERT INTO customers\\(name, gmail, phone, password, status\\) VALUES \\(\\?,\\?,\\?,\\?,\\?\\)").
					ExpectExec().
					WithArgs("John Doe", "john@test.com", "1234567890", "hashedpass", "active").
					WillReturnError(sql.ErrConnDone)
			},
			expectedID:    0,
			expectedError: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.mockSetup()

			id, err := InsertAccount(tt.table, tt.values...)

			if tt.expectedError {
				assert.Error(t, err)
			} else {
				assert.NoError(t, err)
				assert.Equal(t, tt.expectedID, id)
			}

			assert.NoError(t, mock.ExpectationsWereMet())
		})
	}
}

func TestGetAllAccounts(t *testing.T) {
	mockDB, mock, err := sqlmock.New()
	assert.NoError(t, err)
	defer mockDB.Close()

	originalDB := db.DB
	db.DB = mockDB
	defer func() { db.DB = originalDB }()

	t.Run("Successful retrieval", func(t *testing.T) {
		rows := sqlmock.NewRows([]string{"id", "gmail", "name", "phone", "password"}).
			AddRow(1, "user1@test.com", "User One", "1111111111", "pass1").
			AddRow(2, "admin@test.com", "Admin User", "2222222222", "pass2")

		mock.ExpectQuery("SELECT \\* FROM users UNION SELECT \\* FROM admin").
			WillReturnRows(rows)

		accounts, err := GetAllAccounts()

		assert.NoError(t, err)
		assert.Len(t, accounts, 2)
		assert.Equal(t, "user1@test.com", accounts[0].Email)
		assert.Equal(t, "admin@test.com", accounts[1].Email)
		assert.NoError(t, mock.ExpectationsWereMet())
	})

	t.Run("Database error", func(t *testing.T) {
		mock.ExpectQuery("SELECT \\* FROM users UNION SELECT \\* FROM admin").
			WillReturnError(sql.ErrConnDone)

		accounts, err := GetAllAccounts()

		assert.Error(t, err)
		assert.Nil(t, accounts)
		assert.NoError(t, mock.ExpectationsWereMet())
	})
}
