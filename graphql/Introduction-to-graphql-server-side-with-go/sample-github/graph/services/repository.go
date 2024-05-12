package services

import (
	"context"
	"log"
	"time"

	"github.com/kudoas/graphql-sample/graph/db"
	"github.com/kudoas/graphql-sample/graph/model"
	"github.com/volatiletech/sqlboiler/boil"
	"github.com/volatiletech/sqlboiler/v4/queries/qm"
)

type repositoryService struct {
	exec boil.ContextExecutor
}

func (r *repositoryService) GetRepositoryByName(ctx context.Context, name, owner string) (*model.Repository, error) {
	boil.DebugMode = true
	defer func() { boil.DebugMode = false }()
	repo, err := db.Repositories(
		qm.Select(db.RepositoryColumns.ID, db.RepositoryColumns.Name, db.RepositoryColumns.CreatedAt),
		db.RepositoryWhere.Name.EQ(name),
		db.RepositoryWhere.Owner.EQ(owner),
	).One(ctx, r.exec)
	if err != nil {
		return nil, err
	}
	return convertRepository(repo), nil
}

func convertRepository(repo *db.Repository) *model.Repository {
	c, err := time.Parse(time.DateTime, repo.CreatedAt)
	if err != nil {
		log.Printf("failed to parse time: %v", err)
	}
	return &model.Repository{
		ID:        repo.ID,
		Name:      repo.Name,
		CreatedAt: c,
	}
}
