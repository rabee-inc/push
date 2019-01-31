package scenario

import "github.com/rabee-inc/push/test/service"

type normal struct {
	dSvc service.Document
	rSvc service.Rest
	jSvc service.JSONRPC2
}

func (s *normal) Run() {
	// リクエスト送る
	s.Send()

	// リクエストとレスポンスをまとめる
	apis := s.jSvc.GetAPIs()

	// お掃除
	s.dSvc.RemoveAll()

	// フォーマットに適用
	s.dSvc.Distributes("template/api.tmpl", apis)
}

func (s *normal) Send() {
	s.jSvc.Send("トークン登録_iOS", "entry", map[string]interface{}{
		"user_id":   "test_user_id",
		"platform":  "ios",
		"device_id": "test_device_id_ios",
		"token":     "test_token_ios",
	})
	s.jSvc.Send("トークン登録_Android", "entry", map[string]interface{}{
		"user_id":   "test_user_id",
		"platform":  "android",
		"device_id": "test_device_id_android",
		"token":     "test_token_android",
	})
	s.jSvc.Send("トークン登録_Web", "entry", map[string]interface{}{
		"user_id":   "test_user_id",
		"platform":  "web",
		"device_id": "test_device_id_web",
		"token":     "test_token_web",
	})

	s.jSvc.Send("即時送信", "send", map[string]interface{}{
		"user_ids": []string{"test_user_id"},
		"message": map[string]interface{}{
			"title": "テストタイトル",
			"body":  "テストボディ",
			"data": map[string]string{
				"hoge": "任意のデータ",
				"fuga": "12345",
			},
			"ios": map[string]interface{}{
				"badge": 1,
				"sound": "好きなサウンドファイル名（空でデフォルト音）",
			},
			"android": map[string]interface{}{
				"click_action": "任意のaction名",
				"sound":        "好きなサウンドファイル名（空でデフォルト音）",
				"tag":          "タグ",
			},
			"web": map[string]interface{}{
				"icon": "アイコン",
			},
		},
	})
}

// NewNormal ... Normalを作成する
func NewNormal(
	dSvc service.Document,
	rSvc service.Rest,
	jSvc service.JSONRPC2) Interfaces {
	return &normal{
		dSvc: dSvc,
		rSvc: rSvc,
		jSvc: jSvc,
	}
}
