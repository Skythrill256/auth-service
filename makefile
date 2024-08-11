build:
	@go build -o bin/$(APP_NAME) cmd/main/main.go

run:
	@go run cmd/main/main.go

