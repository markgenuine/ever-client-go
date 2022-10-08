package main

import (
	"encoding/json"
	"fmt"
	"github.com/move-ton/ever-client-go/domain"
	"github.com/move-ton/ever-client-go/gateway/client"
	everNet "github.com/move-ton/ever-client-go/usecase/net"
	"log"
	"time"
)

func main() {
	config := domain.NewDefaultConfig("", domain.GetDevNetBaseUrls(), "")
	clientConn, err := client.NewClientGateway(config)
	if err != nil {
		log.Fatal("Error to connect: ", err.Error())
	}
	defer clientConn.Destroy()

	netUC := everNet.NewNet(config, clientConn)

	nowTime := int(time.Now().Unix())
	queryParams := &domain.ParamsOfSubscribeCollection{
		Collection: "messages",
		Filter:     json.RawMessage(fmt.Sprintf(`{"created_at":{"gt":%d}}`, nowTime)),
		Result:     "id, src, dst, boc, body, msg_type, value(format:DEC), created_at",
	}

	// # Create generator
	generator, handle, err := netUC.SubscribeCollection(queryParams)
	log.Println("generator: ", generator)
	log.Println("handle: ", handle)
	log.Println("err: ", err)

	log.Println("messages")
	var ts []interface{}
	go func() {
		for g := range generator {
			defer netUC.Unsubscribe(&domain.ResultOfSubscribeCollection{Handle: handle.Handle})
			if err != nil {
				log.Fatal(err.Error())
				break
			}
			log.Println(string(g))
		}
	}()
	time.Sleep(2 * time.Minute)
	fmt.Println(ts)
}
