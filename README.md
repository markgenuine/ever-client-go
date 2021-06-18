# TON Client
TON SDK Client library Golang bindings based itself on [TON-SDK](https://github.com/tonlabs/TON-SDK).

[![version](https://img.shields.io/github/v/tag/move-ton/ton-client-go.svg)](https://github.com/move-ton/ton-client-go/releases/latest)
[![license](https://img.shields.io/github/license/move-ton/ton-client-go.svg)](https://github.com/move-ton/ton-client-go/blob/master/LICENSE)
[![Go version](https://img.shields.io/badge/go-1.14+-blue.svg)](https://github.com/moovweb/gvm)
[![Chat on Telegram RU](https://img.shields.io/badge/Chat%20on-Telegram%20RU-blue)](https://t.me/MOVETON_SDK_RU)
[![Chat on Telegram EN](https://img.shields.io/badge/Chat%20on-Telegram%20EN-blue)](https://t.me/MOVETON_SDK_EN)

Many thanks to [@temamagic](https://github.com/temamagic) for advice on architecture, tests, code and commit style.

## Installation

```sh
$ go get -u github.com/move-ton/ton-client-go
```
or

```sh
$ git clone https://github.com/move-ton/ton-client-go.git
$ cd ton-client-go
```

#### Installation for MAC OS 
```
#Set path to library
install_name_tool -id PATH_WITH_BINDING/gateway/client/lib/darwin/libton_client.dylib PATH_WITH_BINDING/gateway/client/lib/darwin/libton_client.dylib```

#Add to ~/.bashrc or execute everytime 
export CGO_LDFLAGS="-LPATH_WITH_BINDING/gateway/client/lib/darwin -lton_client"
```
#### Installation for Linux
```
#Add to ~/.bashrc or execute everytime 
export LD_LIBRARY_PATH=PATH_WITH_BINDING/gateway/client/lib/linux/:$LD_LIBRARY_PATH
export CGO_LDFLAGS="-LPATH_WITH_BINDING/gateway/client/lib/linux -lton_client"
```

#### Or use "-exec" for example:
```
go build
go run  -exec "env DYLD_LIBRARY_PATH=/path-with-lib/" main.go
go test -exec "env DYLD_LIBRARY_PATH=/path-with-lib/ ./... " -v
```

## Tests
```
$ go test ./... -v
$ go run ./example/*.go
```

## Usage
```golang
import goton "github.com/move-ton/ton-client-go"
```

## Example
```golang
package main

import (
	"fmt"
	"github.com/move-ton/ton-client-go/domain"
	"log"

	goton "github.com/move-ton/ton-client-go"
)

func main() {
	ton, err := goton.NewTon(domain.BaseUrl)
	if err != nil {
		log.Fatal(err)
	}

	defer ton.Client.Destroy()

	value, err := ton.Client.Version()
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Version bindings is: ", value.Version)
}
```
For more examples see *_test.go files
[ton-client-go/usecase](https://github.com/move-ton/ton-client-go/tree/master/usecase)
