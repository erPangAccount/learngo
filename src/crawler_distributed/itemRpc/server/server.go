package main

import (
	"crawler_distributed/config"
	"crawler_distributed/itemRpc"
	"crawler_distributed/rpc"
	"flag"
	"fmt"
	"gopkg.in/olivere/elastic.v6"
	"log"
)

var port = flag.Int("port", 0, "itemService port")

func main() {
	flag.Parse()
	if *port == 0 {
		panic("must have a port!")
	}

	log.Fatal(startServer(fmt.Sprintf(":%d", *port), config.ElasticIndex))
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
