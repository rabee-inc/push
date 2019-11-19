package service

import (
	"context"

	"github.com/rabee-inc/push/appengine/default/src/config"
	"github.com/rabee-inc/push/appengine/default/src/lib/cloudtasks"
	"github.com/rabee-inc/push/appengine/default/src/lib/log"
	"github.com/rabee-inc/push/appengine/default/src/model"
	"github.com/rabee-inc/push/appengine/default/src/repository"
)

type sender struct {
	tRepo repository.Token
	fRepo repository.Fcm
	tCli  *cloudtasks.Client
}

func (s *sender) MessageByAll(ctx context.Context, appID string, msg *model.Message) error {
	// プッシュ通知を送信
	err := s.fRepo.SendMessageByTopic(ctx, appID, config.TopicAll, msg)
	if err != nil {
		log.Warningm(ctx, "s.fRepo.SendMessageByTopic", err)
		return err
	}
	return nil
}

func (s *sender) MessageByUserIDs(ctx context.Context, appID string, userIDs []string, msg *model.Message) error {
	for _, userID := range userIDs {
		src := &model.CloudTasksParamSendUserID{
			AppID:   appID,
			UserID:  userID,
			Message: msg,
		}
		err := s.tCli.AddTask(ctx, config.QueueSendUser, "/worker/send/user", src)
		if err != nil {
			log.Warningm(ctx, "s.tCli.AddTask", err)
			return err
		}
	}
	return nil
}

func (s *sender) MessageByUserID(ctx context.Context, appID string, userID string, msg *model.Message) error {
	// ユーザーに紐づくTokenを取得
	tokens, err := s.tRepo.ListByUser(ctx, appID, userID)
	if err != nil {
		log.Errorm(ctx, "s.tRepo.GetListByUserID", err)
		return err
	}

	// プッシュ通知を送信
	err = s.fRepo.SendMessageByTokens(ctx, appID, tokens, msg)
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
	tCli *cloudtasks.Client) Sender {
	return &sender{
		tRepo: tRepo,
		fRepo: fRepo,
		tCli:  tCli,
	}
}
