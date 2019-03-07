package model

// TaskQueueParamSendUserIDs ... UserIDsでプッシュ通知を送信するパラメータ
type TaskQueueParamSendUserIDs struct {
	UserIDs []string `json:"user_ids" validate:"required"`
	Message *Message `json:"message"  validate:"required"`
}

// TaskQueueParamSendUserID ... UserIDでプッシュ通知を送信するパラメータ
type TaskQueueParamSendUserID struct {
	UserID  string   `json:"user_id" validate:"required"`
	Message *Message `json:"message" validate:"required"`
}

// TaskQueueParamSendToken ... Tokenでプッシュ通知を送信するパラメータ
type TaskQueueParamSendToken struct {
	Token   string   `json:"token"   validate:"required"`
	Message *Message `json:"message" validate:"required"`
}
