package service

import (
	"context"

	"github.com/rabee-inc/push/appengine/src/config"
	"github.com/rabee-inc/push/appengine/src/lib/cloudtasks"
	"github.com/rabee-inc/push/appengine/src/lib/log"
	"github.com/rabee-inc/push/appengine/src/model"
	"github.com/rabee-inc/push/appengine/src/repository"
)

type sender struct {
	tRepo           repository.Token
	fRepo           repository.Fcm
	tCli            *cloudtasks.Client
	workerServiceID string
}

// MessageByUserIDs ... メッセージを複数のユーザーIDに対して送信する
func (s *sender) MessageByUserIDs(ctx context.Context, appID string, userIDs []string, msg *model.Message) error {
	for _, userID := range userIDs {
		src := &model.TaskQueueParamSendUserID{
			AppID:   appID,
			UserID:  userID,
			Message: msg,
		}
		err := s.tCli.AddTask(ctx, config.QueueSendUser, s.workerServiceID, "/worker/send/user", src)
		if err != nil {
			log.Warningm(ctx, "s.tCli.AddTask", err)
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
		err = s.tCli.AddTask(ctx, config.QueueSendToken, s.workerServiceID, "/worker/send/token", src)
		if err != nil {
			log.Warningm(ctx, "s.tCli.AddTask", err)
			return err
		}
	}
	return nil
}

// MessageByToken ... メッセージをトークンに対して送信する
func (s *sender) MessageByToken(ctx context.Context, appID string, token string, msg *model.Message) error {
	err := s.fRepo.SendMessage(ctx, appID, token, msg)
	if err != nil {
		log.Warningm(ctx, "s.fRepo.SendMessage", err)
		return err
	}
	return nil
}

// NewSender ... サービスを作成する
func NewSender(
	tRepo repository.Token,
	fRepo repository.Fcm,
	tCli *cloudtasks.Client,
	workerServiceID string) Sender {
	return &sender{
		tRepo:           tRepo,
		fRepo:           fRepo,
		tCli:            tCli,
		workerServiceID: workerServiceID,
	}
}
