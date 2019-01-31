package service

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/rabee-inc/push/test/config"
	"github.com/rabee-inc/push/test/model"
	"github.com/rabee-inc/push/test/repository"
)

const (
	apiTypeJSONRPC2 string = "JSONRPC2"
)

type jsonrpc2 struct {
	hRepo repository.HTTPClient

	apis    []*model.API
	url     string
	uri     string
	headers map[string]string

	ovStagingURL    string
	ovProductionURL string
}

func (s *jsonrpc2) Send(name string, method string, params map[string]interface{}) {
	ov := &model.APIOverview{
		Type: apiTypeRest,
		URL: &model.APIOverviewURL{
			Staging:    s.ovStagingURL,
			Production: s.ovProductionURL,
		},
		URI: s.uri,
	}

	s.headers["Content-Type"] = "application/json"
	hs := s.createHeadersString(s.headers)

	reqParams := map[string]interface{}{
		"jsonrpc": "2.0",
		"id":      "1",
		"method":  method,
		"params":  params,
	}

	jPs, sPs := s.createJSON(reqParams)

	req := &model.APIRequest{
		Method:  http.MethodPost,
		URI:     s.uri,
		Headers: hs,
		Params:  sPs,
	}

	status, body := s.hRepo.PostJSON(s.url+s.uri, jPs, s.headers)

	res := &model.APIResponse{
		StatusCode: status,
		Body:       body,
	}

	s.apis = append(s.apis, &model.API{
		Name:     name,
		Overview: ov,
		Request:  req,
		Response: res,
	})
}

func (s *jsonrpc2) createHeadersString(headers map[string]string) string {
	hs := ""
	for key, value := range headers {
		if key == "Authorization" {
			// AuthorizationHeaderの値を隠蔽
			hs += fmt.Sprintf("%s: %s%s\n", key, config.AuthorizationPrefix, "XXXXXXXXXX")
		} else {
			hs += fmt.Sprintf("%s: %s\n", key, value)
		}
	}
	return hs
}

func (s *jsonrpc2) createJSON(params map[string]interface{}) ([]byte, string) {
	dstJSON, err := json.Marshal(params)
	if err != nil {
		panic(err.Error())
	}

	out := new(bytes.Buffer)
	err = json.Indent(out, dstJSON, "", "    ")
	if err != nil {
		panic(err.Error())
	}
	dstStr := out.String()

	return dstJSON, dstStr
}

func (s *jsonrpc2) GetAPIs() []*model.API {
	return s.apis
}

// NewJSONRPC2 ... サービスを作成する
func NewJSONRPC2(
	hRepo repository.HTTPClient,
	url string,
	uri string,
	headers map[string]string,
	stagingURL string,
	productionURL string) JSONRPC2 {
	return &jsonrpc2{
		hRepo:           hRepo,
		apis:            []*model.API{},
		url:             url,
		uri:             uri,
		headers:         headers,
		ovStagingURL:    stagingURL,
		ovProductionURL: productionURL,
	}
}
