package service

import (
	"context"

	"github.com/rabee-inc/go-pkg/log"
	"github.com/rabee-inc/push/appengine/src/config"
	"github.com/rabee-inc/push/appengine/src/model"
	"github.com/rabee-inc/push/appengine/src/repository"
)

type register struct {
	rToken repository.Token
	rFCM   repository.FCM
}

func NewRegister(
	rToken repository.Token,
	rFCM repository.FCM,
) Register {
	return &register{
		rToken,
		rFCM,
	}
}

func (s *register) Entry(
	ctx context.Context,
	param *RegisterEntryInput,
) (*RegisterEntryOutput, error) {
	// 取得
	tokenID := model.GenerateTokenDocID(param.Platform, param.DeviceID)
	token, err := s.rToken.Get(
		ctx,
		param.AppID,
		param.UserID,
		tokenID,
	)
	if err != nil {
		log.Error(ctx, err)
		return nil, err
	}
	if token == nil {
		token = model.NewToken(
			tokenID,
			param.Platform,
			param.DeviceID,
			param.Token,
		)
	} else {
		token.Token = param.Token
	}

	// 保存
	err = s.rToken.Set(
		ctx,
		param.AppID,
		param.UserID,
		token,
	)
	if err != nil {
		log.Error(ctx, err)
		return nil, err
	}

	// 全員が所属しているトピックに登録
	err = s.rFCM.SubscribeTopic(
		ctx,
		param.AppID,
		config.TopicAll,
		[]*model.Token{token},
	)
	if err != nil {
		log.Warning(ctx, err)
		// レスポンスを返す事を優先させるため握りつぶす
	}
	return &RegisterEntryOutput{
		Success: true,
	}, nil
}

func (s *register) Leave(
	ctx context.Context,
	param *RegisterLeaveInput,
) (*RegisterLeaveOutput, error) {
	// 取得
	tokenID := model.GenerateTokenDocID(param.Platform, param.DeviceID)
	token, err := s.rToken.Get(
		ctx,
		param.AppID,
		param.UserID,
		tokenID,
	)
	if err != nil {
		log.Error(ctx, err)
		return nil, err
	}
	if token == nil {
		return &RegisterLeaveOutput{
			Success: true,
		}, nil
	}

	// 削除
	err = s.rToken.Delete(
		ctx,
		param.AppID,
		param.UserID,
		token,
	)
	if err != nil {
		log.Error(ctx, err)
		return nil, err
	}

	// 全員が所属しているトピックから削除
	err = s.rFCM.UnsubscribeTopic(
		ctx,
		param.AppID,
		config.TopicAll,
		[]*model.Token{token},
	)
	if err != nil {
		log.Warning(ctx, err)
		// レスポンスを返す事を優先させるため握りつぶす
	}
	return &RegisterLeaveOutput{
		Success: true,
	}, nil
}
