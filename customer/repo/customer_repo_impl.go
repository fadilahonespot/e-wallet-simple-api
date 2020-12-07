package repo

import (
	"database/sql"
	"e-wallet-simple-api/constant"
	"e-wallet-simple-api/customer"
	"e-wallet-simple-api/model"
	"fmt"
)


type CustomerRepoImpl struct {
	db *sql.DB
}

func CreateCustomerRepo(db *sql.DB) customer.CustomerRepo {
	return &CustomerRepoImpl{db}
}

func (e *CustomerRepoImpl) FindCustomers() (*[]model.Account, error) {
	sql := "SELECT * FROM get_customers()"
	res, err := e.db.Query(sql)
	if err != nil {
		fmt.Printf("[CustomerRepoImpl.findCustomers] error execute query %v \n", err)
		return nil, fmt.Errorf(constant.ServerHasWrong)
	}
	var account []model.Account
	for res.Next() {
		var data model.Account
		err := res.Scan(
			&data.AccountNumber,
			&data.CustomerName,
			&data.Balance,
		)
		if err != nil {
			fmt.Printf("[CustomerRepoImpl.FindCustomers] error scan data %v \n", err)
			return nil, fmt.Errorf(constant.ServerHasWrong)
		}
		account = append(account, data)
	}
	defer res.Close()
	return &account, nil
}

func (e *CustomerRepoImpl) FindCustomerByID(accountNumber string) (*model.Account, error) {
	sql := "SELECT * FROM get_detail_customer($1)"
	res, err := e.db.Query(sql, accountNumber)
	if err != nil {
		fmt.Printf("[CustomerRepoImpl.FindCustomerByID] error execute query %v \n", err)
		return nil, fmt.Errorf(constant.ServerHasWrong)
	}
	var account model.Account
	for res.Next() {
		err := res.Scan(
			&account.AccountNumber,
			&account.CustomerName,
			&account.Balance,
		)
		if err != nil {
			fmt.Printf("[CustomerRepoImpl.FindCustomerById] error scan data %v \n", err)
			return nil, fmt.Errorf(constant.ServerHasWrong)
		} 
	}
	defer res.Close()
	return &account, nil
} 

func (e *CustomerRepoImpl) InsertCustomer(customer *model.Account) bool {
	sql := "SELECT * FROM add_customer($1, $2)"
	res, err := e.db.Query(sql, customer.CustomerName, customer.Balance)
	if err != nil {
		fmt.Printf("[CustomerRepoImpl.Insertcustomer] error execute query %v \n", err)
		return false
	}
	defer res.Close()
	return true
}

func (e *CustomerRepoImpl) CheckAccoutExist(accountNumber string) bool {
	sql := "SELECT * FROM is_account_exist($1)"
	res, err := e.db.Query(sql, accountNumber)
	if err != nil {
		fmt.Printf("[CustomerrepoImpl.CheckAccountExist] error execute query %v \n", err)
		return false
	}
	var isValid bool
	for res.Next() {
		err := res.Scan(&isValid)
		if err != nil {
			fmt.Printf("[CustomerRepoImpl.checkAccountExist] error scan %v \n", err)
			return false
		}
	}
	defer res.Close()
	return isValid
}

func (e *CustomerRepoImpl) CheckCustomerExist(customerNumber string) bool {
	sql := "SELECT * FROM is_customer_exist($1)"
	res, err := e.db.Query(sql, customerNumber)
	if err != nil {
		fmt.Printf("[CustomerRepoImpl.checkCustomertExist] error scan %v \n", err)
		return false
	}
	var isValid bool
	for res.Next() {
		err := res.Scan(&isValid)
		if err != nil {
			fmt.Printf("[CustomerRepoImpl.checkCustomerExist] error scan %v \n", err)
			return false
		}
	}
	defer res.Close()
	return isValid
}

func (e *CustomerRepoImpl) TransferBalance(balance *model.Transfer) bool {
	sql := "SELECT * FROM transfer_balance($1, $2, $3)"
	res, err := e.db.Query(sql, balance.MyAccountNumber, balance.ToAccountNumber, balance.Amount)
	if err != nil {
		fmt.Printf("[CustomerRepoImpl.TransferBalance] error execute query %v \n", err)
		return false
	}
	var isSuccess bool
	for res.Next() {
		err := res.Scan(&isSuccess)
		if err != nil {
			fmt.Printf("[CustomerRepoImpl.TrasferBalance] error scan data %v \n", err)
			return false
		}
	}
	defer res.Close()
	return isSuccess
}