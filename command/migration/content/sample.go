package content

import (
	"context"

	"cloud.google.com/go/firestore"
)

// Sample ... サンプルのシードデータ
type Sample struct {
	FCli *firestore.Client
}

// Migrate ... マイグレーションを実行する
func (m *Sample) Migrate(ctx context.Context) {
}
