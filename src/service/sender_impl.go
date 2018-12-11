package service

import (
	"context"
	"fmt"

	"github.com/aikizoku/push/src/config"
	"github.com/aikizoku/push/src/lib/log"
	"github.com/aikizoku/push/src/lib/taskqueue"
	"github.com/aikizoku/push/src/model"
	"github.com/aikizoku/push/src/repository"
)

type sender struct {
	tRepo repository.Token
	fRepo repository.Fcm
}

func (s *sender) SendMessageByUserIDs(ctx context.Context, userIDs []string, msg *model.Message) error {
	for _, userID := range userIDs {
		src := &model.SendUserID{
			UserID:  userID,
			Message: msg,
		}
		err := taskqueue.AddTaskByJSON(ctx, config.QueueSendUser, "/worker/send/user", src)
		if err != nil {
			log.Warningm(ctx, "taskqueue.AddTaskByJSON", err)
			return err
		}
	}
	return nil
}

func (s *sender) SendMessageByUserID(ctx context.Context, userID string, msg *model.Message) error {
	tokens, err := s.tRepo.GetListByUserID(ctx, userID)
	if err != nil {
		log.Errorm(ctx, "s.tRepo.GetListByUserID", err)
		return err
	}
	for _, token := range tokens {
		src := &model.SendToken{
			Token:   token,
			Message: msg,
		}
		err = taskqueue.AddTaskByJSON(ctx, config.QueueSendToken, "/worker/send/token", src)
		if err != nil {
			log.Warningm(ctx, "taskqueue.AddTaskByJSON", err)
			return err
		}
	}
	return nil
}

func (s *sender) SendMessageByToken(ctx context.Context, token string, msg *model.Message) error {
	if token == "" {
		err := fmt.Errorf("token is empty")
		log.Errorf(ctx, err.Error())
		return err
	}
	err := s.fRepo.SendMessage(ctx, token, msg)
	if err != nil {
		log.Warningm(ctx, "s.fRepo.SendMessage", err)
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
