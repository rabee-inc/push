# これは何？
iOS, Android, Webにプッシュ通知を送信するサーバーです。

各サービスのユーザーID毎に各プラットフォームに良い感じに負荷分散しながら送信してくれます。

プッシュ通知送信機能を導入したいプロジェクトに個別でデプロイして使用してください。

# 対応状況

## 対応しているデータベース
- Cloud Datastore
- Cloud Firestore
- MySQL

## 機能
- ユーザー＆Token登録
- 即時送信
- 予約送信 ToBe...
- 定期送信 ToBe...

# セットアップ

## 準備
```bash
cp env.example.json env.json
dep ensure
```

## Docker(LocalのMySQLを使う場合)
```bash
# build
docker-compose build

# start
docker-compose up -d

# stop
docker-compose down
```

## 実行
```bash
make run app=default
```

## デプロイ
```bash
make deploy app=default
make deploy-production app=default
```

# API
JSONRPC2.0を使用しています。
詳細はrundoc/docs内を参照してください。
