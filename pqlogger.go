package penquin

import (
	"fmt"
	"net/url"
	"time"

	"github.com/pkg/errors"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"gopkg.in/natefinch/lumberjack.v2"
)

var (
	lg      *zap.Logger
	sugerLg *zap.SugaredLogger
)

func init() {
	l, err := NewLogger()
	if err != nil {
		fmt.Println("init logger failed")
		return
	}
	lg = l
	sugerLg = l.Sugar()
}

func Debug(msg string, filed ...zap.Field) {
	lg.Debug(msg, filed...)
}

func Info(msg string, filed ...zap.Field) {
	lg.Info(msg, filed...)
}

func Warn(msg string, filed ...zap.Field) {
	lg.Warn(msg, filed...)
}

func Error(msg string, filed ...zap.Field) {
	lg.Error(msg, filed...)
}

func DPanic(msg string, filed ...zap.Field) {
	lg.DPanic(msg, filed...)
}

func Panic(msg string, filed ...zap.Field) {
	lg.Panic(msg, filed...)
}

func Fatal(msg string, filed ...zap.Field) {
	lg.Fatal(msg, filed...)
}

func Sync() {
	lg.Sync()
}

func SDebug(tpl string, args ...interface{}) {
	sugerLg.Debugf(tpl, args...)
}

func SInfo(tpl string, args ...interface{}) {
	sugerLg.Infof(tpl, args...)
}

func SWarn(tpl string, args ...interface{}) {
	sugerLg.Warnf(tpl, args...)
}

func SError(tpl string, args ...interface{}) {
	sugerLg.Errorf(tpl, args...)
}

func SDPanic(tpl string, args ...interface{}) {
	sugerLg.DPanicf(tpl, args...)
}

func SPanic(tpl string, args ...interface{}) {
	sugerLg.Panicf(tpl, args...)
}

func SFatal(tpl string, args ...interface{}) {
	sugerLg.Fatalf(tpl, args...)
}

type LogOptions struct {
	Path            string
	MaxSize         int
	MaxBackups      int
	MaxAge          int
	Stdout          bool
	EncodingConsole bool
}

var defaultLogOptions = LogOptions{
	Path:            "./logs/sm.log",
	MaxSize:         1024,
	MaxBackups:      50,
	MaxAge:          3,
	Stdout:          false,
	EncodingConsole: false,
}

func WithPath(v string) logOptionsFunc {
	return func(o *LogOptions) {
		o.Path = v
	}
}

func WithMaxSize(v int) logOptionsFunc {
	return func(o *LogOptions) {
		o.MaxSize = v
	}
}

func WithMaxBackups(v int) logOptionsFunc {
	return func(o *LogOptions) {
		o.MaxBackups = v
	}
}

func WithMaxAge(v int) logOptionsFunc {
	return func(o *LogOptions) {
		o.MaxAge = v
	}
}

func WithStdout(v bool) logOptionsFunc {
	return func(o *LogOptions) {
		o.Stdout = v
	}
}

func WithEncodingConsole(v bool) logOptionsFunc {
	return func(o *LogOptions) {
		o.EncodingConsole = v
	}
}

type logOptionsFunc func(*LogOptions)

type logRotationConfig struct {
	*lumberjack.Logger
}

func (l logRotationConfig) Sync() error {
	return nil
}

func NewLogger(opt ...logOptionsFunc) (*zap.Logger, error) {
	opts := defaultLogOptions
	for _, o := range opt {
		o(&opts)
	}
	cfg := logRotationConfig{
		&lumberjack.Logger{
			MaxSize:    opts.MaxSize,
			MaxBackups: opts.MaxBackups,
			MaxAge:     opts.MaxAge,
		},
	}
	if err := zap.RegisterSink("rotate", func(u *url.URL) (zap.Sink, error) {
		cfg.Filename = u.Path[1:]
		return &cfg, nil
	}); err != nil {
		return nil, errors.Wrap(err, "")
	}
	zap.AddCallerSkip(1)
	zapCfg := zap.NewProductionConfig()
	zapCfg.EncoderConfig.EncodeTime = zapcore.ISO8601TimeEncoder
	if opts.EncodingConsole {
		zapCfg.Encoding = "console"
		zapCfg.EncoderConfig.EncodeLevel = func(level zapcore.Level, enc zapcore.PrimitiveArrayEncoder) {
			enc.AppendString(fmt.Sprintf("%s %s", level.CapitalString(), time.Now().Format("2006-01-02 15:04:05.999")))
		}
		zapCfg.EncoderConfig.EncodeTime = func(t time.Time, encoder zapcore.PrimitiveArrayEncoder) {}
	}
	zapCfg.OutputPaths = []string{fmt.Sprintf("rotate://%s", opts.Path)}
	if opts.Stdout {
		zapCfg.OutputPaths = append(zapCfg.OutputPaths, "stdout")
	}
	logger, err := zapCfg.Build()
	if err != nil {
		return nil, errors.Wrap(err, "")
	}
	return logger, nil
}
