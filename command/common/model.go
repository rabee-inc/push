package common

// Env ... 環境変数ファイルの定義
type Env struct {
	Apps        []string `json:"apps"`
	Credentials EnvData  `json:"credentials"`
	Values      EnvData  `json:"values"`
}

// EnvData ... 環境変数ファイルの環境毎のデータの定義
type EnvData struct {
	Local      map[string]string `json:"local"`
	Staging    map[string]string `json:"staging"`
	Production map[string]string `json:"production"`
}

// ProjectIDs ... 各環境のProjectIDの定義
type ProjectIDs struct {
	Local      string
	Staging    string
	Production string
}
