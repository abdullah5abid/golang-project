build:
	GOOS=linux GOARCH=amd64 go build -o bin/admin-api cmd/backend/*.go

test:
	go test -count=1 -p 1 ./...
