package config

const (
	// QueueSendUser ... UserID変換処理のQueue
	QueueSendUser = "PushSendUser"

	// TopicAll ... 全員が登録されているトピック
	TopicAll = "all"
)

// ReserveStatus ... 予約ステータス
type ReserveStatus string

const (
	// ReserveStatusReserved ... 予約ステータス: 予約中
	ReserveStatusReserved ReserveStatus = "reserved"
	// ReserveStatusCanceled ... 予約ステータス: キャンセル
	ReserveStatusCanceled ReserveStatus = "canceled"
	// ReserveStatusProcessing ... 予約ステータス: 処理中
	ReserveStatusProcessing ReserveStatus = "processing"
	// ReserveStatusFailure ... 予約ステータス: 送信失敗
	ReserveStatusFailure ReserveStatus = "failure"
	// ReserveStatusSuccess ... 予約ステータス: 送信成功
	ReserveStatusSuccess ReserveStatus = "success"
)
