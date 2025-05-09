include .env

build-local:
	go build -o main main.go

build-production:
	GOOS=linux GOARCH=amd64 go build -ldflags="-s -w" -o main main.go

start:
	./main

dev:
	GO_ENV=development go run main.go

fmt:
	gofmt -w ./ && golangci-lint run

test-all:
	GO_ENV=testing go test ./tests/...

test:
	GO_ENV=testing go test $(FILE) -v

migrate-up:
	migrate -path migrations -database "mysql://$(MYSQL_USER):$(MYSQL_PASSWORD)@tcp($(MYSQL_HOST))/$(MYSQL_DATABASE)" -verbose up

migrate-down:
	migrate -path migrations -database "mysql://$(MYSQL_USER):$(MYSQL_PASSWORD)@tcp($(MYSQL_HOST))/$(MYSQL_DATABASE)" -verbose down 1

migrate-test:
	migrate -path migrations -database "mysql://$(MYSQL_USER):$(MYSQL_PASSWORD)@tcp($(MYSQL_HOST))/$(MYSQL_DATABASE)_testing" -verbose up 

migrate-force:
	migrate -path migrations -database "mysql://$(MYSQL_USER):$(MYSQL_PASSWORD)@tcp($(MYSQL_HOST))/$(MYSQL_DATABASE)" -verbose force $(VERSION)

migrate-force-test:
	migrate -path migrations -database "mysql://$(MYSQL_USER):$(MYSQL_PASSWORD)@tcp($(MYSQL_HOST))/$(MYSQL_DATABASE)_testing" -verbose force $(VERSION)

create-migration:
	migrate create -ext sql -dir migrations -seq $(NAME)
