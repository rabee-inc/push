# post_api_rpc_create_reserve

プッシュ通知を予約する

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
    "method": "create_reserve",
    "params": {
        "app_id": "test",
        "message": {
            "body": "test_body_reserved",
            "title": "test_title_reserved"
        },
        "reserved_at": 1686113627892,
        "unmanaged": false
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
            "reserved_at": 1686113627892,
            "unmanaged": false,
            "status": "reserved",
            "created_at": 1686113768330,
            "updated_at": 1686113768330
        }
    }
}
```

