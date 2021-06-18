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