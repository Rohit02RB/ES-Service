package model

import (
	"github.com/elastic/go-elasticsearch/v7"
)

type Person struct {
	Name       string
	Age        int
	Email      string
	DocumentId int
}

type Member struct {
	Name string `json:"name"`
	URL  string `json:"url"`
}

type recipeRepo struct {
	client *elasticsearch.Client
}

type SearchResult struct {
	Hits HitsSearchResult `json:"hits"`
}

type HitsSearchResult struct {
	ArrayHits []ArrayHits `json:"hits"`
}

type ArrayHits struct {
	Source map[string]interface{} `json:"_source"`
}
