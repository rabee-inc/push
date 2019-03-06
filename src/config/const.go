package config

const (
	// CollectionUsers ... FirestoreのUsersコレクション
	CollectionUsers = "push_users"
	// CollectionTokens ... FirestoreのTokensコレクション
	CollectionTokens = "tokens"

	// KindPushToken ... DatastoreのPushTokenカインド
	KindPushToken = "PushTokenDatastore"

	// QueueSendUser ... UserID変換処理のQueue
	QueueSendUser = "PushSendUser"
	// QueueSendToken ... Tokenから通知を送信するQueue
	QueueSendToken = "PushSendToken"
)
