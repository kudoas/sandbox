package main

import (
	"context"
	"fmt"
	"log"
	"os"

	"github.com/shurcooL/githubv4"
	"golang.org/x/oauth2"
)

type ProjectV2QueryResponse struct {
	Node struct {
		ProjectV2 struct {
			Fields struct {
				Nodes []struct {
					ProjectV2Field struct {
						ID   githubv4.ID
						Name githubv4.String
					} `graphql:"... on ProjectV2Field"`
				}
			} `graphql:"fields(first: 1)"`
		} `graphql:"... on ProjectV2"`
	} `graphql:"node(id: $projectNodeID)"`
}

func NewProjectV2QueryResponse() *ProjectV2QueryResponse {
	return &ProjectV2QueryResponse{}
}

var githubToken = os.Getenv("GITHUB_TOKEN")
var projectNodeID = os.Getenv("PROJECT_NODE_ID")

func main() {
	src := oauth2.StaticTokenSource(
		&oauth2.Token{AccessToken: githubToken},
	)
	httpClient := oauth2.NewClient(context.Background(), src)
	client := githubv4.NewClient(httpClient)

	p := NewProjectV2QueryResponse()
	if err := client.Query(context.Background(), &p, map[string]interface{}{
		"projectNodeID": githubv4.ID(projectNodeID),
	}); err != nil {
		log.Println(err)
	}

	fmt.Println("Project Name:", p.Node.ProjectV2.Fields.Nodes[0].ProjectV2Field.Name)
	fmt.Println("Project ID:", p.Node.ProjectV2.Fields.Nodes[0].ProjectV2Field.ID)
}
