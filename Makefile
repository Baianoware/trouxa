build: cmd/trouxa.go
	@go build -o build/trouxa cmd/trouxa.go
install:
	@cp build/trouxa /usr/bin/
