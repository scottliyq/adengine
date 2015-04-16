//  Copyright (c) 2014 Couchbase, Inc.
//  Licensed under the Apache License, Version 2.0 (the "License"); you may not use this file
//  except in compliance with the License. You may obtain a copy of the License at
//    http://www.apache.org/licenses/LICENSE-2.0
//  Unless required by applicable law or agreed to in writing, software distributed under the
//  License is distributed on an "AS IS" BASIS, WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND,
//  either express or implied. See the License for the specific language governing permissions
//  and limitations under the License.
package main

import (
	"log"

	"github.com/blevesearch/bleve"
)

func openIndex(path string) bleve.Index {
	index, err := bleve.Open(path)
	if err == bleve.ErrorIndexPathDoesNotExist {
		log.Printf("Creating new index...")
		// create a mapping
		indexMapping := buildIndexMapping()
		index, err = bleve.New(path, indexMapping)
		if err != nil {
			log.Fatal(err)
		}
	} else if err == nil {
		log.Printf("Opening existing index...")
	} else {
		log.Fatal(err)
	}
	return index
}

func processUpdate(index bleve.Index, path string) {
	log.Printf("updated: %s", path)

}

func processDelete(index bleve.Index, path string) {
	log.Printf("delete: %s", path)

}

func buildIndexMapping() *bleve.IndexMapping {

	indexMapping := bleve.NewIndexMapping()

	return indexMapping
}
