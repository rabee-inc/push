# これは何？
iOS, Android, Webにプッシュ通知を送信するサーバーです。
プッシュ通知送信機能を導入したいプロジェクトに個別でデプロイして使用してください。

# セットアップ

## 実行
```bash
make run app=push
```

## デプロイ
```bash
make deploy app=push
make deploy-production app=push
```

# API
JSONRPC2.0を使用しています。
詳細はdoc内を参照してください。
