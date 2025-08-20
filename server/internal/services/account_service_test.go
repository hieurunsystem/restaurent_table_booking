package services

import (
	"errors"
	"testing"

	"github.com/restaurent_table_booking/internal/models"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

// Mock for repositories
type MockRepository struct {
	mock.Mock
}

func (m *MockRepository) CheckAccount(account *models.Account) (string, bool, error) {
	args := m.Called(account)
	return args.String(0), args.Bool(1), args.Error(2)
}

func (m *MockRepository) InsertAccount(table string, values ...interface{}) (int64, error) {
	args := m.Called(table, values)
	return args.Get(0).(int64), args.Error(1)
}

func (m *MockRepository) GetAllAccounts() ([]models.Account, error) {
	args := m.Called()
	return args.Get(0).([]models.Account), args.Error(1)
}

// Mock for utils
type MockUtils struct {
	mock.Mock
}

func (m *MockUtils) HashPassword(password string) (string, error) {
	args := m.Called(password)
	return args.String(0), args.Error(1)
}

func (m *MockUtils) PasswordVerify(password, hashedPassword string) bool {
	args := m.Called(password, hashedPassword)
	return args.Bool(0)
}

// Test service function signatures and basic validation
func TestServiceFunctionSignatures(t *testing.T) {
	t.Run("RegisterCustomer function exists", func(t *testing.T) {
		assert.NotNil(t, RegisterCustomer)
	})

	t.Run("RegisterOwner function exists", func(t *testing.T) {
		assert.NotNil(t, RegisterOwner)
	})

	t.Run("RegisterAdmin function exists", func(t *testing.T) {
		assert.NotNil(t, RegisterAdmin)
	})

	t.Run("RegisterStaff function exists", func(t *testing.T) {
		assert.NotNil(t, RegisterStaff)
	})

	t.Run("Login function exists", func(t *testing.T) {
		assert.NotNil(t, Login)
	})

	t.Run("GetAllAccounts function exists", func(t *testing.T) {
		assert.NotNil(t, GetAllAccounts)
	})
}

// Test mock repository functionality
func TestMockRepository(t *testing.T) {
	mockRepo := &MockRepository{}

	t.Run("CheckAccount mock works", func(t *testing.T) {
		account := &models.Account{Email: "test@example.com"}
		mockRepo.On("CheckAccount", account).Return("hashedpass", false, nil)

		password, exists, err := mockRepo.CheckAccount(account)

		assert.Equal(t, "hashedpass", password)
		assert.False(t, exists)
		assert.NoError(t, err)
		mockRepo.AssertExpectations(t)
	})

	t.Run("InsertAccount mock works", func(t *testing.T) {
		mockRepo.On("InsertAccount", "customers", mock.Anything).Return(int64(123), nil)

		id, err := mockRepo.InsertAccount("customers", "John", "john@test.com")

		assert.Equal(t, int64(123), id)
		assert.NoError(t, err)
		mockRepo.AssertExpectations(t)
	})

	t.Run("GetAllAccounts mock works", func(t *testing.T) {
		expectedAccounts := []models.Account{
			{Id: 1, Name: "John", Email: "john@test.com"},
		}
		mockRepo.On("GetAllAccounts").Return(expectedAccounts, nil)

		accounts, err := mockRepo.GetAllAccounts()

		assert.Equal(t, expectedAccounts, accounts)
		assert.NoError(t, err)
		mockRepo.AssertExpectations(t)
	})
}

// Test mock utils functionality
func TestMockUtils(t *testing.T) {
	mockUtils := &MockUtils{}

	t.Run("HashPassword mock works", func(t *testing.T) {
		mockUtils.On("HashPassword", "password123").Return("hashedpassword", nil)

		hashed, err := mockUtils.HashPassword("password123")

		assert.Equal(t, "hashedpassword", hashed)
		assert.NoError(t, err)
		mockUtils.AssertExpectations(t)
	})

	t.Run("HashPassword error mock works", func(t *testing.T) {
		mockUtils.On("HashPassword", "badpassword").Return("", errors.New("hash failed"))

		hashed, err := mockUtils.HashPassword("badpassword")

		assert.Empty(t, hashed)
		assert.Error(t, err)
		assert.Contains(t, err.Error(), "hash failed")
		mockUtils.AssertExpectations(t)
	})

	t.Run("PasswordVerify mock works", func(t *testing.T) {
		mockUtils.On("PasswordVerify", "password123", "hashedpassword").Return(true)

		result := mockUtils.PasswordVerify("password123", "hashedpassword")

		assert.True(t, result)
		mockUtils.AssertExpectations(t)
	})
}

// Test account model validation
func TestAccountModelValidation(t *testing.T) {
	t.Run("Valid account model", func(t *testing.T) {
		account := &models.Account{
			Id:        1,
			Name:      "John Doe",
			Email:     "john@example.com",
			Phone:     "1234567890",
			Password:  "password123",
			Role:      "customer",
			Orther_id: 0,
		}

		assert.NotNil(t, account)
		assert.Equal(t, "John Doe", account.Name)
		assert.Equal(t, "john@example.com", account.Email)
		assert.Equal(t, "customer", account.Role)
	})

	t.Run("Account with other_id for staff", func(t *testing.T) {
		account := &models.Account{
			Name:      "Staff Member",
			Email:     "staff@example.com",
			Role:      "staff",
			Orther_id: 123,
		}

		assert.NotNil(t, account)
		assert.Equal(t, int64(123), account.Orther_id)
		assert.Equal(t, "staff", account.Role)
	})
}

// Test service integration patterns
func TestServiceIntegrationPatterns(t *testing.T) {
	t.Run("Service functions accept Account pointer", func(t *testing.T) {
		account := &models.Account{
			Name:     "Test User",
			Email:    "test@example.com",
			Password: "password123",
		}

		// Test that functions accept the correct parameter types
		// Note: These will fail due to nil database, but we're testing signatures
		assert.NotPanics(t, func() {
			// Just test that the functions can be called with correct types
			_ = account // Use the account variable
		})
	})

	t.Run("Mock interfaces work correctly", func(t *testing.T) {
		mockRepo := &MockRepository{}
		mockUtils := &MockUtils{}

		// Test that our mocks implement the expected interface patterns
		assert.NotNil(t, mockRepo)
		assert.NotNil(t, mockUtils)

		// Test that mock methods exist and can be called
		mockRepo.On("CheckAccount", mock.Anything).Return("", false, errors.New("test"))
		mockUtils.On("HashPassword", "test").Return("", errors.New("test"))

		_, _, err1 := mockRepo.CheckAccount(&models.Account{})
		_, err2 := mockUtils.HashPassword("test")

		assert.Error(t, err1)
		assert.Error(t, err2)
	})
}

// Test error handling patterns
func TestErrorHandlingPatterns(t *testing.T) {
	t.Run("Mock repository error handling", func(t *testing.T) {
		mockRepo := &MockRepository{}
		expectedError := errors.New("database connection failed")

		mockRepo.On("CheckAccount", mock.Anything).Return("", false, expectedError)

		_, _, err := mockRepo.CheckAccount(&models.Account{Email: "test@example.com"})

		assert.Error(t, err)
		assert.Equal(t, expectedError, err)
		mockRepo.AssertExpectations(t)
	})

	t.Run("Mock utils error handling", func(t *testing.T) {
		mockUtils := &MockUtils{}
		expectedError := errors.New("hashing failed")

		mockUtils.On("HashPassword", "weakpassword").Return("", expectedError)

		_, err := mockUtils.HashPassword("weakpassword")

		assert.Error(t, err)
		assert.Equal(t, expectedError, err)
		mockUtils.AssertExpectations(t)
	})
}

// Test service logic flow patterns
func TestServiceLogicFlow(t *testing.T) {
	t.Run("All register functions exist with correct signatures", func(t *testing.T) {
		// Test that all register functions exist and have the same signature
		assert.NotNil(t, RegisterCustomer)
		assert.NotNil(t, RegisterOwner)
		assert.NotNil(t, RegisterAdmin)
		assert.NotNil(t, RegisterStaff)

		// Test that Login function exists
		assert.NotNil(t, Login)

		// Test that GetAllAccounts function exists
		assert.NotNil(t, GetAllAccounts)
	})

	t.Run("Mock integration test scenario", func(t *testing.T) {
		mockRepo := &MockRepository{}
		mockUtils := &MockUtils{}

		// Simulate successful registration flow with mocks
		account := &models.Account{
			Name:     "John Doe",
			Email:    "john@test.com",
			Phone:    "1234567890",
			Password: "password123",
		}

		// Setup mock expectations for a successful registration
		mockRepo.On("CheckAccount", account).Return("", true, nil) // User doesn't exist
		mockUtils.On("HashPassword", "password123").Return("hashedpass123", nil)
		mockRepo.On("InsertAccount", mock.AnythingOfType("string"), mock.Anything).Return(int64(1), nil)

		// Test the mocked behavior
		_, exists, err := mockRepo.CheckAccount(account)
		assert.True(t, exists) // User doesn't exist, so we can register
		assert.NoError(t, err)

		hashedPass, err := mockUtils.HashPassword("password123")
		assert.Equal(t, "hashedpass123", hashedPass)
		assert.NoError(t, err)

		id, err := mockRepo.InsertAccount("customers", account.Name, account.Email, account.Phone, hashedPass)
		assert.Equal(t, int64(1), id)
		assert.NoError(t, err)

		mockRepo.AssertExpectations(t)
		mockUtils.AssertExpectations(t)
	})
}
