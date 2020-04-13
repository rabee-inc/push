_STG_SCHEDULE     := * * * * *
_STG_DESCRIPTION  := 予約したプッシュ通知を送信する
_STG_SERVICE      := default
_STG_RELATIVE_URL := /worker/send/reserved?app_id=xxx
_STG_HTTP_METHOD  := GET
_STG_HTTP_HEADERS := Authorization=xxx
_STG_HTTP_BODY    :=
_STG_TIME_ZONE    := Asia/Tokyo
_STG_TIMEOUT      := 60s

_PRD_SCHEDULE     := * * * * *
_PRD_DESCRIPTION  := 予約したプッシュ通知を送信する
_PRD_SERVICE      := default
_PRD_RELATIVE_URL := /worker/send/reserved?app_id=xxx
_PRD_HTTP_METHOD  := GET
_PRD_HTTP_HEADERS := Authorization=xxx
_PRD_HTTP_BODY    :=
_PRD_TIME_ZONE    := Asia/Tokyo
_PRD_TIMEOUT      := 60s
