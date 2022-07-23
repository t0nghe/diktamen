package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"backend-go/graph/generated"
	"backend-go/graph/model"
	"backend-go/internal/auth"
	"backend-go/internal/queries"
	"context"
	"fmt"
)

// UserSignUp allows the user to sign up via GraphQL mutation `userSignUp` using a username and an email.
func (r *mutationResolver) UserSignUp(ctx context.Context, input model.UserCredentials) (*model.Message, error) {
	email := *input.Email
	err := auth.CreateUser(input.Username, email, input.Password)
	var success bool
	var message string
	if err != nil {
		success = false
		message = err.Error()
	} else {
		success = true
		message = "Successful"
	}
	return &model.Message{Success: success, Message: &message}, err
}

// UserLogin allows the user to log in via mutation `userLogin`.
func (r *mutationResolver) UserLogIn(ctx context.Context, input model.UserCredentials) (*model.LoginToken, error) {
	token, err := auth.UserLogin(input.Username, input.Password)

	return &model.LoginToken{Success: true, Token: &token}, err
}

// TrySent is the resolver for the trySent field.
func (r *mutationResolver) TrySent(ctx context.Context, input *model.TrySentInput) (*model.TrySentScore, error) {
	panic(fmt.Errorf("not implemented"))
}

// ListUserArticles resolver returns all articles with the progress of the user.
func (r *queryResolver) ListUserArticles(ctx context.Context) ([]*model.UserArticle, error) {
	userId := auth.FromContext(ctx, "user_id")
	fmt.Println(userId)
	if userId == 0 {
		return nil, fmt.Errorf("please log in")
	}
	ret, err := queries.GetUserArticlesList(userId)

	fmt.Println(ret)

	if err != nil {
		return nil, err
	}

	return ret, nil
}

// ListUserUnseenArticles is the resolver for the listUserUnseenArticles field.
func (r *queryResolver) ListUserUnseenArticles(ctx context.Context) ([]*model.UserArticle, error) {
	userId := auth.FromContext(ctx, "user_id")
	if userId == 0 {
		return nil, fmt.Errorf("please log in")
	}

	ret, err := queries.GetUserUnseenArticlesList(userId)

	fmt.Println(ret)

	if err != nil {
		return nil, err
	}

	return ret, nil
}

// DisplayUnseenSents queries for sentences the user has not seen in an article.
func (r *queryResolver) DisplayUnseenSents(ctx context.Context, articleID *int) ([]*model.UnseenSent, error) {
	// check if user is logged in
	userId := auth.FromContext(ctx, "user_id")
	if userId == 0 {
		return nil, fmt.Errorf("please log in")
	}

	// query for sentences in an article, where userid and articleid match; and userFinishedIndex is 0
	result, err := queries.QueryUserUnseenSents(userId, *articleID)
	if err != nil {
		return nil, err
	}

	return result, nil
}

// DisplaySeenSents is the resolver for the displaySeenSents field.
func (r *queryResolver) DisplaySeenSents(ctx context.Context, articleID *int) ([]*model.SeenSent, error) {
	panic(fmt.Errorf("not implemented"))
}

// Mutation returns generated.MutationResolver implementation.
func (r *Resolver) Mutation() generated.MutationResolver { return &mutationResolver{r} }

// Query returns generated.QueryResolver implementation.
func (r *Resolver) Query() generated.QueryResolver { return &queryResolver{r} }

type mutationResolver struct{ *Resolver }
type queryResolver struct{ *Resolver }
