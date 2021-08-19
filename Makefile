build: cmd/main.go
	@go build -o build/trouxa ./cmd/
install:
	@cp build/trouxa /usr/bin/
