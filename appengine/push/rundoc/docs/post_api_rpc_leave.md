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
        "app_id": "rec",
        "device_id": "75BF4085-F382-4600-A949-AF1A27EF5F11",
        "platform": "ios",
        "user_id": "wnuLkLqflFWWtuhu7dvML0sy6Rh1"
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

