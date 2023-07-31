package main

import (
	"context"
	"crypto/tls"
	"fmt"
	"net/http"
	"os"
	"strings"
	"time"

	"github.com/opensearch-project/opensearch-go"
	"github.com/opensearch-project/opensearch-go/opensearchapi"
)

type Data struct {
	Name      string    `json:"name"`
	Timestamp time.Time `json:"timestamp"`
}

func main() {
	client, err := opensearch.NewClient(opensearch.Config{
		Transport: &http.Transport{
			TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
		},
		Addresses: []string{"http://localhost:9200"},
		Username:  "admin", // For testing only. Don't store credentials in code.
		Password:  "admin",
	})
	if err != nil {
		fmt.Println("cannot initialize", err)
		os.Exit(1)
	}

	// Define index settings.
	settings := strings.NewReader(`{
     'settings': {
       'index': {
            'number_of_shards': 1,
            'number_of_replicas': 2
            }
          }
     }`)
	res := opensearchapi.IndicesCreateRequest{
		Index: "users",
		Body:  settings,
	}
	fmt.Println("Creating index")
	fmt.Println(res)

	for i := 1; i < 100; i++ {
		// Create an index with non-default settings.
		// Add a document to the index.
		document := strings.NewReader(
			fmt.Sprintf(`{
        "name": "Moneyball",
        "timestamp": "%v",
    }`, time.Now()),
		)

		docId := fmt.Sprint(i)
		req := opensearchapi.IndexRequest{
			Index:      "users",
			DocumentID: docId,
			Body:       document,
		}
		fmt.Println(i)
		insertResponse, err := req.Do(context.Background(), client)
		if err != nil {
			fmt.Println("failed to insert document ", err)
			os.Exit(1)
		}
		fmt.Println("Inserting a document")
		fmt.Println(insertResponse)
		defer insertResponse.Body.Close()

	}
}
