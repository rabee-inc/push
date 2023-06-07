package service

import (
	"context"

	"cloud.google.com/go/firestore"
	"github.com/rabee-inc/go-pkg/cloudfirestore"
	"github.com/rabee-inc/go-pkg/cloudtasks"
	"github.com/rabee-inc/go-pkg/log"
	"github.com/rabee-inc/push/appengine/src/config"
	"github.com/rabee-inc/push/appengine/src/model/input"
	"github.com/rabee-inc/push/appengine/src/repository"
)

type sender struct {
	rToken     repository.Token
	rFCM       repository.FCM
	rReserve   repository.Reserve
	cTasks     *cloudtasks.Client
	cFirestore *firestore.Client
}

func NewSender(
	rToken repository.Token,
	rFCM repository.FCM,
	rReserve repository.Reserve,
	cTasks *cloudtasks.Client,
	cFirestore *firestore.Client,
) Sender {
	return &sender{
		rToken,
		rFCM,
		rReserve,
		cTasks,
		cFirestore,
	}
}

func (s *sender) AllUsers(
	ctx context.Context,
	param *SenderAllUsersInput,
) (*SenderAllUsersOutput, error) {
	// 全員が所属しているトピックにプッシュ通知を送信
	err := s.rFCM.SendMessageByTopic(
		ctx,
		param.AppID,
		config.TopicAll,
		param.PushID,
		param.Message,
	)
	if err != nil {
		log.Warning(ctx, err)
		return nil, err
	}
	return &SenderAllUsersOutput{
		Success: true,
	}, nil
}

func (s *sender) Users(
	ctx context.Context,
	param *input.WorkerSendUsers,
) error {
	for _, userID := range param.UserIDs {
		src := &input.WorkerSendUser{
			AppID:   param.AppID,
			UserID:  userID,
			PushID:  param.PushID,
			Message: param.Message,
		}
		err := s.cTasks.AddTask(ctx, config.QueueSendUser, "/worker/send/user", src)
		if err != nil {
			log.Warning(ctx, err)
			return err
		}
	}
	return nil
}

func (s *sender) User(
	ctx context.Context,
	param *input.WorkerSendUser,
) error {
	tokens, err := s.rToken.List(
		ctx,
		param.AppID,
		param.UserID,
	)
	if err != nil {
		log.Warning(ctx, err)
		return err
	}
	if len(tokens) == 0 {
		return nil
	}

	err = s.rFCM.SendMessageByTokens(
		ctx,
		param.AppID,
		tokens,
		param.PushID,
		param.Message,
	)
	if err != nil {
		log.Warning(ctx, err)
		return err
	}
	return nil
}

func (s *sender) Reserved(
	ctx context.Context,
	param *SenderReservedInput,
) error {
	// 送信対象の予約を取得
	reserves, _, err := s.rReserve.List(
		ctx,
		param.AppID,
		&repository.ReserveListQuery{
			OverdueReserved: true,
			FilterStatuses: []config.ReserveStatus{
				config.ReserveStatusReserved,
			},
			Limit: 500,
		})
	if err != nil {
		log.Error(ctx, err)
		return err
	}
	if len(reserves) == 0 {
		return nil
	}

	// 処理中に設定
	ctx = cloudfirestore.RunBulkWriter(ctx, s.cFirestore)
	for _, reserve := range reserves {
		reserve.Status = config.ReserveStatusProcessing
		err = s.rReserve.Update(ctx, param.AppID, reserve)
		if err != nil {
			log.Error(ctx, err)
			return err
		}
	}
	if ctx, err = cloudfirestore.CommitBulkWriter(ctx); err != nil {
		log.Error(ctx, err)
		return err
	}

	ctx = cloudfirestore.RunBulkWriter(ctx, s.cFirestore)
	for _, reserve := range reserves {
		// 送信
		if len(reserve.UserIDs) > 0 {
			err = s.Users(ctx, &input.WorkerSendUsers{
				AppID:   param.AppID,
				UserIDs: reserve.UserIDs,
				PushID:  reserve.ID,
				Message: reserve.Message,
			})
			if err != nil {
				log.Error(ctx, err)
			}
		} else {
			_, err = s.AllUsers(ctx, &SenderAllUsersInput{
				AppID:   param.AppID,
				PushID:  reserve.ID,
				Message: reserve.Message,
			})
		}
		if err != nil {
			// 失敗
			reserve.Status = config.ReserveStatusFailure
		} else {
			// 成功
			reserve.Status = config.ReserveStatusSuccess
		}

		// ステータスを変更
		err = s.rReserve.Update(ctx, param.AppID, reserve)
		if err != nil {
			log.Error(ctx, err)
			return err
		}
	}
	if ctx, err = cloudfirestore.CommitBulkWriter(ctx); err != nil {
		log.Error(ctx, err)
		return err
	}
	return nil
}
