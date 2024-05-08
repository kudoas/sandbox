package services

import (
	"context"

	"github.com/kudoas/graphql-sample/graph/db"
	"github.com/kudoas/graphql-sample/graph/model"
	"github.com/volatiletech/sqlboiler/boil"
	"github.com/volatiletech/sqlboiler/v4/queries/qm"
)

type userService struct {
	exec boil.ContextExecutor
}

func (u *userService) GetUserByName(ctx context.Context, name string) (*model.User, error) {
	user, err := db.Users(
		qm.Select(db.UserColumns.ID, db.UserColumns.Name),
		db.UserWhere.Name.EQ(name),
	).One(ctx, u.exec)
	if err != nil {
		return nil, err
	}
	return convertUser(user), nil
}

func convertUser(user *db.User) *model.User {
	return &model.User{
		ID:   user.ID,
		Name: user.Name,
	}
}
