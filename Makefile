build:
	GOOS=linux GOARCH=amd64 go build -o main && mv main ./cmd/amd64
build-arm:
	GOOS=linux GOARCH=arm64 go build -o main && mv main ./cmd/arm64
build-mac:
	GOOS=darwin GOARCH=amd64 go build -o main && mv main ./cmd/amd64