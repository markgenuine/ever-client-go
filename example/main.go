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

	idReq, err := ton.Client.GetBuildInfo()
	if err != nil {
		log.Fatal(idReq)
	}
	value, err := ton.Client.GetResp(idReq)
	if err != nil {
		log.Fatal(idReq)
	}

	fmt.Println("value: ", value)
	// ton, err := goton.NewTon(0)
	// if err != nil {
	// 	log.Fatal(err)
	// }
	// idReq, err := ton.Crypto.Factorize(domain.ParamsOfFactorize{Composite: "17ED48941A08F981"})
	// if err != nil {
	// 	log.Fatal(idReq)
	// }
	// value, err := ton.Client.GetResp(idReq)
	// if err != nil {
	// 	log.Fatal(idReq)
	// }
	// if len(value.(domain.ResultOfFactorize).Factors) == 0 || err != nil {
	// 	fmt.Println("test Failed - Error get Factorize method, err: %s", err)
	// }
	// if !(value.(domain.ResultOfFactorize).Factors[0] == "494C553B" && value.(domain.ResultOfFactorize).Factors[1] == "53911073") {
	// 	fmt.Println("test Failed - error value different factorize value")
	// }

	// fmt.Println(value)
}
