# E-Wallet Simple Api
project as a learning tool to build a simple API with dokerizer and unit testing

## Go Version
- 1.5
## Installation
### Using Make Command
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
- Remove project in docker
```run
 $~ make clean
```
### Using Docker Command
```docker
 $~ docker-compose up --build -d
```
### Without Docker
- Setting database in .env file
- Using postgres database
- Run project in terminal
```run
 $~ go run main.go
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
account_number = {{account_number}}
```
- Request body
```body
{
    "to_account_number": "555002",
    "amount": 100
}
```
