package service

import (
	"context"

	"github.com/rabee-inc/push/src/config"
	"github.com/rabee-inc/push/src/lib/log"
	"github.com/rabee-inc/push/src/lib/taskqueue"
	"github.com/rabee-inc/push/src/model"
	"github.com/rabee-inc/push/src/repository"
)

type sender struct {
	tRepo repository.Token
	fRepo repository.Fcm
}

// MessageByUserIDs ... メッセージを複数のユーザーIDに対して送信する
func (s *sender) MessageByUserIDs(ctx context.Context, appID string, userIDs []string, msg *model.Message) error {
	for _, userID := range userIDs {
		src := &model.TaskQueueParamSendUserID{
			AppID:   appID,
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

// MessageByUserID ... メッセージをユーザーIDに対して送信する
func (s *sender) MessageByUserID(ctx context.Context, appID string, userID string, msg *model.Message) error {
	tokens, err := s.tRepo.GetListByUserID(ctx, appID, userID)
	if err != nil {
		log.Errorm(ctx, "s.tRepo.GetListByUserID", err)
		return err
	}
	for _, token := range tokens {
		src := &model.TaskQueueParamSendToken{
			AppID:   appID,
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

// MessageByToken ... メッセージをトークンに対して送信する
func (s *sender) MessageByToken(ctx context.Context, token string, msg *model.Message) error {
	err := s.fRepo.SendMessage(ctx, token, msg)
	if err != nil {
		log.Warningm(ctx, "s.fRepo.SendMessage", err)
		return err
	}
	return nil
}

// NewSender ... サービスを作成する
func NewSender(tRepo repository.Token, fRepo repository.Fcm) Sender {
	return &sender{
		tRepo: tRepo,
		fRepo: fRepo,
	}
}
