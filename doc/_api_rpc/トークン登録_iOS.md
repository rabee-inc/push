# トークン登録_iOS

## Overview
```
【Type】
REST

【Endpoint】
https://push.staging.xxxx.rabee.jp
https://push.xxxx.rabee.jp

【URI】
/api/rpc
```

## Request
```json
【Method】
POST

【URL】
/api/rpc

【Headers】
Authorization: XXXXXXXXXX
Content-Type: application/json


【Params】
{
    "id": "1",
    "jsonrpc": "2.0",
    "method": "entry",
    "params": {
        "device_id": "test_device_id_ios",
        "platform": "ios",
        "token": "test_token_ios",
        "user_id": "test_user_id"
    }
}
```

## Response
```json
【StatusCode】
200

【Body】
{
    "jsonrpc": "2.0",
    "id": "1",
    "result": {
        "success": true
    }
}

```
