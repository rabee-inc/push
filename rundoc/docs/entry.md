# entry

ユーザーをプッシュ通知に登録する

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
    "method": "entry",
    "params": {
        "app_id": "test_app_id",
        "device_id": "test_device_id",
        "platform": "ios",
        "token": "test_token",
        "user_id": "test_user_id"
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

