# send

ユーザーにプッシュ通知を送信する

## Request

|ENV|URL|
|---|---|
|Local|`POST` http://localhost:8080/api/rpc|
|Staging|`POST` https://push-dot-staging.appspot.com/api/rpc|
|Production|`POST` https://push-dot.appspot.com/api/rpc|

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
        "app_id": "test_app_id",
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
Status 200
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

