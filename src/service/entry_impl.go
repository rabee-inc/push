package service

import (
	"context"

	"github.com/aikizoku/push/src/lib/log"
	"github.com/aikizoku/push/src/repository"
)

type entry struct {
	repo repository.Token
}

func (s *entry) Token(ctx context.Context, userID string, platform string, deviceID string, token string) error {
	err := s.repo.Put(ctx, userID, platform, deviceID, token)
	if err != nil {
		log.Errorf(ctx, "s.repo.Put error: %s", err.Error())
		return err
	}
	return nil
}

// NewEntry ... Entryを作成する
func NewEntry(repo repository.Token) Entry {
	return &entry{
		repo: repo,
	}
}
