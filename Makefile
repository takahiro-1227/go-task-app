build-local:
	go build -o main main.go

build-production:
	GOOS=linux GOARCH=amd64 go build -ldflags="-s -w" -o main main.go

start:
	./main

dev:
	GO_ENV=development go run main.go

format:
	gofmt -w ./ && golangci-lint run
  
migrate-up:
	migrate -path migrations -database "mysql://admin:admin@tcp(localhost:3306)/go_task_app" -verbose up

migrate-down:
	migrate -path migrations -database "mysql://admin:admin@tcp(localhost:3306)/go_task_app" -verbose down 1

migrate-force:
	migrate -path migrations -database "mysql://admin:admin@tcp(localhost:3306)/go_task_app" -verbose force $(VERSION)

create-migration:
	migrate create -ext sql -dir migrations -seq $(NAME)
