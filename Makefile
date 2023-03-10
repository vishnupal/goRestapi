build:
	@go build -o bin/goRestapi


run: build
	@./bin/goRestapi

test:
	@go test -v ./...