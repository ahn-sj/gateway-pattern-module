package client

import (
	"fmt"
	"gateway-module/common"
	"gateway-module/config"
	"github.com/confluentinc/confluent-kafka-go/kafka"
	"github.com/go-resty/resty/v2"
)

const (
	_defaultBatchTime = 2
)

type HttpClient struct {
	client *resty.Client
	cfg    config.App

	producer kafka.Producer
}

func NewHttpClient(
	cfg config.App,
	producer map[string]kafka.Producer,
) HttpClient {
	batchTime := cfg.Producer.BatchTime

	if batchTime == 0 {
		batchTime = _defaultBatchTime
	}

	if cfg.Http.BaseUrl == "" {
		panic("BaseUrl is required")
	}

	client := resty.New().
		SetJSONMarshaler(common.JsonHandler.Marshal).     // sonic Marshal
		SetJSONUnmarshaler(common.JsonHandler.Unmarshal). // sonic Unmarshal
		SetBaseURL(cfg.Http.BaseUrl)
	return HttpClient{
		cfg:      cfg,
		client:   client,
		producer: producer[cfg.App.Name],
	}
}

func (h HttpClient) Get(url string, router config.Router) (interface{}, error) {
	var buffer interface{} // 외부 호출한 결과를 받을 변수
	var err error
	var req *resty.Request
	var resp *resty.Response

	// defer // 함수가 종료할 때 실행

	req = getRequest(h.client, router).SetResult(&buffer)
	resp, err = req.Get(url)

	fmt.Println(resp)

	if err != nil {
		return nil, err
	}
	return buffer, nil
}

func (h HttpClient) Post(url string, router config.Router, requestBody interface{}) (interface{}, error) {
	var buffer interface{} // 외부 호출한 결과를 받을 변수
	var err error
	var req *resty.Request
	var resp *resty.Response

	// defer // 함수가 종료할 때 실행

	req = getRequest(h.client, router).SetResult(&buffer).SetBody(requestBody)
	resp, err = req.Post(url)

	fmt.Println(resp)

	if err != nil {
		return nil, err
	}
	return buffer, nil
}

func (h HttpClient) Put(url string, router config.Router, requestBody interface{}) (interface{}, error) {
	var buffer interface{} // 외부 호출한 결과를 받을 변수
	var err error
	var req *resty.Request
	var resp *resty.Response

	// defer // 함수가 종료할 때 실행

	req = getRequest(h.client, router).SetResult(&buffer).SetBody(requestBody)
	resp, err = req.Put(url)

	fmt.Println(resp)

	if err != nil {
		return nil, err
	}
	return buffer, nil
}

func (h HttpClient) Delete(url string, router config.Router, requestBody interface{}) (interface{}, error) {
	var buffer interface{} // 외부 호출한 결과를 받을 변수
	var err error
	var req *resty.Request
	var resp *resty.Response

	// defer // 함수가 종료할 때 실행

	req = getRequest(h.client, router).SetResult(&buffer).SetBody(requestBody)
	resp, err = req.Delete(url)

	fmt.Println(resp)

	if err != nil {
		return nil, err
	}
	return buffer, nil
}

func getRequest(client *resty.Client, router config.Router) *resty.Request {
	// h.client.R().SetAuthScheme().SetAuthToken().SetHeaders()
	req := client.R().EnableTrace()

	if router.Auth != nil {
		if len(router.Auth.Key) != 0 {
			req.SetAuthScheme(router.Auth.Key)
		}
		req.SetAuthScheme(router.Auth.Token)
	}
	if router.Header != nil {
		req.SetHeaders(router.Header)
	}
	return req
}
