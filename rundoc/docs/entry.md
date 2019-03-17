# entry

ユーザーをプッシュ通知に登録する

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
Content-Type: application/json
Authorization: xxxxxxxxxx
```
```json
{
    "id": "1",
    "jsonrpc": "2.0",
    "method": "entry",
    "params": {
        "device_id": "test_device_id",
        "platform": "ios",
        "token": "test_token",
        "user_id": "test_user_id"
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

