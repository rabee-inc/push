# post_api_rpc_list_reserve

プッシュ通知の予約リストを取得する

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
    "method": "list_reserve",
    "params": {
        "app_id": "test",
        "cursor": "",
        "limit": 10,
        "reserve_id": ""
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
        "reserves": [
            {
                "id": "OhyeVdVpin8nNYsQKGql",
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
                "created_at": 1686113722998,
                "updated_at": 1686113722998
            },
            {
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
        ],
        "next_cursor": ""
    }
}
```

