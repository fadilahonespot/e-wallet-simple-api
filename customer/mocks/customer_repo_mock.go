package mocks

import (
	"e-wallet-simple-api/model"
	"github.com/stretchr/testify/mock"
)


type CustomerRepoMock struct {
	mock.Mock // Type Embedding
}

func (m *CustomerRepoMock) FindCustomerByID(accountNumber string) (*model.Account, error) {
	ret := m.Called(accountNumber)

	var mAccount *model.Account
	if ret.Get(0) != nil {
		mAccount = ret.Get(0).(*model.Account) // assert type
	}

	return mAccount, ret.Error(1)
}

func (m *CustomerRepoMock) InsertCustomer(customer *model.Account) bool {
	ret := m.Called(customer)

	err := ret.Error(0)
	if err != nil {
		return false
	}
	return true
}

func (m *CustomerRepoMock) CheckAccoutExist(accountNumber string) bool {
	ret := m.Called(accountNumber)

	err := ret.Error(0)
	if err != nil {
		return false
	}
	return true
}

func (m *CustomerRepoMock) CheckCustomerExist(customerNumber string) bool {
	ret := m.Called(customerNumber)

	err := ret.Error(0)
	if err != nil {
		return false
	}
	return true
}

func (m *CustomerRepoMock) TransferBalance(balance *model.Transfer) bool {
	ret := m.Called(balance)

	err := ret.Error(0)
	if err != nil {
		return false
	}
	return true
}

func (m *CustomerRepoMock) FindCustomers() (*[]model.Account, error) {
	ret := m.Called()

	var r0 []model.Account
	if rf, ok := ret.Get(0).(func() []model.Account); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]model.Account)
		}
	}

	var r2 error
	if rf, ok := ret.Get(1).(func() error); ok {
		r2 = rf()
	} else {
		r2 = ret.Error(1)
	}

	return &r0, r2
}
