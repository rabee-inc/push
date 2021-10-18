package model

// Message ... プッシュ通知メッセージ
type Message struct {
	Title   string            `json:"title"   firestore:"title"`
	Body    string            `json:"body"    firestore:"body"`
	Data    map[string]string `json:"data"    firestore:"data"`
	IOS     *MessageIOS       `json:"ios"     firestore:"ios"`
	Android *MessageAndroid   `json:"android" firestore:"android"`
	Web     *MessageWeb       `json:"web"     firestore:"web"`
}

// MessageIOS ... プッシュ通知メッセージ(iOS独自部分)
type MessageIOS struct {
	Badge    int    `json:"badge,omitempty"     firestore:"badge"`
	Sound    string `json:"sound,omitempty"     firestore:"sound"`
	ImageURL string `json:"image_url,omitempty" firestore:"image_url"`
}

// MessageAndroid ... プッシュ通知メッセージ(Android独自部分)
type MessageAndroid struct {
	ClickAction string `json:"click_action,omitempty" firestore:"click_action"`
	Sound       string `json:"sound,omitempty"        firestore:"sound"`
	Tag         string `json:"badge,omitempty"        firestore:"badge"`
}

// MessageWeb ... プッシュ通知メッセージ(Web独自部分)
type MessageWeb struct {
	Icon string `json:"icon,omitempty" firestore:"icon"`
}
