package common

import (
	"github.com/bytedance/sonic"
	"log"
)

type jsonHandler struct {
	marshal   func(v interface{}) ([]byte, error)
	unmarshal func(data []byte, v interface{}) error
}

var JsonHandler jsonHandler

func init() {
	JsonHandler = jsonHandler{
		marshal:   sonic.Marshal,
		unmarshal: sonic.Unmarshal,
	}
}

// 에러 로깅을 위해 jsonHandler를 사용하도록 변경
func (j jsonHandler) Marshal(v interface{}) ([]byte, error) {
	bytes, err := j.marshal(v)

	if err != nil {
		log.Println("Failed to marshal json", err.Error())
		return nil, err
	}
	return bytes, nil
}

// 에러 로깅을 위해 jsonHandler를 사용하도록 변경
func (j jsonHandler) Unmarshal(data []byte, v interface{}) error {
	err := j.unmarshal(data, v)

	if err != nil {
		log.Println("Failed to unmarshal json", err.Error())
		return err
	}
	return nil
}
