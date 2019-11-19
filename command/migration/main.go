package main

import (
	"context"

	"github.com/rabee-inc/push/appengine/default/src/lib/cloudfirestore"
	"github.com/rabee-inc/push/command/lib"
	"github.com/rabee-inc/push/command/migration/content"
)

func main() {
	ctx := context.Background()

	// env.jsonの読み込み
	pID := lib.GetProjectID(lib.Staging)

	// Inject
	fCli := cloudfirestore.NewClient(pID)

	u := &content.Sample{
		FCli: fCli,
	}

	// 実行
	u.Generate(ctx)
}
