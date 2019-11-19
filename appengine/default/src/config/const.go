package config

const (
	// QueueSendUser ... UserID変換処理のQueue
	QueueSendUser = "PushSendUser"
	// QueueSendToken ... Tokenから通知を送信するQueue
	QueueSendToken = "PushSendToken"

	// TopicAll ... 全員が登録されているトピック
	TopicAll = "all"
)
