package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"backend-go/graph/generated"
	"backend-go/graph/model"
	"backend-go/internal/auth"
	"backend-go/internal/dummymessages"
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

// UserDictation is the resolver for the userDictation field.
func (r *mutationResolver) UserDictation(ctx context.Context, input *model.UserDictation) (*model.Message, error) {
	panic(fmt.Errorf("not implemented"))
}

// CreateDummyMessage is a resolver used to show if the server works.
func (r *mutationResolver) CreateDummyMessage(ctx context.Context, input *model.InputMessage) (*model.Message, error) {
	var dmsg dummymessages.DummyMessage
	dmsg.Success = input.Success
	dmsg.Message = input.Message
	dmsgId := dmsg.Save()
	msgStr := fmt.Sprintf("Message ID %d saved.", dmsgId)
	gqlReturnMsg := &model.Message{
		Success: true,
		Message: &msgStr,
	}
	return gqlReturnMsg, nil
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

// GetUserFullArticle is the resolver for the getUserFullArticle field.
func (r *queryResolver) GetUserFullArticle(ctx context.Context, articleID string) (*model.FullArticle, error) {
	userId := auth.FromContext(ctx, "user_id")
	if userId == 0 {
		return nil, fmt.Errorf("please log in")
	}
	queries.QueryUserFullArticle(userId, articleID)

	return nil, nil
}

// FetchSentAudio is the resolver for the fetchSentAudio field.
func (r *queryResolver) FetchSentAudio(ctx context.Context, sentID string) (*model.SentDetails, error) {
	sentDetails := &model.SentDetails{SentID: sentID, MediaURI: "https://abc.com/file.mp3"}
	return sentDetails, nil
}

// ScoreArticle is the resolver for the scoreArticle field.
func (r *queryResolver) ScoreArticle(ctx context.Context, articleID string) ([]*model.SentScore, error) {
	panic(fmt.Errorf("not implemented"))
}

// DummyMessage is the resolver for the dummyMessage field.
func (r *queryResolver) DummyMessage(ctx context.Context) ([]*model.Message, error) {
	var resultMsgs []*model.Message
	dbMsgs := dummymessages.GetDummyMessages()

	for _, msg := range dbMsgs {
		// using &msg.Message on line 67
		// all returned records will be the string
		// as the last item.
		currentMsg := msg.Message
		resultMsg := &model.Message{
			Success: msg.Success,
			Message: &currentMsg,
		}
		resultMsgs = append(resultMsgs, resultMsg)
	}
	return resultMsgs, nil
}

// Mutation returns generated.MutationResolver implementation.
func (r *Resolver) Mutation() generated.MutationResolver { return &mutationResolver{r} }

// Query returns generated.QueryResolver implementation.
func (r *Resolver) Query() generated.QueryResolver { return &queryResolver{r} }

type mutationResolver struct{ *Resolver }
type queryResolver struct{ *Resolver }

// !!! WARNING !!!
// The code below was going to be deleted when updating resolvers. It has been copied here so you have
// one last chance to move it out of harms way if you want. There are two reasons this happens:
//  - When renaming or deleting a resolver the old code will be put in here. You can safely delete
//    it when you're done.
//  - You have helper methods in this file. Move them out to keep these resolver files clean.
func (r *mutationResolver) UserAttempt(ctx context.Context, input *model.UserDictation) (*model.Message, error) {
	panic(fmt.Errorf("not implemented"))
}
