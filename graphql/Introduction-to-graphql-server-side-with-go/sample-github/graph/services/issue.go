package services

import (
	"context"
	"net/url"

	"github.com/kudoas/graphql-sample/graph/db"
	"github.com/kudoas/graphql-sample/graph/model"
	"github.com/volatiletech/sqlboiler/boil"
	"github.com/volatiletech/sqlboiler/v4/queries/qm"
)

type issueService struct {
	exec boil.ContextExecutor
}

func (i *issueService) GetIssueByRepoAndNumber(ctx context.Context, repoID string, number int) (*model.Issue, error) {
	issue, err := db.Issues(
		qm.Select(
			db.IssueColumns.ID,
			db.IssueColumns.URL,
			db.IssueColumns.Title,
			db.IssueColumns.Closed,
			db.IssueColumns.Number,
			db.IssueColumns.Repository,
		),
		db.IssueWhere.Repository.EQ(repoID),
		db.IssueWhere.Number.EQ(int64(number)),
	).One(ctx, i.exec)
	if err != nil {
		return nil, err
	}
	return convertIssue(issue), nil
}

func convertIssue(issue *db.Issue) *model.Issue {
	url, _ := url.Parse(issue.URL)

	return &model.Issue{
		ID:         issue.ID,
		Title:      issue.Title,
		URL:        model.URL{URL: *url},
		Closed:     (issue.Closed == 1),
		Number:     int(issue.Number),
		Repository: &model.Repository{ID: issue.Repository},
	}
}
