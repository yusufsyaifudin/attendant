package attendant

import (
	"encoding/json"
	"net/http"

	"github.com/stretchr/testify/mock"
)

// RequestMock local struct that holds mock instance
type RequestMock struct {
	mock.Mock
}

// _ ensures that RequestMock implements Request
var _ Request = (*RequestMock)(nil)

// ContentType mock Request.ContentType
func (r *RequestMock) ContentType() string {
	args := r.Called()
	return args.String(0)
}

// Bind mock Request.Bind
func (r *RequestMock) Bind(out interface{}) error {
	args := r.Called(out)

	resultByte, _ := json.Marshal(args.Get(0))
	_ = json.Unmarshal(resultByte, out)

	return args.Error(1)
}

// RawRequest mock Request.RawRequest
func (r *RequestMock) RawRequest() *http.Request {
	args := r.Called()
	return args.Get(0).(*http.Request)
}

// GetParam mock Request.GetParam
func (r *RequestMock) GetParam(param string) string {
	args := r.Called(param)
	return args.String(0)
}

// GetQueryParam mock Request.GetQueryParam
func (r *RequestMock) GetQueryParam(param string) string {
	args := r.Called(param)
	return args.String(0)
}

// Set mock Request.Set
func (r *RequestMock) Set(key string, val interface{}) error {
	args := r.Called(key, val)
	return args.Error(0)
}

// Get mock Request.Get
func (r *RequestMock) Get(key string) interface{} {
	args := r.Called(key)
	return args.Get(0)
}

// NewRequestMock returns new instance that implements Request interface
func NewRequestMock() *RequestMock {
	return &RequestMock{}
}
