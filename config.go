package attendant

import (
	"time"

	"github.com/opentracing/opentracing-go"
	"go.uber.org/zap"
)

// Config is a configuration needed by Server to run
type Config struct {
	EnableProfiling bool
	ListenAddress   int
	WriteTimeout    time.Duration
	ReadTimeout     time.Duration

	ZapLogger   *zap.Logger
	OpenTracing opentracing.Tracer
}
