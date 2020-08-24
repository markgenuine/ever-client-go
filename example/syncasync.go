package main

import (
	"fmt"
	"log"
	"strconv"

	goton "github.com/move-ton/go-ton-sdk"
)

func syncasync() {

	config, err := goton.ParseConfigFile("config.toml")
	if err != nil {
		log.Println("Error read config file, err: . Settings setup on default.", err)
		config = goton.NewConfig(0)
	}

	client, err := goton.InitClient(config)
	if err != nil {
		log.Fatal("Init client error", err)
	}
	defer client.Destroy()

	req, err := client.Version()
	fmt.Println("Version sync: req1: ", req, " err: ", err)

	client.RequestType = 1
	req1, _ := client.Version()
	fmt.Println("req1: ", req1)

	req2, _ := client.Version()
	fmt.Println("req2: ", req2)
	if err != nil {
		log.Fatal("Error get version, err: ", err)
	}

	req3, _ := client.Version()
	fmt.Println("req3: ", req3)
	if err != nil {
		log.Fatal("Error get version, err: ", err)
	}

	fmt.Println(req1, " ", req2, " ", req3)
	fmt.Println("Version async ress1: ", client.GetResp(convertToInt(req1)))
	fmt.Println("Version async  ress2: ", client.GetResp(convertToInt(req2)))
	fmt.Println("Version async  ress3: ", client.GetResp(convertToInt(req3)))

	client.RequestType = 0
	req, err = client.Version()
	fmt.Println("Version sync2 : req4: ", req, " err: ", err)

}

func convertToInt(req string) (value int) {
	value, _ = strconv.Atoi(req)
	return
}
