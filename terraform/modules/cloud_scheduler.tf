resource "google_cloud_scheduler_job" "push_send_by_reserved" {
  project = var.project_id
  name = "push-send-by-reserved"
  schedule = "* * * * *"
  description = "予約したプッシュ通知を送信する"
  region = var.region
  time_zone = var.time_zone
  attempt_deadline = "60s"

  app_engine_http_target {
    http_method = "GET"

    app_engine_routing {
      service  = "push"
    }

    relative_uri = "/worker/send/reserved?app_id=xxx"

    headers = {
        Authorization = var.internal_auth_token
    }
  }
}
