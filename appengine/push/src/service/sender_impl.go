package service

import (
	"context"

	"cloud.google.com/go/firestore"

	"github.com/rabee-inc/go-pkg/cloudtasks"
	"github.com/rabee-inc/go-pkg/log"
	"github.com/rabee-inc/go-pkg/timeutil"
	"github.com/rabee-inc/push/appengine/push/src/config"
	"github.com/rabee-inc/push/appengine/push/src/model"
	"github.com/rabee-inc/push/appengine/push/src/repository"
)

type sender struct {
	tRepo repository.Token
	fRepo repository.Fcm
	rRepo repository.Reserve
	tCli  *cloudtasks.Client
	fCli  *firestore.Client
}

func (s *sender) AllUsers(ctx context.Context, appID string, pushID string, msg *model.Message) error {
	// プッシュ通知を送信
	err := s.fRepo.SendMessageByTopic(ctx, appID, config.TopicAll, pushID, msg)
	if err != nil {
		log.Warning(ctx, err)
		return err
	}
	return nil
}

func (s *sender) Users(ctx context.Context, appID string, userIDs []string, pushID string, msg *model.Message) error {
	for _, userID := range userIDs {
		src := &model.CloudTasksParamSendUser{
			AppID:   appID,
			UserID:  userID,
			PushID:  pushID,
			Message: msg,
		}
		err := s.tCli.AddTask(ctx, config.QueueSendUser, "/worker/send/user", src)
		if err != nil {
			log.Warning(ctx, err)
			return err
		}
	}
	return nil
}

func (s *sender) User(ctx context.Context, appID string, userID string, pushID string, msg *model.Message) error {
	// ユーザーに紐づくTokenを取得
	tokens, err := s.tRepo.ListByUser(ctx, appID, userID)
	if err != nil {
		log.Warning(ctx, err)
		return err
	}
	if len(tokens) == 0 {
		return nil
	}

	// プッシュ通知を送信
	err = s.fRepo.SendMessageByTokens(ctx, appID, tokens, pushID, msg)
	if err != nil {
		log.Warning(ctx, err)
		return err
	}
	return nil
}

func (s *sender) Reserved(ctx context.Context, appID string) error {
	// 送信対象の予約を取得
	now := timeutil.NowUnix()
	rsvs, _, err := s.rRepo.ListBySend(ctx, appID, now, 30, nil)
	if err != nil {
		log.Error(ctx, err)
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
		log.Error(ctx, err)
		return err
	}

	bt = s.fCli.Batch()
	for _, rsv := range rsvs {
		// 送信
		if len(rsv.UserIDs) > 0 {
			err = s.Users(ctx, appID, rsv.UserIDs, rsv.ID, rsv.Message)
			if err != nil {
				log.Error(ctx, err)
			}
		} else {
			err = s.AllUsers(ctx, appID, rsv.ID, rsv.Message)
			if err != nil {
				log.Error(ctx, err)
			}
		}
		if err != nil {
			// 失敗
			rsv.Status = config.ReserveStatusFailure
		} else {
			// 成功
			rsv.Status = config.ReserveStatusSuccess
		}
		// ステータスを変更
		now := timeutil.NowUnix()
		_ = s.rRepo.BtUpdate(ctx, bt, appID, rsv, now)
	}
	_, err = bt.Commit(ctx)
	if err != nil {
		log.Error(ctx, err)
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
