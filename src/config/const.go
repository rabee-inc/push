package config

const (
	// CollectionApps ... FirestoreのAppsコレクション
	CollectionApps = "push_apps"
	// CollectionUsers ... FirestoreのUsersコレクション
	CollectionUsers = "users"
	// CollectionTokens ... FirestoreのTokensコレクション
	CollectionTokens = "tokens"

	// KindPushToken ... DatastoreのPushTokenカインド
	KindPushToken = "PushToken"

	// QueueSendUser ... UserID変換処理のQueue
	QueueSendUser = "PushSendUser"
	// QueueSendToken ... Tokenから通知を送信するQueue
	QueueSendToken = "PushSendToken"
)
