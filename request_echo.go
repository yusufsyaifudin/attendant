package attendant

import (
	"fmt"
	"net/http"
	"strings"

	"github.com/labstack/echo/v4"
)

type echoRequest struct {
	echoCtx echo.Context
}

func newEchoRequest(echoCtx echo.Context) Request {
	return &echoRequest{
		echoCtx: echoCtx,
	}
}

func (r *echoRequest) ContentType() string {
	return r.echoCtx.Request().Header.Get("Content-Type")
}

func (r *echoRequest) Bind(out interface{}) error {
	err := r.echoCtx.Bind(out)
	if err == nil {
		return nil
	}

	return fmt.Errorf(err.Error())
}

func (r *echoRequest) RawRequest() *http.Request {
	return r.echoCtx.Request()
}

func (r *echoRequest) GetParam(key string) string {
	return r.echoCtx.Param(key)
}

func (r *echoRequest) GetQueryParam(key string) string {
	return r.echoCtx.QueryParam(key)
}

func (r *echoRequest) Set(key string, val interface{}) error {
	key = strings.TrimSpace(key)
	if key == "" {
		return fmt.Errorf("empty key on set")
	}

	r.echoCtx.Set(key, val)
	return nil
}

func (r *echoRequest) Get(key string) interface{} {
	return r.echoCtx.Get(key)
}
