package service

import (
	"context"
	"net/http"

	"github.com/rabee-inc/push/appengine/default/src/config"
	"github.com/rabee-inc/push/appengine/default/src/lib/log"
	"github.com/rabee-inc/push/appengine/default/src/repository"
)

type register struct {
	tRepo repository.Token
	fRepo repository.Fcm
}

func (s *register) SetToken(ctx context.Context, appID string, userID string, platform string, deviceID string, token string) error {
	// 保存
	err := s.tRepo.Put(ctx, appID, userID, platform, deviceID, token)
	if err != nil {
		log.Errorm(ctx, "s.tRepo.Put", err)
		return err
	}

	// 全員のトピックに登録
	err = s.fRepo.SubscribeTopic(ctx, appID, config.TopicAll, []string{token})
	if err != nil {
		log.Errorm(ctx, "s.fRepo.SubscribeTopic", err)
		return err
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
		err = log.Warningc(ctx, http.StatusNotFound, "not exist token: app_id: %s, platform: %s, device_id: %s", appID, platform, deviceID)
		return err
	}

	// 削除
	err = s.tRepo.Delete(ctx, appID, userID, platform, deviceID)
	if err != nil {
		log.Errorm(ctx, "s.tRepo.Delete", err)
		return err
	}

	// 全員のトピックから削除
	err = s.fRepo.UnsubscribeTopic(ctx, appID, config.TopicAll, []string{token})
	if err != nil {
		log.Errorm(ctx, "s.fRepo.UnsubscribeTopic", err)
		return err
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
