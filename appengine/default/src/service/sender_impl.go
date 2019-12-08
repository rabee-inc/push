package service

import (
	"context"

	"cloud.google.com/go/firestore"

	"github.com/rabee-inc/push/appengine/default/src/config"
	"github.com/rabee-inc/push/appengine/default/src/lib/cloudtasks"
	"github.com/rabee-inc/push/appengine/default/src/lib/log"
	"github.com/rabee-inc/push/appengine/default/src/lib/util"
	"github.com/rabee-inc/push/appengine/default/src/model"
	"github.com/rabee-inc/push/appengine/default/src/repository"
)

type sender struct {
	tRepo repository.Token
	fRepo repository.Fcm
	rRepo repository.Reserve
	tCli  *cloudtasks.Client
	fCli  *firestore.Client
}

func (s *sender) AllUsers(ctx context.Context, appID string, msg *model.Message) error {
	// プッシュ通知を送信
	err := s.fRepo.SendMessageByTopic(ctx, appID, config.TopicAll, msg)
	if err != nil {
		log.Warningm(ctx, "s.fRepo.SendMessageByTopic", err)
		return err
	}
	return nil
}

func (s *sender) Users(ctx context.Context, appID string, userIDs []string, msg *model.Message) error {
	for _, userID := range userIDs {
		src := &model.CloudTasksParamSendUser{
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

func (s *sender) User(ctx context.Context, appID string, userID string, msg *model.Message) error {
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

func (s *sender) Reserved(ctx context.Context, appID string) error {
	// 送信対象の予約を取得
	now := util.TimeNowUnix()
	rsvs, _, err := s.rRepo.ListBySend(ctx, appID, now, 30, nil)
	if err != nil {
		log.Errorm(ctx, "s.rRepo.ListBySend", err)
		return err
	}
	if len(rsvs) == 0 {
		return nil
	}

	// 処理中に設定
	bt := s.fCli.Batch()
	for _, rsv := range rsvs {
		rsv.Status = config.ReserveStatusProcessing
		s.rRepo.BtUpdate(ctx, bt, appID, rsv, now)
	}
	_, err = bt.Commit(ctx)
	if err != nil {
		log.Errorm(ctx, "bt.Commit", err)
		return err
	}

	bt = s.fCli.Batch()
	for _, rsv := range rsvs {
		// 送信
		err := s.AllUsers(ctx, appID, rsv.Message)
		if err != nil {
			// 失敗
			log.Errorm(ctx, "s.AllUsers", err)
			rsv.Status = config.ReserveStatusFailure
		} else {
			// 成功
			rsv.Status = config.ReserveStatusSuccess
		}
		// ステータスを変更
		now := util.TimeNowUnix()
		_ = s.rRepo.BtUpdate(ctx, bt, appID, rsv, now)
	}
	_, err = bt.Commit(ctx)
	if err != nil {
		log.Errorm(ctx, "bt.Commit", err)
		return err
	}
	return nil
}

// NewSender ... サービスを作成する
func NewSender(
	tRepo repository.Token,
	fRepo repository.Fcm,
	rRepo repository.Reserve,
	tCli *cloudtasks.Client,
	fCli *firestore.Client) Sender {
	return &sender{
		tRepo: tRepo,
		fRepo: fRepo,
		rRepo: rRepo,
		tCli:  tCli,
		fCli:  fCli,
	}
}
