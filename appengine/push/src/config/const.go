package config

const (
	QueueSendUser = "PushSendUser"
	TopicAll      = "all"
)

// ReserveStatus ... 予約ステータス
type ReserveStatus string

const (
	ReserveStatusReserved   ReserveStatus = "reserved"
	ReserveStatusCanceled   ReserveStatus = "canceled"
	ReserveStatusProcessing ReserveStatus = "processing"
	ReserveStatusFailure    ReserveStatus = "failure"
	ReserveStatusSuccess    ReserveStatus = "success"
)
