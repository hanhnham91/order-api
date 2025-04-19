build:
	go build -o ./tmp/app.exe ./cmd/main.go

prepare:
	@go install github.com/air-verse/air@latest
	@go install github.com/swaggo/swag/cmd/swag@latest
	@go install github.com/daixiang0/gci@latest
	@go install github.com/rubenv/sql-migrate/...@latest
	@go install github.com/securego/gosec/v2/cmd/gosec@latest
	@curl -sSfL https://raw.githubusercontent.com/golangci/golangci-lint/master/install.sh | sh -s -- -b ${GOPATH}/bin v1.61.0

lint:
	golangci-lint run

mod:
	@go mod tidy
	@go mod download

run:
	go run cmd/main.go

sec:
	@gosec ./...

gen:
	@mockery

test:	
	@mkdir coverage || true
	@go test -race -v -coverprofile=coverage/coverage.txt.tmp -count=1 ./...  
	@cat coverage/coverage.txt.tmp | grep -v "mock_" > coverage/coverage.txt
	@go tool cover -func=coverage/coverage.txt
	@go tool cover -html=coverage/coverage.txt -o coverage/index.html

gci:
	@gci write -s standard -s default -s "Prefix(github.com/hanhnham91/mfv-test)" .

