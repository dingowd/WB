BINARY_NAME=weather

build:
	export GOOS=linux
	go build -a -installsuffix cgo -o cmd/command/${BINARY_NAME} cmd/command/weather.go

run:
	go run cmd/command/weather.go
