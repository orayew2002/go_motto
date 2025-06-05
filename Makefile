test:
	go test ./...

run:
	go mod tidy
	go run cmd/app/main.go
