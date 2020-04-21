package persist

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"log"

	"github.com/corrots/go-demo/crawler/engine"
	"github.com/elastic/go-elasticsearch/v7"
	"github.com/elastic/go-elasticsearch/v7/esapi"
)

func ItemSaver(index string) (chan engine.Item, error) {
	out := make(chan engine.Item)
	es, err := elasticsearch.NewDefaultClient()
	if err != nil {
		return nil, fmt.Errorf("Error creating elasticsearch client: %s\n", err)
	}
	go func() {
		itemCount := 0
		for {
			item := <-out
			fmt.Printf("Item Saver: got item #%d: %v\n", itemCount, item)
			itemCount++
			if err := save(item, es, index); err != nil {
				log.Printf("Item %v save err: %v\n", item, err)
			}
		}
	}()
	return out, nil
}

func save(item engine.Item, es *elasticsearch.Client, index string) error {
	jsonData, err := json.Marshal(item)
	if err != nil {
		return fmt.Errorf("Json Marshal err: %s\n", err)
	}

	req := esapi.IndexRequest{
		Index:      index,
		DocumentID: item.Id,
		Body:       bytes.NewReader(jsonData),
		Refresh:    "true",
		Pretty:     true,
	}
	res, err := req.Do(context.Background(), es)
	defer res.Body.Close()
	//fmt.Printf("es Index res: %+v\n", res)
	return err
}
