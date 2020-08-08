package main

import (
	"fmt"
	"log"

	goton "github.com/move-ton/go-ton-sdk"
)

func main() {

	config, err := goton.ParseConfigFile("./config.toml")
	if err != nil {
		log.Fatal("Error read config file, err: ", err)
	}

	client, err := goton.InitClient(*config)
	if err != nil {
		log.Fatal("Init client error", err)
	}
	defer client.Destroy()

	value, err := client.version()
	if err != nil {
		log.Fatal("Error get version, err: ", err)
	}

	fmt.Println(value)
}
