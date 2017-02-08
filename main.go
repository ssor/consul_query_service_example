package main

import (
	"log"

	"fmt"

	"github.com/davecgh/go-spew/spew"
	"github.com/hashicorp/consul/api"
)

func main() {
	// // Make client config
	conf := api.DefaultConfig()

	// Create client
	client, err := api.NewClient(conf)
	if err != nil {
		log.Fatalf("err: %v", err)
	}
	health := client.Health()
	services, _, err := health.Service("redis_svc", "", true, nil)
	if err != nil {
		log.Fatalf("get service err: %s", err)
	}
	if len(services) <= 0 {
		fmt.Println("no service alive")
	} else {

		for index, entry := range services {
			fmt.Println("--->  ", index)
			spew.Dump(entry.Service)
		}
	}
}
