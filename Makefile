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
add_darwin_amd64:
	install_name_tool -id gateway/client/lib/darwin/amd64/libton_client.dylib gateway/client/lib/darwin/amd64/libton_client.dylib
	export CGO_LDFLAGS="-Lgateway/client/lib/darwin/amd64 -lton_client"

add_darwin_arm64:
	install_name_tool -id gateway/client/lib/darwin/arm64/libton_client.dylib gateway/client/lib/darwin/arm64/libton_client.dylib
	export CGO_LDFLAGS="-Lgateway/client/lib/darwin/arm64 -lton_client"

lib_install:
	go build -o ./tools/downloadLibs ./tools/downloadLibs.go
	./tools/downloadLibs 1.44.4
