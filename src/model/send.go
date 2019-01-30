package model

// TaskQueueParamSendUserIDs ... UserIDsでプッシュ通知を送信するパラメータ
type TaskQueueParamSendUserIDs struct {
	UserIDs []string `json:"user_ids"`
	Message *Message `json:"message"`
}

// TaskQueueParamSendUserID ... UserIDでプッシュ通知を送信するパラメータ
type TaskQueueParamSendUserID struct {
	UserID  string   `json:"user_id"`
	Message *Message `json:"message"`
}

// TaskQueueParamSendToken ... Tokenでプッシュ通知を送信するパラメータ
type TaskQueueParamSendToken struct {
	Token   string   `json:"token"`
	Message *Message `json:"message"`
}
