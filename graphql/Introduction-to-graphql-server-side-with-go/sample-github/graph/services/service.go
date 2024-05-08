package services

import (
	"context"

	"github.com/kudoas/graphql-sample/graph/model"
	"github.com/volatiletech/sqlboiler/boil"
)

type UserService interface {
	GetUserByName(ctx context.Context, name string) (*model.User, error)
}

type Services interface {
	UserService
}

type services struct {
	*userService
}

func New(exec boil.ContextExecutor) Services {
	return &services{
		userService: &userService{
			exec: exec,
		},
	}
}
