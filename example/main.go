package main

import (
	"fmt"
	"log"

	goton "github.com/move-ton/go-ton-sdk"
)

func main() {

	config, err := goton.ParseConfigFile("config.toml")
	if err != nil {
		log.Println("Error read config file, err: . Settings setup on default.", err)
	}

	config = goton.NewConfig(1)

	client, err := goton.InitClient(config)
	if err != nil {
		log.Fatal("Init client error", err)
	}
	defer client.Destroy()

	value, err := client.Version()
	if err != nil {
		log.Fatal("Error get version, err: ", err)
	}

	fmt.Println(value)

	value, err = client.MnemonicWords()

	if err != nil {
		log.Fatal("Error get mnemonic words, err: ", err)
	}

	fmt.Println(value)

}
