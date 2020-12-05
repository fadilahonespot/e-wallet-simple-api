package usecase_test

import (
	"e-wallet-simple-api/customer/mocks"
	"e-wallet-simple-api/model"
	usecase "e-wallet-simple-api/customer/usecase"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
)

var mockAccount = model.Account{
	AccountNumber: "879999",
	CustomerName:  "Fadilah",
	Balance:       10000,
}

func TestCostomerUsecaseImpl_FindCustomerByID(t *testing.T) {
	t.Run("Shoul return data", func(t *testing.T) {

		mockAccountRepo := new(mocks.CustomerRepoMock)
		mockAccountRepo.On("FindCustomerByID", mockAccount.AccountNumber).Return(&mockAccount, nil)

		customerUsecase := usecase.CreateCustomerUsecase(mockAccountRepo)
		res, err := customerUsecase.FindCustomerByID(mockAccount.AccountNumber)
		if res == nil {
			t.Fatalf("should return data but return nil")
		}
		if res.AccountNumber != mockAccount.AccountNumber {
			t.Fatalf("Account ID should return %v, got %v", mockAccount.AccountNumber, res.AccountNumber)
		}
		if err != nil {
			t.Fatalf("error return nil, got %v", err)
		}
	})
}

func TestCustomerUsecaseImpl_InsertCustomer(t *testing.T) {
	t.Run("Test Normal Case", func(t *testing.T) {
		db, mockSql, err := sqlmock.New()
		if err != nil {
			t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
		}
		defer db.Close()
		mockSql.ExpectBegin()

		mockAccoundRepo := new(mocks.CustomerRepoMock)
		mockAccoundRepo.On("InsertCustomer", &mockAccount).Return(nil)

		u := usecase.CreateCustomerUsecase(mockAccoundRepo)
		isSucces := u.InsertCustomer(&mockAccount)

		if isSucces != true {
			t.Errorf("Error should be false, but get %v", isSucces)
		}
	})
}

func TestCustomerUsecaseImpl_CheckAccoutExist(t *testing.T) {
	t.Run("Test Normal Case", func(t *testing.T) {

		mockAccountRepo := new(mocks.CustomerRepoMock)
		mockAccountRepo.On("CheckAccoutExist", mockAccount.AccountNumber).Return(nil)

		customerUsecase := usecase.CreateCustomerUsecase(mockAccountRepo)
		isExist := customerUsecase.CheckAccoutExist(mockAccount.AccountNumber)
		if isExist != true {
			t.Errorf("Error should be false, but get %v", isExist)
		}
	})
}

func TestCustomerUsecaseImpl_CheckCustomerExist(t *testing.T) {
	t.Run("Test Normal Case", func(t *testing.T) {
		mockAccountRepo := new(mocks.CustomerRepoMock)
		customerNumber := "10001"
		mockAccountRepo.On("CheckCustomerExist", customerNumber).Return(nil)

		customerUsecase := usecase.CreateCustomerUsecase(mockAccountRepo)
		isExist := customerUsecase.CheckCustomerExist(customerNumber)
		if isExist != true {
			t.Errorf("Error should be false, but get %v", isExist)
		}
	})
}

func TestCustomerUsecaseImpl_TransferBalance(t *testing.T) {
	t.Run("Test Normal Case", func(t *testing.T) {
		mockAccountRepo := new(mocks.CustomerRepoMock)
		var balance = model.Transfer{
			CostomerNumber:  "10004",
			ToAccountNumber: mockAccount.AccountNumber,
			Amount:          mockAccount.Balance,
		}
		mockAccountRepo.On("TransferBalance", &balance).Return(nil)

		customerUsecase := usecase.CreateCustomerUsecase(mockAccountRepo)
		isSucces := customerUsecase.TransferBalance(&balance)
		if isSucces != true {
			t.Errorf("Error should be false, but get %v", isSucces)
		}
	})
}

func TestCustomerUsecaseImpl_FindCustomers(t *testing.T) {
	t.Run("Shoul return data", func(t *testing.T) {

		mockAccountRepo := new(mocks.CustomerRepoMock)
		var mockAccounds []model.Account
		mockAccounds = append(mockAccounds, mockAccount)
		mockAccountRepo.On("FindCustomers").Return(mockAccounds, nil)

		customerUsecase := usecase.CreateCustomerUsecase(mockAccountRepo)
		res, err := customerUsecase.FindCustomers()
		if res == nil {
			t.Fatalf("should return data but return nil")
		}
		if len(*res) != len(mockAccounds) {
			t.Fatalf("Len should return %v, got %v", len(mockAccounds), len(*res))
		}
		if err != nil {
			t.Fatalf("error return nil, got %v", err)
		}
	})
} 
