package usecase

import (
	"e-wallet-simple-api/customer"	
	"e-wallet-simple-api/model"
)


type CostomerUsecaseImpl struct {
	customerRepo customer.CustomerRepo
}

func CreateCustomerUsecase(customerRepo customer.CustomerRepo) customer.CustomerUsecase {
	return &CostomerUsecaseImpl{customerRepo}
}

func (e *CostomerUsecaseImpl) FindCustomerByID(accountNumber string) (*model.Account, error) {
	return e.customerRepo.FindCustomerByID(accountNumber)
}

func (e *CostomerUsecaseImpl) InsertCustomer(customer *model.Account) bool {
	return e.customerRepo.InsertCustomer(customer)
}

func (e *CostomerUsecaseImpl) TransferBalance(balance *model.Transfer) bool {
	return e.customerRepo.TransferBalance(balance)
}

func (e *CostomerUsecaseImpl) CheckAccoutExist(accountNumber string) bool {
	return e.customerRepo.CheckAccoutExist(accountNumber)
}

func (e *CostomerUsecaseImpl) CheckCustomerExist(customerNumber string) bool {
	return e.customerRepo.CheckCustomerExist(customerNumber)
}

func (e *CostomerUsecaseImpl) FindCustomers() (*[]model.Account, error) {
	return e.customerRepo.FindCustomers()
}