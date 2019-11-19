package model

// CloudTasksParamSendUserIDs ... UserIDsでプッシュ通知を送信するパラメータ
type CloudTasksParamSendUserIDs struct {
	AppID   string   `json:"app_id"   validate:"required"`
	UserIDs []string `json:"user_ids" validate:"required"`
	Message *Message `json:"message"  validate:"required"`
}

// CloudTasksParamSendUserID ... UserIDでプッシュ通知を送信するパラメータ
type CloudTasksParamSendUserID struct {
	AppID   string   `json:"app_id"  validate:"required"`
	UserID  string   `json:"user_id" validate:"required"`
	Message *Message `json:"message" validate:"required"`
}

// CloudTasksParamSendToken ... Tokenでプッシュ通知を送信するパラメータ
type CloudTasksParamSendToken struct {
	AppID   string   `json:"app_id"  validate:"required"`
	Token   string   `json:"token"   validate:"required"`
	Message *Message `json:"message" validate:"required"`
}
