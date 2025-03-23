build-local:
	go build -o main main.go

build-production:
	GOOS=linux GOARCH=amd64 go build -ldflags="-s -w" -o main main.go

start:
	./main

dev:
	GO_ENV=development go run main.go
  
migrate-up:
	migrate -path migrations -database "mysql://admin:admin@tcp(localhost:3306)/go_task_app" -verbose up

migrate-down:
	migrate -path migrations -database "mysql://admin:admin@tcp(localhost:3306)/go_task_app" -verbose down

