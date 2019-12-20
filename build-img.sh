BIN_NAME=apiservice

GOOS=linux GOARCH=amd64 go build -o $BIN_NAME main.go
docker build -t $BIN_NAME .