# post_api_rpc_update_reserve

プッシュ通知の予約を編集する

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
    "method": "update_reserve",
    "params": {
        "app_id": "test",
        "message": {
            "body": "test_body_reserved",
            "title": "test_title_reserved"
        },
        "reserve_id": "BZoYtjvGcL3w5T9SkYCw",
        "reserved_at": 1700000000000,
        "status": "canceled"
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
        "reserve": {
            "id": "BZoYtjvGcL3w5T9SkYCw",
            "user_ids": [],
            "message": {
                "title": "test_title_reserved",
                "body": "test_body_reserved",
                "data": null,
                "ios": null,
                "android": null,
                "web": null
            },
            "reserved_at": 1700000000000,
            "unmanaged": false,
            "status": "canceled",
            "created_at": 1686113768330,
            "updated_at": 1686114673712
        }
    }
}
```

