package http

import "errors"

// 지원하는 메서드
type HttpMethod string

const (
	GET    = HttpMethod("GET")
	POST   = HttpMethod("POST")
	DELETE = HttpMethod("DELETE")
	PUT    = HttpMethod("PUT")
)

func (h HttpMethod) ToString() string {
	return string(h)
}

type GetType string // Type

const (
	QUERY = GetType("query")
	URL   = GetType("url")
)

func (h GetType) ToString() string {
	return string(h)
}

func (h GetType) CheckType() error {
	switch h {
	case QUERY, URL:
		return nil
	default:
		return errors.New("invalid GetType")
	}
}
