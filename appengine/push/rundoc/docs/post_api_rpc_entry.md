# post_api_rpc_entry

プッシュ通知に登録する

## Request

|ENV|URL|
|---|---|
|Local|`POST` http://localhost:8080/api/rpc|
|Staging|`POST` https://staging.appspot.com/api/rpc|
|Production|`POST` https://appspot.com/api/rpc|

```
Content-Type: application/json
Authorization: xxxxxxxxxx
```
```json
{
    "id": "0",
    "jsonrpc": "2.0",
    "method": "entry",
    "params": {
        "app_id": "test",
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
    "id": "0",
    "result": {
        "success": true
    }
}
```

