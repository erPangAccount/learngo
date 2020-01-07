package persist

import (
	"crawler/engine"
	"github.com/pkg/errors"
	"golang.org/x/net/context"
	"gopkg.in/olivere/elastic.v6"
)

func ItemServer(index string) (chan engine.Item, error) {
	out := make(chan engine.Item)
	client, err := elastic.NewClient(
		elastic.SetURL(engine.ElasticHost),
		elastic.SetSniff(false),
	)
	if err != nil {
		return nil, err
	}
	go func() {
		itemCount := 0
		for {
			item := <-out
			//log.Printf("item server : got item #%d: %v", itemCount, item)
			itemCount++

			_, err := Save(client, index, item)
			if err != nil {
				//log.Printf("item server: save error %v: %v", item, err)
			}
		}
	}()
	return out, nil
}

func Save(client *elastic.Client, index string, data engine.Item) (string, error) {

	if data.Type == "" {
		data.Type = "index"
	}

	indexService := client.Index().Index(index).Type(data.Type).BodyJson(data)
	if data.Id != "" {
		//判断数据是否存在
		existsService := elastic.NewExistsService(client)
		if exists, err := existsService.Index(index).Type(data.Type).Id(data.Id).Do(context.Background()); err != nil || exists {
			if err != nil {
				return "", err
			}
			if exists {
				return "exists", errors.New("The data already exists")
			}
		}
		indexService.Id(data.Id)
	}

	response, err := indexService.Do(context.Background())
	if err != nil {
		return "", nil
	}

	return response.Id, nil
}
