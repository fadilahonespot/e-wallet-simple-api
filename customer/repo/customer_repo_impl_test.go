package repo_test

import (
	"database/sql"
	customerRepoTest "e-wallet-simple-api/customer/repo"
	"e-wallet-simple-api/model"
	"log"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/stretchr/testify/assert"
)

var u = model.Account{
	AccountNumber: "555004",
	CustomerName:  "Asep Saripudin",
	Balance:       1000,
}

func NewMock() (*sql.DB, sqlmock.Sqlmock) {
	db, mock, err := sqlmock.New(sqlmock.QueryMatcherOption(sqlmock.QueryMatcherEqual))
	if err != nil {
		log.Fatalf("an error '%s' was not expected when opening a stub database connect", err)
	}
	return db, mock
}

func TestCustomerRepoImpl_FindCustomers(t *testing.T) {
	db, mock := NewMock()
	rows := sqlmock.NewRows([]string{"account_number", "name", "balance"}).
		AddRow(u.AccountNumber, u.CustomerName, u.Balance)

	query := "SELECT * FROM get_customers()"

	mock.ExpectQuery(query).WillReturnRows(rows)
	defer db.Close()

	repo := customerRepoTest.CreateCustomerRepo(db)

	user, err := repo.FindCustomers()
	assert.NotEmpty(t, user)
	assert.NoError(t, err)
	assert.Len(t, *user, 1)
}

func TestCustomerRepoImpl_FindCustomerByID(t *testing.T) {
	db, mock := NewMock()
	rows := sqlmock.NewRows([]string{"account_number", "name", "balance"}).
		AddRow(u.AccountNumber, u.CustomerName, u.Balance)

	query := "SELECT * FROM get_detail_customer($1)"

	mock.ExpectQuery(query).WithArgs(u.AccountNumber).WillReturnRows(rows)
	defer db.Close()

	repo := customerRepoTest.CreateCustomerRepo(db)

	user, err := repo.FindCustomerByID(u.AccountNumber)
	
	assert.NotNil(t, user)
	assert.NoError(t, err)
}

func TestCustomerRepoImpl_FindCustomerByIDError(t *testing.T) {
	db, mock := NewMock()
	rows := sqlmock.NewRows([]string{"account_number", "name", "balance"})

	query := "SELECT * FROM get_detail_customer($1)"

	mock.ExpectQuery(query).WithArgs(u.AccountNumber).WillReturnRows(rows)
	defer db.Close()

	repo := customerRepoTest.CreateCustomerRepo(db)

	user, err := repo.FindCustomerByID(u.AccountNumber)
	if (err != nil) != false {
		t.Errorf("CostomerRepoImpl.FindCustomerByID() error = %v, wantErr %v", err, false)
		return
	}
	assert.Empty(t, user)
}

func TestCustomerRepoImpl_InsertCustomer(t *testing.T) {
	db, mock := NewMock()

	rows := sqlmock.NewRows([]string{"isValid"}).
		AddRow("true")
	query := "SELECT * FROM add_customer($1, $2)"
	mock.ExpectQuery(query).WithArgs(u.CustomerName, u.Balance).WillReturnRows(rows)
	defer db.Close()

	repo := customerRepoTest.CreateCustomerRepo(db)

	isSuccess := repo.InsertCustomer(&u)
	assert.Equal(t, true, isSuccess)
}


func TestCustomerImpl_CheckAccoutExist(t *testing.T) {
	db, mock := NewMock()

	rows := sqlmock.NewRows([]string{"isValid"}).
		AddRow("true")
	query := "SELECT * FROM is_account_exist($1)"
	mock.ExpectQuery(query).WithArgs(u.AccountNumber).WillReturnRows(rows)
	defer db.Close()

	repo := customerRepoTest.CreateCustomerRepo(db)

	isSuccess := repo.CheckAccoutExist(u.AccountNumber)
	assert.Equal(t, true, isSuccess)
}

func TestCustomerImpl_CheckCustomerExist(t *testing.T) {
	db, mock := NewMock()

	var customerNumber = "1001"
	rows := sqlmock.NewRows([]string{"isValid"}).
		AddRow("true")
	query := "SELECT * FROM is_customer_exist($1)"
	mock.ExpectQuery(query).WithArgs(customerNumber).WillReturnRows(rows)
	defer db.Close()

	repo := customerRepoTest.CreateCustomerRepo(db)

	isSuccess := repo.CheckCustomerExist(customerNumber)
	assert.Equal(t, true, isSuccess)
}

func TestCustomerImpl_TransferBalance(t *testing.T) {
	db, mock := NewMock()

	var b = model.Transfer{
		MyAccountNumber:  "1001",
		ToAccountNumber: u.AccountNumber,
		Amount:          u.Balance,
	}
	rows := sqlmock.NewRows([]string{"isValid"}).
		AddRow("true")
	query := "SELECT * FROM transfer_balance($1, $2, $3)"
	mock.ExpectQuery(query).WithArgs(b.MyAccountNumber, b.ToAccountNumber, b.Amount).WillReturnRows(rows)
	defer db.Close()

	repo := customerRepoTest.CreateCustomerRepo(db)

	isSuccess := repo.TransferBalance(&b)
	assert.Equal(t, true, isSuccess)
}


