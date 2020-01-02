package main

import (
	"crawler/engine"
	"crawler/service"
	"crawler/service/itemRpc"
	"gopkg.in/olivere/elastic.v6"
	"log"
)

func main() {
	log.Fatal(startServer(":1234", "test"))
}

func startServer(host string, index string) error {
	client, e := elastic.NewClient(
		elastic.SetURL(engine.ElasticHost),
		elastic.SetSniff(false),
	)
	if e != nil {
		return e
	}

	return service.ServeRpc(
		&itemRpc.ItemService{
			Client: client,
			Index:  index,
		}, host)
}
