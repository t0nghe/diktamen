package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"backend-go/graph/generated"
	"backend-go/graph/model"
	"context"
	"fmt"
)

// UserSignUp is the resolver for the userSignUp field.
func (r *mutationResolver) UserSignUp(ctx context.Context, input model.UserCredentials) (*model.Message, error) {
	panic(fmt.Errorf("not implemented"))
}

// UserLogIn is the resolver for the userLogIn field.
func (r *mutationResolver) UserLogIn(ctx context.Context, input model.UserCredentials) (*model.LoginToken, error) {
	panic(fmt.Errorf("not implemented"))
}

// UserAttempt is the resolver for the userAttempt field.
func (r *mutationResolver) UserAttempt(ctx context.Context, input *model.UserAttempt) (*model.Message, error) {
	panic(fmt.Errorf("not implemented"))
}

// ListUserArticles is the resolver for the listUserArticles field.
func (r *queryResolver) ListUserArticles(ctx context.Context, username string) ([]*model.UserArticle, error) {
	panic(fmt.Errorf("not implemented"))
}

// GetUserFullArticle is the resolver for the getUserFullArticle field.
func (r *queryResolver) GetUserFullArticle(ctx context.Context, username string) (*model.FullArticle, error) {
	panic(fmt.Errorf("not implemented"))
}

// FetchSentAudio is the resolver for the fetchSentAudio field.
func (r *queryResolver) FetchSentAudio(ctx context.Context, username string, sentID string) (*model.SentDetails, error) {
	panic(fmt.Errorf("not implemented"))
}

// ScoreArticle is the resolver for the scoreArticle field.
func (r *queryResolver) ScoreArticle(ctx context.Context, username string, articleID string) ([]*model.SentScore, error) {
	panic(fmt.Errorf("not implemented"))
}

// Mutation returns generated.MutationResolver implementation.
func (r *Resolver) Mutation() generated.MutationResolver { return &mutationResolver{r} }

// Query returns generated.QueryResolver implementation.
func (r *Resolver) Query() generated.QueryResolver { return &queryResolver{r} }

type mutationResolver struct{ *Resolver }
type queryResolver struct{ *Resolver }
