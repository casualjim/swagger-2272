package main

import (
	"log"
	"swaggerbug/client"
	"swaggerbug/client/hello_tag"
)

func main() {
	cl := client.NewHTTPClientWithConfig(nil, client.DefaultTransportConfig().WithHost("localhost:8083"))
	ep, err := cl.HelloTag.IDOfHelloEndpoint(hello_tag.NewIDOfHelloEndpointParams())
	if err != nil {
		log.Fatalln(err)
	}
	log.Println(ep)
}
