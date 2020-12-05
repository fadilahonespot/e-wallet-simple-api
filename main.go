package main

import (
	"log"
	"os"

	customerhandler "e-wallet-simple-api/customer/handler"
	customerRepo "e-wallet-simple-api/customer/repo"
	customerUsecase "e-wallet-simple-api/customer/usecase"
	"e-wallet-simple-api/setting"

	"github.com/gin-gonic/gin"
	_ "github.com/joho/godotenv/autoload"
)


func main() {
	db := setting.ConnectDB()

	router := gin.Default()

	customerRepo := customerRepo.CreateCustomerRepo(db)
	customerUsecase := customerUsecase.CreateCustomerUsecase(customerRepo)

	customerhandler.CreateCustomerHandler(router, customerUsecase)
	err := router.Run(":" + os.Getenv("PORT"))
	if err != nil {
		log.Fatal(err)
	}
}
