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
	// catalog := client.Catalog()

	// redis_services, _, err := catalog.Service("redis_svc", "", nil)
	// if err != nil {
	// 	log.Fatalf("get service err: %s", err)
	// }
	// for index, redis_service := range redis_services {
	// 	log.Println("---> ", index)
	// 	spew.Dump(redis_service)
	// }
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

// type RegisterInfo struct {
// 	Datacenter      string                `json:"Datacenter"`
// 	Node            string                `json:"Node"`
// 	Address         string                `json:"Address"`
// 	TaggedAddresses ConsulTaggedAddresses `json:"TaggedAddresses"`
// 	Service         *ConsulService        `json:"Service"`
// 	Check           *ConsulCheck          `json:"Check"`
// }
// type ConsulCheck struct {
// 	Node      string `json:"Node"`
// 	CheckID   string `json:"CheckID"`
// 	Name      string `json:"Name"`
// 	Notes     string `json:"Notes"`
// 	Status    string `json:"Status"`
// 	ServiceID string `json:"ServiceID"`
// }
// type ConsulService struct {
// 	ID      string   `json:"ID"`
// 	Service string   `json:"Service"`
// 	Tags    []string `json:"Tags"`
// 	Address string   `json:"Address"`
// 	Port    int      `json:"Port"`
// }

// func NewConsulService(id, name, address string, port int) *ConsulService {
// 	service := &ConsulService{
// 		ID:      id,
// 		Service: name,
// 		Address: address,
// 		Port:    port,
// 	}
// 	return service
// }

// type ConsulTaggedAddresses struct {
// 	Lan string `json:"lan"`
// 	Wan string `json:"wan"`
// }

// func NewRegisterInfo(node, address string, service *ConsulService, check *ConsulCheck) *RegisterInfo {
// 	info := &RegisterInfo{
// 		Node:    node,
// 		Address: address,
// 		Service: service,
// 		Check:   check,
// 	}
// 	return info
// }
