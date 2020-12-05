package handler_test

import (
	"e-wallet-simple-api/constant"
	"e-wallet-simple-api/customer/handler"
	"e-wallet-simple-api/customer/mocks"
	"e-wallet-simple-api/model"
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

var mockAccount = model.Account{
	AccountNumber: "555004",
	CustomerName:  "Fadilah",
	Balance:       10000,
}

var returnData = model.ResponWrapper{
	Success: true,
	Message: constant.Sucess,
}

var returnMessage= model.ResponWrapperMessage{
	Success: true,
} 

func TestCustomerHandler_getAccountDetail(t *testing.T) {
	t.Run("Test Normal Case", func(t *testing.T) {

		AccountMockUsecase := new(mocks.CustomerUsecaseMock)
		AccountMockUsecase.On("CheckAccoutExist", mock.AnythingOfType("string")).Return(nil)
		AccountMockUsecase.On("FindCustomerByID", mock.AnythingOfType("string")).Return(&mockAccount, nil)

		url := fmt.Sprintf("/account/%v", mockAccount.AccountNumber)
		req, err := http.NewRequest(http.MethodGet, url, nil)
		if err != nil {
			t.Fatalf("Error when creating request %v ", err)
		}
		resp := httptest.NewRecorder()
		router := gin.Default()
		handler.CreateCustomerHandler(router, AccountMockUsecase)
		router.ServeHTTP(resp, req)

		returnData.Data = mockAccount
		data, err := json.Marshal(&returnData)
		assert.NoError(t, err)
		responDat := string(data)

		assert.Equal(t, 200, resp.Code)
		assert.Equal(t, responDat, resp.Body.String())
	})
}

func TestCustomerHandler_getAcccounts(t *testing.T) {
	t.Run("Test Normal Case", func(t *testing.T) {

		AccountMockUsecase := new(mocks.CustomerUsecaseMock)
		var accounts []model.Account
		accounts = append(accounts, mockAccount)
		AccountMockUsecase.On("FindCustomers").Return(accounts, nil)

		url := "/account"
		req, err := http.NewRequest(http.MethodGet, url, nil)
		if err != nil {
			t.Fatalf("Error when creating request %v ", err)
		}
		resp := httptest.NewRecorder()
		router := gin.Default()
		handler.CreateCustomerHandler(router, AccountMockUsecase)
		router.ServeHTTP(resp, req)

		returnData.Data = accounts
		data, err := json.Marshal(&returnData)
		assert.NoError(t, err)
		responDat := string(data)

		assert.Equal(t, 200, resp.Code)
		assert.Equal(t, responDat, resp.Body.String())
	})
}

func TestCustomerhandler_insertCustomer(t *testing.T) {
	t.Run("Test Normal Case", func(t *testing.T) {

		AccountMockUsecase := new(mocks.CustomerUsecaseMock)
		AccountMockUsecase.On("InsertCustomer", mock.AnythingOfType("*model.Account")).Return(nil)

		d, err := json.Marshal(mockAccount)
		assert.NoError(t, err)

		url := "/account"
		req, err := http.NewRequest(http.MethodPost, url, strings.NewReader(string(d)))
		if err != nil {
			t.Fatalf("Error when creating request %v ", err)
		}
		resp := httptest.NewRecorder()
		router := gin.Default()
		handler.CreateCustomerHandler(router, AccountMockUsecase)
		router.ServeHTTP(resp, req)

		returnMessage.Message = constant.SuccessInputData
		data, err := json.Marshal(&returnMessage)
		assert.NoError(t, err)
		responDat := string(data)

		assert.Equal(t, 200, resp.Code)
		assert.Equal(t, responDat, resp.Body.String())
	})
}

func TestCustomerHandler_tranferBalance(t *testing.T) {
	t.Run("Test Normal Case", func(t *testing.T) {

		var customerNumber = "10002"
		var body = model.Transfer{
			ToAccountNumber: mockAccount.AccountNumber,
			Amount:          mockAccount.Balance,
		}

		AccountMockUsecase := new(mocks.CustomerUsecaseMock)
		AccountMockUsecase.On("CheckCustomerExist", mock.AnythingOfType("string")).Return(nil)
		AccountMockUsecase.On("CheckAccoutExist", mock.AnythingOfType("string")).Return(nil)
		AccountMockUsecase.On("TransferBalance", mock.AnythingOfType("*model.Transfer")).Return(nil)

		part_url := "/transfer"
		data, err := json.Marshal(&body)
		assert.NoError(t, err)

		req, err := http.NewRequest(http.MethodPut, part_url, strings.NewReader(string(data)))
		if err != nil {
			t.Fatalf("Error when creating request %v ", err)
		}
		req.Header.Set(constant.CustomerNumber, customerNumber)

		resp := httptest.NewRecorder()
		router := gin.Default()
		handler.CreateCustomerHandler(router, AccountMockUsecase)
		router.ServeHTTP(resp, req)

		returnMessage.Message = constant.SuccessTransferAmount
		output, err := json.Marshal(&returnMessage)
		assert.NoError(t, err)
		responDat := string(output)

		assert.Equal(t, 200, resp.Code)
		assert.Equal(t, responDat, resp.Body.String())
	})
}

