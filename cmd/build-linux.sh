#CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build .
set GOARCH=amd64
set GOOS=linux

CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o migrate .