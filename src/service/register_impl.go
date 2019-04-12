package service

import (
	"context"

	"github.com/rabee-inc/push/src/lib/log"
	"github.com/rabee-inc/push/src/repository"
)

type register struct {
	repo repository.Token
}

// SetToken ... 各プラットフォームで取得したプッシュ通知トークンを登録する
func (s *register) SetToken(ctx context.Context, appID string, userID string, platform string, deviceID string, token string) error {
	err := s.repo.Put(ctx, appID, userID, platform, deviceID, token)
	if err != nil {
		log.Errorm(ctx, "s.repo.Put", err)
		return err
	}
	return nil
}

// DeleteToken ... 指定したプッシュ通知トークンを削除する
func (s *register) DeleteToken(ctx context.Context, appID string, userID string, platform string, deviceID string) error {
	err := s.repo.Delete(ctx, appID, userID, platform, deviceID)
	if err != nil {
		log.Errorm(ctx, "s.repo.Delete", err)
		return err
	}
	return nil
}

// NewRegister ... サービスを作成する
func NewRegister(repo repository.Token) Register {
	return &register{
		repo: repo,
	}
}
