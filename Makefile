.PHONY: dependency unit-test integration-test docker-up docker-down deploy clean docker-remove-image logs

dependency:
	@go get -v ./...

integration-test: docker-up dependency
	@go test -v ./...

unit-test: dependency
	@go test -v -short ./...

coverage-test: dependency
	@echo "=================================================================================="
	@echo "Coverage Test"
	@echo "=================================================================================="
	go test -v -coverprofile coverage.cov ./...
	@echo "\n"
	@echo "=================================================================================="
	@echo "All Package Coverage"
	@echo "=================================================================================="
	go tool cover -func coverage.cov

docker-up:
	@docker-compose up -d

docker-down:
	@docker-compose down

docker-remove-image:
	@docker rmi e-wallet-simple-api_server
	@docker volume rm e-wallet-simple-api_postgres
	@docker volume rm e-wallet-simple-api_go

logs:
	@docker-compose logs -f server

deploy: docker-up

clean: docker-down docker-remove-image

  
