package itemRpc

import (
	"crawler/engine"
	"crawler/persist"
	"gopkg.in/olivere/elastic.v6"
	"log"
)

type ItemService struct {
	Client *elastic.Client
	Index  string
}

func (i *ItemService) Save(item engine.Item, result *string) error {
	id, err := persist.Save(i.Client, i.Index, item)
	log.Printf("saved in itemService")
	if err == nil {
		*result = id
	} else {
		log.Printf("itemService save error: %v : %v", item, err)
	}
	return err
}
