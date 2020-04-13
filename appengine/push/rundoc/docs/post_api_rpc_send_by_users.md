# post_api_rpc_send_by_users

プッシュ通知を複数ユーザーに送信する

## Request

|ENV|URL|
|---|---|
|Local|`POST` http://localhost:8080/api/rpc|
|Staging|`POST` https://staging.appspot.com/api/rpc|
|Production|`POST` https://appspot.com/api/rpc|

```
Authorization: xxxxxxxxxx
Content-Type: application/json
```
```json
{
    "id": "0",
    "jsonrpc": "2.0",
    "method": "send_by_users",
    "params": {
        "app_id": "rec",
        "message": {
            "body": "test_body_users",
            "title": "test_title_users"
        },
        "user_ids": [
            "wnuLkLqflFWWtuhu7dvML0sy6Rh1"
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
    "id": "0",
    "result": {
        "success": true
    }
}
```

