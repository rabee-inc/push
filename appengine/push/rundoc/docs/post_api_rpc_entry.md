# post_api_rpc_entry

プッシュ通知に登録する

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
    "method": "entry",
    "params": {
        "app_id": "rec",
        "device_id": "75BF4085-F382-4600-A949-AF1A27EF5F11",
        "platform": "ios",
        "token": "eo4L3IORW-4:APA91bHydj7jiMSbB4vKt1ht6fIyjHyHqybwb1P5h55li_b21BjBHIzCjv2FnRGnmmZ8_pOWVhDwHT1ef1Bqus2Kp3y8nQmyd8sv7fUjBIOMaC49_B0jDYOXcNg3_RgOSaRgXCUEbHPX",
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

