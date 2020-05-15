package service

import (
	"context"

	"github.com/rabee-inc/go-pkg/log"
	"github.com/rabee-inc/push/appengine/push/src/config"
	"github.com/rabee-inc/push/appengine/push/src/repository"
)

type register struct {
	tRepo repository.Token
	fRepo repository.Fcm
}

func (s *register) SetToken(ctx context.Context, appID string, userID string, platform string, deviceID string, token string) error {
	// 取得
	cToken, err := s.tRepo.Get(ctx, appID, userID, platform, deviceID)
	if err != nil {
		log.Errorm(ctx, "s.tRepo.Get", err)
		return err
	}

	// 保存
	err = s.tRepo.Put(ctx, appID, userID, platform, deviceID, token)
	if err != nil {
		log.Errorm(ctx, "s.tRepo.Put", err)
		return err
	}

	// 全員のトピックに登録
	if token != "" && token != cToken {
		err = s.fRepo.SubscribeTopic(ctx, appID, config.TopicAll, []string{token})
		if err != nil {
			log.Warningm(ctx, "s.fRepo.SubscribeTopic", err)
			// レスポンスを返す事を優先させるため握りつぶす
		}
	}
	return nil
}

func (s *register) DeleteToken(ctx context.Context, appID string, userID string, platform string, deviceID string) error {
	// 取得
	token, err := s.tRepo.Get(ctx, appID, userID, platform, deviceID)
	if err != nil {
		log.Errorm(ctx, "s.tRepo.Get", err)
		return err
	}
	if token == "" {
		log.Warningf(ctx, "not exist token: app_id: %s, platform: %s, device_id: %s", appID, platform, deviceID)
		// 既に削除済みの場合は成功を返す
		return nil
	}

	// 削除
	err = s.tRepo.Delete(ctx, appID, userID, platform, deviceID)
	if err != nil {
		log.Errorm(ctx, "s.tRepo.Delete", err)
		return err
	}

	// 全員のトピックから削除
	if token != "" {
		err = s.fRepo.UnsubscribeTopic(ctx, appID, config.TopicAll, []string{token})
		if err != nil {
			log.Warningm(ctx, "s.fRepo.UnsubscribeTopic", err)
			// レスポンスを返す事を優先させるため握りつぶす
		}
	}
	return nil
}

// NewRegister ... サービスを作成する
func NewRegister(tRepo repository.Token, fRepo repository.Fcm) Register {
	return &register{
		tRepo: tRepo,
		fRepo: fRepo,
	}
}
