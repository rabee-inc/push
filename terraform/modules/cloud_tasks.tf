resource "google_cloud_tasks_queue" "push_send_user" {
  project = var.project_id
  name = "PushSendUser"
  location = var.region

  rate_limits {
    max_dispatches_per_second = 100
  }

  retry_config {
    max_attempts = 1
  }
}
