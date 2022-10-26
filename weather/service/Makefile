BINARY_NAME=weather

build:
	set GOOS=linux
	set GOARCH=amd64
	go build -a -installsuffix cgo -o cmd/command/${BINARY_NAME} cmd/command/weather.go

run:
	go run cmd/command/weather.go
