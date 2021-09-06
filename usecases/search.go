package usecases

import (
	"ES-Service/model"
	"ES-Service/repository"
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strings"

	"github.com/elastic/go-elasticsearch/v7"
)

var es *elasticsearch.Client
var err error

func init() {
	es, err = repository.ESClient()
	if err != nil {
		log.Fatal("error")
	}
}

func FinalResult(w http.ResponseWriter, r *http.Request) {
	fmt.Println("entered into fResult")
	result, err := SearchFResult()
	res, err1 := json.Marshal(result)

	if err1 != nil || err != nil {
		log.Fatal("error")
	}

	w.Write(res)
}

func SearchFResult() ([]model.Person, error) {
	var person []model.Person

	fmt.Println("ES connection: ", es)
	if err != nil {
		log.Fatal("error while making connection: ", err)
	}
	ctx := context.Background()

	// var person []model.Person

	// var query = `
	// "match": {
	// 	"Name.keyword": "AB Testin4"
	//   }`
	var query = `
	"match": {
		"Age": 104
	  }`
	finalQuery := makeQuery(query)
	var mapOutput model.SearchResult
	fmt.Println("json.NewEncoder encoded query:", finalQuery)

	res, err := es.Search(
		es.Search.WithContext(ctx),
		es.Search.WithIndex("person"),
		es.Search.WithBody(finalQuery),
		es.Search.WithTrackTotalHits(true),
	)
	if err != nil {
		log.Fatal("error")
	}

	log.Println("finaloutput: ", res.Body)
	err = json.NewDecoder(res.Body).Decode(&mapOutput)
	if err != nil {
		log.Fatal("error occured while decoding", err)
	}

	ab := mapOutput.Hits.ArrayHits

	for i := range ab {
		docMap := ab[i]
		person = append(person, model.Person{
			Name:       docMap.Source["Name"].(string),
			DocumentId: int(docMap.Source["DocumentId"].(float64)),
			Age:        int(docMap.Source["Age"].(float64)),
			Email:      docMap.Source["Email"].(string),
		})
	}
	return person, nil

}

func makeQuery(q string) *strings.Reader {

	var query = `{"query": {`
	query = query + q
	query = query + `}}`
	fmt.Println("\nquery:", query)

	var b strings.Builder
	b.WriteString(query)
	read := strings.NewReader(b.String())

	return read
}

// func AutoComplete(w http.ResponseWriter, r *http.Request) {
// 	var v elasticMemberRepo
// 	mem, err := v.AutoCompleteSuggestion()
// 	log.Println("output: ", mem)

// 	if err != nil {
// 		log.Fatal("returned error in auto complete: ", err)
// 	}

// }

// func (e *elasticMemberRepo) AutoCompleteSuggestion() (member []model.Member, err error) {

// 	var members []model.Member

// 	keyword := "r"
// 	ctx := context.Background()

// 	searchSuggester := elastic.NewCompletionSuggester("data").Text(keyword).Field("suggest")
// 	searchSource := elastic.NewSearchSource().Suggester(searchSuggester)
// 	searchResult, err := e.Client.Search().Index(e.Index).Type(e.Type).SearchSource(searchSource).Do(ctx)
// 	for _, ops := range searchResult.Suggest["data"] {
// 		for _, op := range ops.Options {
// 			if op.Source == nil {
// 				continue
// 			}
// 			var member model.Member
// 			err := json.Unmarshal(*op.Source, &member)
// 			if err != nil {
// 				log.Println(err)
// 				continue
// 			}
// 			members = append(members, member)
// 		}
// 	}
// 	return members, nil

// }
