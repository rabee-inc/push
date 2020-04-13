# post_api_rpc_create_reserve

プッシュ通知を予約する

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
    "method": "create_reserve",
    "params": {
        "app_id": "rec",
        "message": {
            "body": "test_body_reserved",
            "title": "test_title_reserved"
        },
        "reserved_at": 1575802628763
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
            "id": "--SsKlwzcJUngjOMJCnl3v8x6QxVojf-uvXcGskCd5s",
            "ref": {
                "Parent": {
                    "Parent": {
                        "Parent": {
                            "Parent": null,
                            "Path": "projects/staging-push-rabee-jp/databases/(default)/documents/push_apps",
                            "ID": "push_apps"
                        },
                        "Path": "projects/staging-push-rabee-jp/databases/(default)/documents/push_apps/rec",
                        "ID": "rec"
                    },
                    "Path": "projects/staging-push-rabee-jp/databases/(default)/documents/push_apps/rec/reserves",
                    "ID": "reserves"
                },
                "Path": "projects/staging-push-rabee-jp/databases/(default)/documents/push_apps/rec/reserves/--SsKlwzcJUngjOMJCnl3v8x6QxVojf-uvXcGskCd5s",
                "ID": "--SsKlwzcJUngjOMJCnl3v8x6QxVojf-uvXcGskCd5s"
            },
            "message": {
                "title": "test_title_reserved",
                "body": "test_body_reserved",
                "data": null,
                "ios": null,
                "android": null,
                "web": null
            },
            "reserved_at": 1575802628763,
            "status": "reserved",
            "created_at": 1575802788198,
            "updated_at": 1575802788198
        }
    }
}
```

