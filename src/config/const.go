package config

const (
	// CollectionUsers ... FirestoreのUsersコレクション
	CollectionUsers = "PushUsers"
	// CollectionTokens ... FirestoreのTokensコレクション
	CollectionTokens = "Tokens"

	// KindPushToken ... DatastoreのPushTokenカインド
	KindPushToken = "PushTokenDatastore"

	// QueueSendUser ... UserID変換処理のQueue
	QueueSendUser = "PushSendUser"
	// QueueSendToken ... Tokenから通知を送信するQueue
	QueueSendToken = "PushSendToken"
)
