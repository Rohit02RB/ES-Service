package repository

import (
	"log"

	"github.com/elastic/go-elasticsearch/v7"
)

func ESClient() (*elasticsearch.Client, error) {
	es, err := elasticsearch.NewDefaultClient()
	if err != nil {
		log.Fatal("error during startup ", err)
	}

	log.Println(elasticsearch.Version)
	res, err := es.Info()
	log.Println(res)
	if err != nil {
		log.Fatal("error while making connection: ", err)
	}
	return es, nil

}
