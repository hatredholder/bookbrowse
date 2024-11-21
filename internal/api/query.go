package api

import (
	"context"
	"log"
	"os"

	"github.com/machinebox/graphql"
)

type Document struct {
	Genres      []string
	Pages       int
	Rating      float64
	Title       string
	Year        int `json:"release_year"`
	Description string
	Authors     []string `json:"author_names"`
	Image       struct {
		URL string `json:"url"`
	}
}

type Hits []struct {
	Document `json:"document"`
}

type QueryResponse struct {
	Search struct {
		Results struct {
			Hits `json:"hits"`
		}
	}
}

func Query(title string) Hits {
	// set endpoint
	client := graphql.NewClient("https://api.hardcover.app/v1/graphql")

	// make a request
	req := graphql.NewRequest(`
    query SearchQuery ($title: String!) {
      search(query:$title, per_page: 5) {
        results
      }
    }
  `)

	// set any variables
	req.Var("title", title)

	// set header fields
	req.Header.Set("content-type", "application/json")
	req.Header.Set("authorization", os.Getenv("HARDCOVER_API_KEY"))

	// define a Context for the request
	ctx := context.Background()

	// run it and capture the response
	var respData QueryResponse
	if err := client.Run(ctx, req, &respData); err != nil {
		log.Fatal(err)
	}

	return respData.Search.Results.Hits
}
