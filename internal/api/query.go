package api

import (
	"context"
	"log"
	"os"
	"strings"

	"github.com/machinebox/graphql"
)

type Contributions []struct {
	Author struct {
		Name string `json:"name"`
		Slug string `json:"slug"`
	} `json:"author"`
	Contribution string `json:"contribution"`
}

type Document struct {
	Id                string
	Genres            []string
	Pages             int
	Rating            float64
	RatingsCount      int `json:"ratings_count"`
	Title             string
	Year              int    `json:"release_year"`
	Date              string `json:"release_date"`
	Description       string
	AuthorNames       []string `json:"author_names"`
	ContributionTypes []string `json:"contribution_types"`
	Contributions     Contributions
	Image             struct {
		URL string `json:"url"`
	}
	Moods        []string
	Tags         []string
	Slug         string
	HasAudiobook bool `json:"has_audiobook"`
	HasEbook     bool `json:"has_ebook"`
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

func Query(args []string) Hits {
	// set endpoint
	client := graphql.NewClient("https://api.hardcover.app/v1/graphql")

	// make a request
	req := graphql.NewRequest(`
    query SearchQuery ($title: String!) {
      search(query:$title) {
        results
      }
    }
  `)

	// set any variables
	req.Var("title", strings.Join(args, " "))

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
