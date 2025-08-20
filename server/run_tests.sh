#!/bin/bash

# Restaurant Table Booking - Test Runner Script
# This script runs all unit tests and integration tests for the account system

echo "🧪 Running Restaurant Table Booking Tests"
echo "=========================================="

# Set test environment
export GIN_MODE=test
export GO111MODULE=on

# Colors for output
RED='\033[0;31m'
GREEN='\033[0;32m'
YELLOW='\033[1;33m'
BLUE='\033[0;34m'
NC='\033[0m' # No Color

# Function to print colored output
print_status() {
    echo -e "${BLUE}[INFO]${NC} $1"
}

print_success() {
    echo -e "${GREEN}[SUCCESS]${NC} $1"
}

print_warning() {
    echo -e "${YELLOW}[WARNING]${NC} $1"
}

print_error() {
    echo -e "${RED}[ERROR]${NC} $1"
}

# Check if we're in the correct directory
if [ ! -f "go.mod" ]; then
    print_error "go.mod not found. Please run this script from the server directory."
    exit 1
fi

print_status "Downloading dependencies..."
go mod download
if [ $? -ne 0 ]; then
    print_error "Failed to download dependencies"
    exit 1
fi

print_status "Tidying go modules..."
go mod tidy

echo ""
print_status "Running Unit Tests..."
echo "----------------------"

# Run repository tests
print_status "Testing Account Repository..."
go test -v ./internal/repositories -run TestCheckAccount
go test -v ./internal/repositories -run TestInsertAccount  
go test -v ./internal/repositories -run TestGetAllAccounts

# Run service tests
print_status "Testing Account Services..."
go test -v ./internal/services -run TestRegister
go test -v ./internal/services -run TestLogin
go test -v ./internal/services -run TestGetAllAccounts

# Run controller tests
print_status "Testing Account Controllers..."
go test -v ./internal/controllers -run TestRegisterHandler
go test -v ./internal/controllers -run TestLoginHandler
go test -v ./internal/controllers -run TestLogoutHandler
go test -v ./internal/controllers -run TestGetProfile

echo ""
print_status "Running Integration Tests..."
echo "-----------------------------"

# Run integration tests
go test -v ./internal/tests -run TestAccountRegistrationFlow
go test -v ./internal/tests -run TestLoginLogoutFlow
go test -v ./internal/tests -run TestCompleteUserJourney

echo ""
print_status "Running All Tests with Coverage..."
echo "-----------------------------------"

# Run all tests with coverage
go test -v -cover ./internal/repositories ./internal/services ./internal/controllers ./internal/tests

echo ""
print_status "Generating Coverage Report..."
echo "------------------------------"

# Generate coverage report
go test -coverprofile=coverage.out ./internal/repositories ./internal/services ./internal/controllers ./internal/tests
if [ $? -eq 0 ]; then
    go tool cover -html=coverage.out -o coverage.html
    print_success "Coverage report generated: coverage.html"
else
    print_warning "Coverage report generation failed"
fi

echo ""
print_status "Running Race Condition Tests..."
echo "--------------------------------"

# Test for race conditions
go test -race ./internal/repositories ./internal/services ./internal/controllers ./internal/tests

echo ""
print_success "All tests completed!"
print_status "Check coverage.html for detailed coverage report"

# Clean up
unset GIN_MODE
unset GO111MODULE
