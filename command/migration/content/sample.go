package content

import (
	"context"

	"cloud.google.com/go/firestore"
)

// Sample ... サンプルのシードデータ
type Sample struct {
	FCli *firestore.Client
}

// Generate ... シードデータを作成する
func (m *Sample) Generate(ctx context.Context) {
}
