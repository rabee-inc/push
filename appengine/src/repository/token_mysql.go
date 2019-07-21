package repository

import (
	"context"

	"github.com/rabee-inc/push/appengine/src/lib/log"
	"github.com/rabee-inc/push/appengine/src/lib/mysql"
	"github.com/rabee-inc/push/appengine/src/lib/util"
	"github.com/rabee-inc/push/appengine/src/model"
)

type tokenMySQL struct {
	client *mysql.Client
}

// GetListByUserID ... ユーザーIDに紐づくトークンリストを取得する
func (r *tokenMySQL) GetListByUserID(ctx context.Context, appID string, userID string) ([]string, error) {
	dsts := []string{}
	var tokens []*model.TokenMySQL
	db := r.client.GetDB(ctx).
		Select([]string{
			"id",
			"app_id",
			"user_id",
			"platform",
			"device_id",
			"token",
			"created_at",
			"updated_at",
		}).
		Table("tokens").
		Where("app_id = ?", appID).
		Where("user_id = ?", userID).
		Find(&tokens)
	if err := mysql.HandleErrors(db); err != nil {
		log.Errorm(ctx, "db.Find", err)
		return dsts, err
	}
	for _, token := range tokens {
		dsts = append(dsts, token.Token)
	}
	return dsts, nil
}

// Put ... トークンを登録する
func (r *tokenMySQL) Put(ctx context.Context, appID string, userID string, platform string, deviceID string, token string) error {
	id := model.GenerateTokenID(appID, userID, platform, deviceID)
	now := util.TimeNowUnix()
	src := &model.TokenMySQL{
		ID:        id,
		AppID:     appID,
		UserID:    userID,
		Platform:  platform,
		DeviceID:  deviceID,
		Token:     token,
		CreatedAt: now,
		UpdatedAt: now,
	}
	db := mysql.Upsert(
		r.client.GetDB(ctx),
		"tokens",
		src,
		[]string{"token", "updated_at"})
	if err := mysql.HandleErrors(db); err != nil {
		log.Errorm(ctx, "mysql.Upsert", err)
		return err
	}
	return nil
}

// Delete ... トークンを削除する
func (r *tokenMySQL) Delete(ctx context.Context, appID string, userID string, platform string, deviceID string) error {
	db := r.client.GetDB(ctx).
		Table("tokens").
		Where("app_id = ?", appID).
		Where("user_id = ?", userID).
		Where("platform = ?", platform).
		Where("device_id = ?", deviceID).
		Delete(nil)
	if err := mysql.HandleErrors(db); err != nil {
		log.Errorm(ctx, "db.Delete", err)
		return err
	}
	return nil
}

// NewTokenMySQL ... リポジトリを作成する
func NewTokenMySQL(client *mysql.Client) Token {
	return &tokenMySQL{
		client: client,
	}
}
