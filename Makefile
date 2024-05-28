BINARY_NAME=bootstrap
.DEFAULT_GOAL := build

build:
	rm -rf target
	GOOS=linux GOARCH=arm64 go build -tags lambda.norpc -o target/$(BINARY_NAME) main.go
	zip -j target/$(BINARY_NAME).zip target/$(BINARY_NAME)
