package main

import (
	"fmt"
	elastic "gopkg.in/olivere/elastic.v3"
	"os"
)

func main() {
	esClient, err := elastic.NewClient(
		elastic.SetSniff(false),
		elastic.SetURL("**"),
	)

	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	res, err := esClient.Search().
		Query(elastic.NewMatchAllQuery()).
		Do()
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	for _, h := range res.Hits.Hits {
		a, _ := h.Source.MarshalJSON()
		fmt.Println(string(a))
	}
}
