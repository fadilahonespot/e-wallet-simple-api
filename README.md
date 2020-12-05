# E-Wallet Simple Api
project as a medium for learning to build a simple API

## Go Version
- 1.5
## Installation
### Using Docker
```docker
 $~ docker-compose up --build -d
```
### Without Docker
- Setting database in .env file
- Import file database.sql to postgres
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
customer_number = 1001
```
- Request body
```body
{
    "to_account_number": "555002",
    "amount": 100
}
```
