package main

import (
	"context"
	"fmt"
	"os"

	"github.com/shurcooL/githubv4"
	"golang.org/x/oauth2"
)

func main() {
	src := oauth2.StaticTokenSource(
		&oauth2.Token{AccessToken: os.Getenv("GITHUB_TOKEN")},
	)
	httpClient := oauth2.NewClient(context.Background(), src)

	client := githubv4.NewClient(httpClient)
	query := struct {
		Repository struct {
			DatabaseID githubv4.Int
		} `graphql:"repository(owner: \"shurcooL\", name: \"notifications\")"`
	}{}
	err := client.Query(context.Background(), &query, nil)
	if err != nil {
		panic(err)
	}
	fmt.Println(query.Repository.DatabaseID)
}
