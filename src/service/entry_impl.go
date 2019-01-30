package service

import (
	"context"

	"github.com/rabee-inc/push/src/lib/log"
	"github.com/rabee-inc/push/src/repository"
)

type entry struct {
	repo repository.Token
}

// Token ... 各プラットフォームで取得したプッシュ通知トークンを登録する
func (s *entry) Token(ctx context.Context, userID string, platform string, deviceID string, token string) error {
	err := s.repo.Put(ctx, userID, platform, deviceID, token)
	if err != nil {
		log.Errorm(ctx, "s.repo.Put", err)
		return err
	}
	return nil
}

// NewEntry ... サービスを作成する
func NewEntry(repo repository.Token) Entry {
	return &entry{
		repo: repo,
	}
}
