package log

import (
	stdLog "log"
	"os"
	"path"
	"time"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	log "wx-gitlab.xunlei.cn/galaxy/go-logger"
)

var (
	logLevelMap = map[string]zapcore.Level{
		"debug": zap.DebugLevel,
		"info":  zap.InfoLevel,
		"warn":  zap.WarnLevel,
		"error": zap.ErrorLevel,
	}
	logging *Logger
)

type Options struct {
	path       string
	name       string
	maxAge     int
	maxSize    int
	maxBackups int
	compress   bool
	level      string
}

type Logger struct {
	opts *Options
	*zap.Logger
	ws log.Logger
}

type Option func(o *Options)

func MaxSize(v int) Option {
	return func(o *Options) {
		o.maxSize = v
	}
}

func MaxAge(v int) Option {
	return func(o *Options) {
		o.maxAge = v
	}
}

func MaxBackups(v int) Option {
	return func(o *Options) {
		o.maxBackups = v
	}
}

func Path(v string) Option {
	return func(o *Options) {
		o.path = v
	}
}

func Name(v string) Option {
	return func(o *Options) {
		o.name = v
	}
}

func Compress(v bool) Option {
	return func(o *Options) {
		o.compress = v
	}
}

func Level(v string) Option {
	return func(o *Options) {
		o.level = v
	}
}

func NewLogger(opts ...Option) *Logger {
	options := &Options{
		maxAge:     7,
		maxSize:    100,
		maxBackups: 10,
		compress:   false,
		path:       "./logs",
		name:       "trace.log",
	}

	for _, o := range opts {
		o(options)
	}

	if err := os.MkdirAll(options.path, 0766); err != nil {
		stdLog.Println("failed create log directory:", options.path, ":", err)
		return nil
	}

	encCfg := zap.NewProductionEncoderConfig()
	encCfg.EncodeTime = func(t time.Time, enc zapcore.PrimitiveArrayEncoder) {
		enc.AppendString(t.Format("2006-01-02 15:04:05.000"))
	}
	encCfg.EncodeDuration = zapcore.SecondsDurationEncoder
	encCfg.EncodeCaller = zapcore.FullCallerEncoder

	dl := zap.NewAtomicLevel()
	if l, ok := logLevelMap[options.level]; ok {
		dl.SetLevel(l)
	} else {
		dl.SetLevel(zap.InfoLevel)
	}

	encoder := zapcore.NewConsoleEncoder(encCfg)
	ws := log.NewWriteSyncer(
		log.MaxAge(options.maxAge),
		log.MaxSize(options.maxSize),
		log.MaxBackups(options.maxBackups),
		log.Compress(options.compress),
		log.Filename(path.Join(options.path, options.name)),
	)

	l := zap.New(
		zapcore.NewCore(
			encoder,
			zapcore.NewMultiWriteSyncer(ws.(zapcore.WriteSyncer)),
			dl,
		),
		zap.AddCaller(),
	)
	zap.RedirectStdLog(l)
	return &Logger{
		opts:   options,
		Logger: l,
	}
}

func Init(opts ...Option) {
	logging = NewLogger(opts...)
}

func L() *Logger {
	return logging
}

func FlushLogs() error {
	if logging != nil {
		return logging.Sync()
	}
	return nil
}

func With(fields ...zapcore.Field) *zap.Logger {
	return logging.With(fields...)
}
