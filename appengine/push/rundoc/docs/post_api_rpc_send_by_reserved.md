# post_api_rpc_send_by_reserved

予約されたプッシュ通知を全員に送信する

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
    "method": "send_by_reserved",
    "params": {
        "app_id": "rec"
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
    "error": {
        "code": 400,
        "message": "method not found: send_by_reserved"
    }
}
```

