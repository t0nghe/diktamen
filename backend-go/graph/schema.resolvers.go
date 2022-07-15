package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"backend-go/graph/generated"
	"backend-go/graph/model"
	"backend-go/internal/dummymessages"
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

// CreateDummyMessage is the resolver for the createDummyMessage field.
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

// ListUserArticles is the resolver for the listUserArticles field.
func (r *queryResolver) ListUserArticles(ctx context.Context) ([]*model.UserArticle, error) {
	panic(fmt.Errorf("not implemented"))
}

// GetUserFullArticle is the resolver for the getUserFullArticle field.
func (r *queryResolver) GetUserFullArticle(ctx context.Context, articleID string) (*model.FullArticle, error) {
	panic(fmt.Errorf("not implemented"))
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
	var dbMsgs []dummymessages.DummyMessage

	dbMsgs = dummymessages.GetDummyMessages()

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
