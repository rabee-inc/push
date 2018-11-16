package service

import (
	"context"

	"github.com/aikizoku/push/src/config"
	"github.com/aikizoku/push/src/lib/taskqueue"
	"github.com/aikizoku/push/src/model"
	"github.com/aikizoku/push/src/repository"
	"google.golang.org/appengine/log"
)

type sender struct {
	tRepo repository.Token
	fRepo repository.Fcm
}

func (s *sender) SendMessageToUserIDs(ctx context.Context, userIDs []string, msg *model.Message) error {
	for _, userID := range userIDs {
		src := &model.SendUserID{
			UserID:  userID,
			Message: msg,
		}
		err := taskqueue.AddTaskToJSON(ctx, config.QueueSendUser, "/worker/send/user", src)
		if err != nil {
			log.Warningf(ctx, "taskqueue.NewJSONPostTask error: %s", err.Error())
			return err
		}
	}
	return nil
}

func (s *sender) SendMessageToUserID(ctx context.Context, userID string, msg *model.Message) error {
	tokens, err := s.tRepo.GetMultiToUserID(ctx, userID)
	if err != nil {
		log.Errorf(ctx, "s.tRepo.GetMultiToUserID error: %s", err.Error())
		return err
	}
	for _, token := range tokens {
		src := &model.SendToken{
			Token:   token,
			Message: msg,
		}
		err = taskqueue.AddTaskToJSON(ctx, config.QueueSendToken, "/worker/send/token", src)
		if err != nil {
			log.Warningf(ctx, "taskqueue.NewJSONPostTask error: %s", err.Error())
			return err
		}
	}
	return nil
}

func (s *sender) SendMessageToToken(ctx context.Context, token string, msg *model.Message) error {
	err := s.fRepo.SendMessage(ctx, token, msg)
	if err != nil {
		log.Warningf(ctx, "s.fRepo.SendMessage error: %s", err.Error())
		return err
	}
	return nil
}

// NewSender ... Senderを作成する
func NewSender(tRepo repository.Token, fRepo repository.Fcm) Sender {
	return &sender{
		tRepo: tRepo,
		fRepo: fRepo,
	}
}
