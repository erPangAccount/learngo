package itemRpc

import (
	"crawler/engine"
	"crawler/persist"
	"gopkg.in/olivere/elastic.v6"
)

type ItemService struct {
	Client *elastic.Client
	Index  string
}

func (i *ItemService) Save(item engine.Item, result *string) error {
	id, err := persist.Save(i.Client, i.Index, item)
	if err == nil {
		*result = id
	}
	return err
}
