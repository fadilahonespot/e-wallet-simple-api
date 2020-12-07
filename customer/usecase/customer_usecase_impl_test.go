package usecase_test

import (
	"e-wallet-simple-api/customer/mocks"
	usecase "e-wallet-simple-api/customer/usecase"
	"e-wallet-simple-api/model"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/stretchr/testify/assert"
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
		
		assert.NotNil(t, res)
		assert.Equal(t, mockAccount.AccountNumber, res.AccountNumber)
		assert.Nil(t, err)
	})
}

func TestCustomerUsecaseImpl_InsertCustomer(t *testing.T) {
	t.Run("Test Normal Case", func(t *testing.T) {
		db, mockSql, err := sqlmock.New()
		assert.Nil(t, err)
		defer db.Close()
		mockSql.ExpectBegin()

		mockAccoundRepo := new(mocks.CustomerRepoMock)
		mockAccoundRepo.On("InsertCustomer", &mockAccount).Return(nil)

		u := usecase.CreateCustomerUsecase(mockAccoundRepo)
		isSucces := u.InsertCustomer(&mockAccount)

		assert.Equal(t, true, isSucces)
	})
}

func TestCustomerUsecaseImpl_CheckAccoutExist(t *testing.T) {
	t.Run("Test Normal Case", func(t *testing.T) {

		mockAccountRepo := new(mocks.CustomerRepoMock)
		mockAccountRepo.On("CheckAccoutExist", mockAccount.AccountNumber).Return(nil)

		customerUsecase := usecase.CreateCustomerUsecase(mockAccountRepo)
		isExist := customerUsecase.CheckAccoutExist(mockAccount.AccountNumber)
		
		assert.Equal(t, true, isExist)
	})
}

func TestCustomerUsecaseImpl_CheckCustomerExist(t *testing.T) {
	t.Run("Test Normal Case", func(t *testing.T) {
		mockAccountRepo := new(mocks.CustomerRepoMock)
		customerNumber := "10001"
		mockAccountRepo.On("CheckCustomerExist", customerNumber).Return(nil)

		customerUsecase := usecase.CreateCustomerUsecase(mockAccountRepo)
		isExist := customerUsecase.CheckCustomerExist(customerNumber)
		
		assert.Equal(t, true, isExist)
	})
}

func TestCustomerUsecaseImpl_TransferBalance(t *testing.T) {
	t.Run("Test Normal Case", func(t *testing.T) {
		mockAccountRepo := new(mocks.CustomerRepoMock)
		var balance = model.Transfer{
			MyAccountNumber: "10004",
			ToAccountNumber: mockAccount.AccountNumber,
			Amount:          mockAccount.Balance,
		}
		mockAccountRepo.On("TransferBalance", &balance).Return(nil)

		customerUsecase := usecase.CreateCustomerUsecase(mockAccountRepo)
		isSucces := customerUsecase.TransferBalance(&balance)
		
		assert.Equal(t, true, isSucces)
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

		assert.NotNil(t, res)
		assert.Equal(t, len(mockAccounds), len(*res))
		assert.Nil(t, err)
	})
}
