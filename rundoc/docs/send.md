# send

ユーザーにプッシュ通知を送信する

|ENV|URL|
|---|---|
|Local|http://localhost:8080|
|Staging|https://push-dot-staging.appspot.com|
|Production|https://push-dot.appspot.com|

## Request

```
POST
```
```
/api/rpc
```
```
Authorization: xxxxxxxxxx
Content-Type: application/json
```
```json
{
    "id": "1",
    "jsonrpc": "2.0",
    "method": "send",
    "params": {
        "message": {
            "android": {
                "click_action": "任意のaction名",
                "sound": "好きなサウンドファイル名（空でデフォルト音）",
                "tag": "タグ"
            },
            "body": "テストボディ",
            "data": {
                "hoge": "任意のデータ"
            },
            "ios": {
                "badge": 1,
                "sound": "好きなサウンドファイル名（空でデフォルト音）"
            },
            "title": "テストタイトル",
            "web": {
                "icon": "アイコン"
            }
        },
        "user_ids": [
            "test_user_id"
        ]
    }
}
```

## Response

```
200
```
```json
{
    "jsonrpc": "2.0",
    "id": "1",
    "result": {
        "success": true
    }
}
```

