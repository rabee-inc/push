package config

import "time"

const (
	// HTTPRequestTimeout ... デフォルトのHTTPリクエストタイムアウト
	HTTPRequestTimeout time.Duration = 20

	// CollectionUsers ...
	CollectionUsers = "PushUsers"
	// CollectionTokens ...
	CollectionTokens = "Tokens"
)
