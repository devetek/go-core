// Package jsonlog builds zap loggers that write one JSON object per line.
//
// Unlike jsonerr, the returned *zap.Logger keeps no per-call state, so it is
// safe to share between goroutines, and it reports the caller of the log call.
// Every entry is written: there is no sampling.
package jsonlog

import (
	"os"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

type config struct {
	level   zapcore.Level
	output  zapcore.WriteSyncer
	service string
}

// Option configures New.
type Option func(*config)

// WithLevel sets the minimum level written. The default is info.
func WithLevel(level zapcore.Level) Option {
	return func(c *config) {
		c.level = level
	}
}

// WithService adds a "service" field to every entry.
func WithService(name string) Option {
	return func(c *config) {
		c.service = name
	}
}

// WithOutput sets where entries are written. The default is stdout.
func WithOutput(output zapcore.WriteSyncer) Option {
	return func(c *config) {
		c.output = output
	}
}

// New returns a JSON logger. Entries carry "time" (RFC 3339 with
// nanoseconds), "level", "caller" (file:line of the log call), "msg" and the
// fields passed to the call. Stack traces are only added from DPanic up.
func New(options ...Option) *zap.Logger {
	cfg := &config{
		level:  zapcore.InfoLevel,
		output: zapcore.Lock(os.Stdout),
	}
	for _, option := range options {
		option(cfg)
	}

	encoderConfig := zap.NewProductionEncoderConfig()
	encoderConfig.TimeKey = "time"
	encoderConfig.EncodeTime = zapcore.RFC3339NanoTimeEncoder
	encoderConfig.EncodeDuration = zapcore.MillisDurationEncoder

	core := zapcore.NewCore(zapcore.NewJSONEncoder(encoderConfig), cfg.output, cfg.level)
	logger := zap.New(core, zap.AddCaller(), zap.AddStacktrace(zapcore.DPanicLevel))
	if cfg.service != "" {
		logger = logger.With(zap.String("service", cfg.service))
	}

	return logger
}
