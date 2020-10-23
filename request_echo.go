package attendant

import (
	"fmt"
	"net/http"
	"strings"

	"github.com/labstack/echo/v4"
)

// echoRequest holds implementation of Request using echo
type echoRequest struct {
	echoCtx echo.Context
}

// _ ensures echoRequest implements Request
var _ Request = (*echoRequest)(nil)

// newEchoRequest returns new echoRequest
func newEchoRequest(echoCtx echo.Context) *echoRequest {
	return &echoRequest{
		echoCtx: echoCtx,
	}
}

// ContentType get content type form header
func (r *echoRequest) ContentType() string {
	return r.echoCtx.Request().Header.Get("Content-Type")
}

// Bind binding payload into struct. Error returned when fail binding happened.
func (r *echoRequest) Bind(out interface{}) error {
	err := r.echoCtx.Bind(out)
	if err == nil {
		return nil
	}

	return fmt.Errorf(err.Error())
}

// RawRequest returns *http.Request
func (r *echoRequest) RawRequest() *http.Request {
	return r.echoCtx.Request()
}

// GetParam return URL parameter.
// For example, http://localhost:port/:my_param
// Then user access ttp://localhost:port/foo
// GetParam("my_param") will return foo
func (r *echoRequest) GetParam(key string) string {
	return r.echoCtx.Param(key)
}

// GetQueryParam return URL query parameter.
// For example, http://localhost:port/?query=foo
// GetQueryParam("query") will return foo
func (r *echoRequest) GetQueryParam(key string) string {
	return r.echoCtx.QueryParam(key)
}

// Set set value into echo context using key.
func (r *echoRequest) Set(key string, val interface{}) error {
	key = strings.TrimSpace(key)
	if key == "" {
		return fmt.Errorf("empty key on set")
	}

	r.echoCtx.Set(key, val)
	return nil
}

// Get get value that set in Set method
func (r *echoRequest) Get(key string) interface{} {
	return r.echoCtx.Get(key)
}
