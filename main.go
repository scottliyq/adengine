package main

//import (
//	"flag"
//	"fmt"
//	"github.com/blevesearch/bleve"
//	"log"
//)

//var indexPath = flag.String("index", "beer-search.bleve", "index path")

//type Message struct {
//	ID   string
//	From string
//	Body string
//}

type ItemData struct {
}

func main() {

	initDBData()
	//message := new(Message)
	//message.ID = "test"
	//message.From = "marty.schoch@gmail.com"
	//message.Body = "bleve indexing is easy"

	//index, err := bleve.Open("engine.first")
	//if err == bleve.ErrorIndexPathDoesNotExist {
	//	log.Printf("Creating new index...")
	//	mapping := bleve.NewIndexMapping()
	//	index, err = bleve.New("engine.first", mapping)

	//}

	//query := bleve.NewQueryStringQuery("bleve")
	//searchRequest := bleve.NewSearchRequest(query)
	//searchResult, _ := index.Search(searchRequest)

	//if searchResult != nil {
	//	fmt.Println(searchResult.Hits[0].ID)
	//} else {
	//	fmt.Println("none")
	//}

}
