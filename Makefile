report:
	go test -v ./... -coverprofile=coverage.out
	go tool cover -html=coverage.out

code-log:
	golangci-lint run ./... > error.log

code:
	golangci-lint run ./...

mock:
	mockery --name ClientGateway \
				--case underscore \
				--recursive --dir ./domain \
				--outpkg clientmock --output ./gateway/clientmock