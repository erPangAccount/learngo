package main

import (
	"crawler_distributed/config"
	"crawler_distributed/itemRpc"
	"crawler_distributed/rpc"
	"gopkg.in/olivere/elastic.v6"
	"log"
)

func main() {
	log.Fatal(startServer(config.ItemServiceHost, config.ElasticIndex))
}

func startServer(host string, index string) error {
	client, e := elastic.NewClient(
		elastic.SetURL(config.ElasticHost),
		elastic.SetSniff(false),
	)
	if e != nil {
		return e
	}

	return rpc.ServeRpc(
		&itemRpc.ItemService{
			Client: client,
			Index:  index,
		}, host)
}
