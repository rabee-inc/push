package model

// CloudTasksParamSendUsers ... Usersにプッシュ通知を送信するパラメータ
type CloudTasksParamSendUsers struct {
	AppID   string   `json:"app_id"   validate:"required"`
	UserIDs []string `json:"user_ids" validate:"required"`
	Message *Message `json:"message"  validate:"required"`
}

// CloudTasksParamSendUser ... Userにプッシュ通知を送信するパラメータ
type CloudTasksParamSendUser struct {
	AppID   string   `json:"app_id"  validate:"required"`
	UserID  string   `json:"user_id" validate:"required"`
	Message *Message `json:"message" validate:"required"`
}
