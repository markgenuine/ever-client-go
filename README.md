# Everscale Client
Everscale SDK Client library Golang bindings based itself on [Ever SDK](https://github.com/tonlabs/ever-sdk).

[![version](https://img.shields.io/github/v/tag/move-ton/ton-client-go.svg)](https://github.com/markgenuine/ever-client-go/releases/latest)
[![license](https://img.shields.io/github/license/move-ton/ton-client-go.svg)](https://github.com/markgenuine/ever-client-go/blob/master/LICENSE)
[![Go version](https://img.shields.io/badge/go-1.16+-blue.svg)](https://github.com/moovweb/gvm)

[![Channel on Telegram](https://img.shields.io/badge/%20Channel-Telegram%20-blue)](https://t.me/ever_go)
[![Chat on Telegram RU](https://img.shields.io/badge/Chat%20on-Telegram%20RU-blue)](https://t.me/ever_go_ru)
[![Chat on Telegram EN](https://img.shields.io/badge/Chat%20on-Telegram%20EN-blue)](https://t.me/ever_go_en)

Many thanks to [@temamagic](https://github.com/temamagic) for advice on architecture, tests, code and commit style.

## Installation

```sh
$ go get -u github.com/markgenuine/ever-client-go
```
or

```sh
$ git clone https://github.com/markgenuine/ever-client-go.git
$ cd ever-client-go
```

#### Installation for MAC OS 
```
#Set path to library
install_name_tool -id PATH_WITH_BINDING/gateway/client/lib/darwin/libton_client.dylib PATH_WITH_BINDING/gateway/client/lib/darwin/libton_client.dylib

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
import goever "github.com/markgenuine/ever-client-go"
```

## Example
```golang
package main

import (
	"fmt"
	"github.com/markgenuine/ever-client-go/domain"
	"log"

	goton "github.com/markgenuine/ever-client-go"
)

func main() {
	ever, err := goever.NewEver("", domain.GetDevNetBaseUrls())
	if err != nil {
		log.Fatal(err)
	}

	defer ever.Client.Destroy()

	value, err := ever.Client.Version()
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Version bindings is: ", value.Version)
}
```
For more examples see *_test.go files
[ever-client-go/usecase](https://github.com/markgenuine/ever-client-go/tree/master/usecase)
