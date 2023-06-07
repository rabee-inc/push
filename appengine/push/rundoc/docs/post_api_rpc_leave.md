# post_api_rpc_leave

プッシュ通知を解除する

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
    "method": "leave",
    "params": {
        "app_id": "test",
        "device_id": "test_device_id",
        "platform": "ios",
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

