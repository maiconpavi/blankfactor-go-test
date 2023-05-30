swag-init:
	go get github.com/swaggo/swag/gen@latest
	go get github.com/swaggo/swag/cmd/swag@latest
	go run github.com/swaggo/swag/cmd/swag init

clean:
	rm -rf ./bin ./vendor

build: clean swag-init
	
	go get ./...
	mkdir bin
	env GOOS=linux GOARCH=amd64 CGO_ENABLED=1 go build -ldflags="-s -w" -o bin/main main.go




run: swag-init
	go run .
