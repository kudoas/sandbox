package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.
// Code generated by github.com/99designs/gqlgen version v0.17.45

import (
	"context"
	"fmt"

	"github.com/kudoas/graphql-sample/graph/model"
	"github.com/kudoas/graphql-sample/internal"
)

// AddProjectV2ItemByID is the resolver for the addProjectV2ItemById field.
func (r *mutationResolver) AddProjectV2ItemByID(ctx context.Context, input model.AddProjectV2ItemByIDInput) (*model.AddProjectV2ItemByIDPayload, error) {
	panic(fmt.Errorf("not implemented: AddProjectV2ItemByID - addProjectV2ItemById"))
}

// Repository is the resolver for the repository field.
func (r *queryResolver) Repository(ctx context.Context, name string, owner string) (*model.Repository, error) {
	panic(fmt.Errorf("not implemented: Repository - repository"))
}

// User is the resolver for the user field.
func (r *queryResolver) User(ctx context.Context, name string) (*model.User, error) {
	panic(fmt.Errorf("not implemented: User - user"))
}

// Node is the resolver for the node field.
func (r *queryResolver) Node(ctx context.Context, id string) (model.Node, error) {
	panic(fmt.Errorf("not implemented: Node - node"))
}

// Mutation returns internal.MutationResolver implementation.
func (r *Resolver) Mutation() internal.MutationResolver { return &mutationResolver{r} }

// Query returns internal.QueryResolver implementation.
func (r *Resolver) Query() internal.QueryResolver { return &queryResolver{r} }

type mutationResolver struct{ *Resolver }
type queryResolver struct{ *Resolver }
