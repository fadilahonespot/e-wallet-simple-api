# E-Wallet Simple Api
project as a medium for learning to build a simple API with dokerizer and unit testing

## Go Version
- 1.5
## Installation
### Using Docker
```docker
 $~ docker-compose up --build -d
```
### Without Docker
- Setting database in .env file
- Use the postgres database
- Run project in terminal
```run
 $~ go run main.go
```
### Using Makefile Command
- Deploy project in docker
```run
 $~ make deploy
```
- Integration test
```run
 $~ make integration-test
```
- Unit test
```run
 $~ make unit-test
```
- Clean project in docker
```run
 $~ make clean
```
## Endpoint
### Add User Method POST
- Url
```endpoint
localhost:8090/account
```
- Request body
```
{
    "customer_name": "Siti Husna",
    "balance": 10000
}
```
### Get Account Method GET
- url
```endpoit
localhost:8090/account
```
### Get Account Detail Methode GET
- url
```endpoint
localhost:8090/account/{{account_number}}
```
### Transfer method PUT
- url
```endpoint
localhost:8090/transfer
```
- Request header
```header
customer_number = {{customer_number}}
```
- Request body
```body
{
    "to_account_number": "555002",
    "amount": 100
}
```
