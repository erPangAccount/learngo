package client

import (
	"crawler/engine"
	"crawler_distributed/rpc"
	"log"
)

func ItemServer(host string) (chan engine.Item, error) {
	client, e := rpc.NewRpcClient(host)
	if e != nil {
		return nil, e
	}

	out := make(chan engine.Item)
	go func() {
		itemCount := 0
		for {
			item := <-out
			itemCount++

			// Call RPC to save item
			var result string
			err := client.Call("ItemService.Save", item, &result)
			if err != nil {
				log.Printf("item server: save error %v: %v", item, err)
			}
		}
	}()
	return out, nil
}
