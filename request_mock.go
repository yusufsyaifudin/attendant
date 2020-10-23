package attendant

import (
	"encoding/json"
	"net/http"
	"time"

	"github.com/stretchr/testify/mock"
)

// requestMock local struct that holds mock instance
type requestMock struct {
	mock.Mock

	startTime time.Time
	traceID   string
}

// _ ensures that requestMock implements Request
var _ Request = (*requestMock)(nil)

// ContentType mock Request.ContentType
func (r *requestMock) ContentType() string {
	args := r.Called()
	return args.String(0)
}

// Bind mock Request.Bind
func (r *requestMock) Bind(out interface{}) error {
	args := r.Called(out)

	resultByte, _ := json.Marshal(args.Get(0))
	_ = json.Unmarshal(resultByte, out)

	return args.Error(1)
}

// RawRequest mock Request.RawRequest
func (r *requestMock) RawRequest() *http.Request {
	args := r.Called()
	return args.Get(0).(*http.Request)
}

// GetParam mock Request.GetParam
func (r *requestMock) GetParam(param string) string {
	args := r.Called(param)
	return args.String(0)
}

// GetQueryParam mock Request.GetQueryParam
func (r *requestMock) GetQueryParam(param string) string {
	args := r.Called(param)
	return args.String(0)
}

// Set mock Request.Set
func (r *requestMock) Set(key string, val interface{}) error {
	args := r.Called(key, val)
	return args.Error(0)
}

// Get mock Request.Get
func (r *requestMock) Get(key string) interface{} {
	args := r.Called(key)
	return args.Get(0)
}

// NewRequestMock returns new instance that implements Request interface
func NewRequestMock() *requestMock {
	return &requestMock{}
}
