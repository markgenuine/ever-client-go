package main

import (
	"fmt"
	"github.com/markgenuine/ever-client-go/domain"
	"log"

	goever "github.com/markgenuine/ever-client-go"
)

func main() {
	ever, err := goever.NewEver("", domain.GetDevNetBaseUrls(), "")
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
