build: main.go
	@go build -o build/trouxa .
install:
	@cp build/trouxa /usr/bin/
