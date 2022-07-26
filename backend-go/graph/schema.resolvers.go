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

// TrySent takes user input, parses and scores user input.
func (r *mutationResolver) TrySent(ctx context.Context, input *model.TrySentInput) (*model.TrySentScore, error) {
	userId := auth.FromContext(ctx, "user_id")
	if userId == 0 {
		return nil, fmt.Errorf("please log in")
	}

	score, err := queries.TrySentence(userId, input.SentID, input.UserInputJSON)
	if err != nil {
		return nil, err
	}
	success := true

	return &model.TrySentScore{Success: &success, Score: &score}, nil
}

// ListUserArticles resolver returns all articles with the progress of the user.
func (r *queryResolver) ListUserArticles(ctx context.Context) ([]*model.UserArticle, error) {
	userId := auth.FromContext(ctx, "user_id")
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

// DisplaySeenSents returns sentences in an article that the user hasn't tried.
func (r *queryResolver) DisplaySeenSents(ctx context.Context, articleID *int) ([]*model.SeenSent, error) {
	userId := auth.FromContext(ctx, "user_id")
	if userId == 0 {
		return nil, fmt.Errorf("please log in")
	}

	result, err := queries.QuerySeenSents(userId, *articleID)
	if err != nil {
		return nil, err
	}

	return result, nil
}

// ExamineCorrectSents return user intput text on the last trial to be shown on examine page.
func (r *queryResolver) ExamineCorrectSents(ctx context.Context, articleID *int) ([]*model.SeenSent, error) {
	userId := auth.FromContext(ctx, "user_id")
	if userId == 0 {
		return nil, fmt.Errorf("please log in")
	}

	result, err := queries.QueryExamineCorrectSents(userId, *articleID)
	if err != nil {
		return nil, err
	}

	return result, nil
}

// ExamineIncorrectSents is the resolver for the examineIncorrectSents field.
func (r *queryResolver) ExamineIncorrectSents(ctx context.Context, articleID *int) ([]*model.IncorrectSeenSent, error) {
	userId := auth.FromContext(ctx, "user_id")
	if userId == 0 {
		return nil, fmt.Errorf("please log in")
	}
	result, err := queries.QueryExamineIncorrectSents(userId, *articleID)
	if err != nil {
		return nil, err
	}
	return result, nil
}

// DisplayDueSents is the resolver for the displayDueSents field.
func (r *queryResolver) DisplayDueSents(ctx context.Context, daysAhead *int) ([]*model.DueSent, error) {
	userId := auth.FromContext(ctx, "user_id")
	if userId == 0 {
		return nil, fmt.Errorf("please log in")
	}
	result, err := queries.QueryDueSents(userId, *daysAhead)
	if err != nil {
		return nil, err
	}
	return result, nil
}

// ZzzDevQuery is the resolver for the zzzDevQuery field.
func (r *queryResolver) ZzzDevQuery(ctx context.Context) (*int, error) {
	// 	// DUMMY DB OPERATIONS

	queries.QueryDueSents(14, 10)

	ret := -1
	return &ret, nil
}

// Mutation returns generated.MutationResolver implementation.
func (r *Resolver) Mutation() generated.MutationResolver { return &mutationResolver{r} }

// Query returns generated.QueryResolver implementation.
func (r *Resolver) Query() generated.QueryResolver { return &queryResolver{r} }

type mutationResolver struct{ *Resolver }
type queryResolver struct{ *Resolver }
