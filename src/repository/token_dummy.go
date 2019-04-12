package repository

import (
	"context"

	"github.com/rabee-inc/push/src/lib/log"
)

type tokenDummy struct {
}

// GetListByUserID ... ユーザーIDに紐づくトークンリストを取得する
func (r *tokenDummy) GetListByUserID(ctx context.Context, appID string, userID string) ([]string, error) {
	log.Debugf(ctx, "call token get list by user id")
	return []string{}, nil
}

// Put ... トークンを登録する
func (r *tokenDummy) Put(ctx context.Context, appID string, userID string, platform string, deviceID string, token string) error {
	log.Debugf(ctx, "call token put")
	return nil
}

// Delete ... トークンを削除する
func (r *tokenDummy) Delete(ctx context.Context, appID string, userID string, platform string, deviceID string) error {
	log.Debugf(ctx, "call token delete")
	return nil
}

// NewTokenDummy ... リポジトリを作成する
func NewTokenDummy() Token {
	return &tokenDummy{}
}
