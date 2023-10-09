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

	owner := "kudoas"
	name := "sandbox"
	issue := 151
	variables := map[string]interface{}{
		"repositoryOwner": githubv4.String(owner),
		"repositoryName":  githubv4.String(name),
		"issueNumber":     githubv4.Int(issue),
	}
	var query struct {
		Repository struct {
			Issue struct {
				ID              githubv4.ID
				TrackedInIssues struct {
					Nodes []struct {
						ID githubv4.ID
					}
				} `graphql:"trackedInIssues(first: 5)"`
			} `graphql:"issue(number: $issueNumber)"`
		} `graphql:"repository(owner: $repositoryOwner, name: $repositoryName)"`
	}
	err := client.Query(context.Background(), &query, variables)
	if err != nil {
		panic(err)
	}
	fmt.Println("children issue", query.Repository.Issue)

	var query2 struct {
		Node struct {
			Issue struct {
				Assignees struct {
					Nodes []struct {
						ID   githubv4.ID
						Name githubv4.String
					}
				} `graphql:"assignees(first: 5)"`
				Labels struct {
					Nodes []struct {
						ID   githubv4.ID
						Name githubv4.String
					}
				} `graphql:"labels(first: 10)"`
				Milestone struct {
					ID    githubv4.ID
					Title githubv4.String
				} `graphql:"milestone"`
			} `graphql:"... on Issue"`
		} `graphql:"node(id: $issueID)"`
	}
	variables2 := map[string]interface{}{
		"issueID": githubv4.ID(query.Repository.Issue.TrackedInIssues.Nodes[0].ID),
	}
	err = client.Query(context.Background(), &query2, variables2)
	if err != nil {
		println(fmt.Errorf("error: %v", err))
	}

	fmt.Println("issues", query2.Node.Issue)

	var mutation struct {
		UpdateIssue struct {
			Issue struct {
				ID githubv4.ID
			}
		} `graphql:"updateIssue(input: $input)"`
	}

	input := githubv4.UpdateIssueInput{
		ID:          query.Repository.Issue.ID,
		AssigneeIDs: &[]githubv4.ID{query2.Node.Issue.Assignees.Nodes[0].ID},
		LabelIDs:    &[]githubv4.ID{query2.Node.Issue.Labels.Nodes[0].ID},
		MilestoneID: &query2.Node.Issue.Milestone.ID,
	}
	err = client.Mutate(context.Background(), &mutation, input, nil)
	if err != nil {
		println(fmt.Errorf("error: %v", err))
	}
}
