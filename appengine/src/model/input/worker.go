package input

import "github.com/rabee-inc/push/appengine/src/model"

type WorkerSendUsers struct {
	AppID   string         `json:"app_id"   validate:"required"`
	UserIDs []string       `json:"user_ids" validate:"required"`
	PushID  string         `json:"push_id"  validate:"required"`
	Message *model.Message `json:"message"  validate:"required"`
}

type WorkerSendUser struct {
	AppID   string         `json:"app_id"  validate:"required"`
	UserID  string         `json:"user_id" validate:"required"`
	PushID  string         `json:"push_id" validate:"required"`
	Message *model.Message `json:"message" validate:"required"`
}
