# ton-client-go
FreeTON SDK Golang language based itself on [TON-SDK](https://github.com/tonlabs/TON-SDK).
1. For complited from source code SDK-lib:
```
export CGO_LDFLAGS="-L/path-to-installation/TON-SDK/target/release/deps/ -lton_client"
```
Linux:
```
export LD_LIBRARY_PATH=/path-to-installation/TON-SDK/target/release/deps/
```
MacOS:
```
export DYLD_LIBRARY_PATH=/path-to-installation/TON-SDK/target/release/deps/
```
and add file to "lib" directory darwin for macOS and linux.

2. For completed of binary lib complited:
```
export CGO_LDFLAGS="-L/path-with-lib/ -lton_client"
```
Linux:
```
export LD_LIBRARY_PATH=/path-with-lib/
```
MacOS:
```
export DYLD_LIBRARY_PATH=/path-with-lib/
```

Or use "-exec" for example:
```
go run  -exec "env DYLD_LIBRARY_PATH=/path-with-lib/" main.go
go test -exec "env DYLD_LIBRARY_PATH=/path-with-lib/ " -v
```

## Install
```sh
$ go get -u github.com/move-ton/ton-client-go
```
## Test
$ go test -v
$ go run ./example/*.go

## Usage
```go
import goton "github.com/move-ton/ton-client-go"
```
