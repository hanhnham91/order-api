build:
	go build -o ./tmp/app.exe ./cmd/main.go

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

