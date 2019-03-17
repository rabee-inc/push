# leave

ユーザーのプッシュ通知を解除する

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
    "method": "leave",
    "params": {
        "device_id": "test_device_id",
        "platform": "ios",
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

