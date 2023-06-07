# get_worker_send_by_reserved

予約されたプッシュ通知を全員に送信する

## Request

|ENV|URL|
|---|---|
|Local|`GET` http://localhost:8080/worker/send/reserved|
|Staging|`GET` https://staging.appspot.com/worker/send/reserved|
|Production|`GET` https://appspot.com/worker/send/reserved|

```
Authorization: xxxxxxxxxx
Content-Type: application/json
```
```json
{
    "app_id": "test"
}
```

## Response

```
Status 200
```
```json
{
    "status": 200
}
```

