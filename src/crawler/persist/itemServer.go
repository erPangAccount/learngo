package persist

import (
	"crawler/engine"
	"github.com/olivere/elastic"
	"github.com/pkg/errors"
	"golang.org/x/net/context"
	"log"
)

func ItemServer() (chan engine.Item, error) {
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
			log.Printf("item server : got item #%d: %v", itemCount, item)
			itemCount++

			_, err := save(client, item)
			if err != nil {
				log.Printf("item server: save error %v: %v", item, err)
			}
		}
	}()
	return out, nil
}

func save(client *elastic.Client, data engine.Item) (string, error) {

	if data.Type == "" {
		data.Type = "index"
	}

	indexService := client.Index().Index("test").Type(data.Type).BodyJson(data)
	if data.Id != "" {
		//判断数据是否存在
		existsService := elastic.NewExistsService(client)
		if exists, err := existsService.Index("test").Type(data.Type).Id(data.Id).Do(context.Background()); err != nil || exists {
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
