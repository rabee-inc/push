package model

// SendUserIDs ... UserIDsでプッシュ通知を送信するパラメータ
type SendUserIDs struct {
	UserIDs []string `json:"user_ids"`
	Message *Message `json:"message"`
}

// SendUserID ... UserIDでプッシュ通知を送信するパラメータ
type SendUserID struct {
	UserID  string   `json:"user_id"`
	Message *Message `json:"message"`
}

// SendToken ... Tokenでプッシュ通知を送信するパラメータ
type SendToken struct {
	Token   string   `json:"token"`
	Message *Message `json:"message"`
}
