package main

import (
	"fmt"
	"log"

	goton "github.com/markgenuine/ton-client-go"
)

func main() {

	ton, err := goton.NewTon(2)
	if err != nil {
		log.Fatal(err)
	}

	defer ton.Client.Destroy()

	idReq, err := ton.Client.Version()
	if err != nil {
		log.Fatal(idReq)
	}

	value, err := ton.Client.GetResp(idReq)
	if err != nil {
		log.Fatal(idReq)
	}

	fmt.Println("Version bindings is: ", value)
}
