build:
	GOOS=linux GOARCH=amd64 go build && chmod 777 store_toolkit
build-arm:
	GOOS=linux GOARCH=arm64 go build && chmod 777 store_toolkit