package attendant

import (
	"net/http"
)

// Request represents an api request
type Request interface {
	ContentType() string
	Bind(out interface{}) error
	RawRequest() *http.Request
	GetParam(string) string
	GetQueryParam(string) string

	// Add Value
	Set(key string, val interface{}) error
	Get(key string) interface{}
}
