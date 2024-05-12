package services

import (
	"context"

	"github.com/kudoas/graphql-sample/graph/model"
	"github.com/volatiletech/sqlboiler/boil"
)

type UserService interface {
	GetUserByName(ctx context.Context, name string) (*model.User, error)
}

type RepositoryService interface {
	GetRepositoryByName(ctx context.Context, name, owner string) (*model.Repository, error)
}

type Services interface {
	UserService
	RepositoryService
}

type services struct {
	*userService
	*repositoryService
}

func New(exec boil.ContextExecutor) Services {
	return &services{
		userService: &userService{
			exec: exec,
		},
		repositoryService: &repositoryService{
			exec: exec,
		},
	}
}
