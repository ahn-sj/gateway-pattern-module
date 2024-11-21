package client

import (
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
) {
	batchTime := cfg.Http.BatchTime

	if batchTime == 0 {
		batchTime = _defaultBatchTime
	}

	client := resty.New().
		SetJSONMarshaler(common.JsonHandler.Marshal).    // sonic Marshal
		SetJSONUnmarshaler(common.JsonHandler.Unmarshal) // sonic Unmarshal
}
