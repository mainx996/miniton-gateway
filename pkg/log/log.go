package log

import (
	"context"
	"os"
	"sync"
	"www.miniton-gateway.com/pkg/config"

	"github.com/gin-gonic/gin"
	"github.com/natefinch/lumberjack"
	uuid "github.com/segmentio/ksuid"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

type (
	Conf struct {
		FileName  string
		MaxSize   int  // 单个文件大小, 单位兆
		MaxBackup int  // 保留旧文件的最大个数
		MaxAge    int  // 保留旧文件的最大天数
		Compress  bool // 是否压缩/归档旧文件
		LocalTime bool // 日志备份文件名称
		Level     int  // 日志级别
	}
	// Logger wrap zap.logger
	Logger struct {
		sync.Mutex
		l *zap.Logger
	}

	// Field wrap
	Field = zap.Field

	// Option wrap
	Option = zap.Option

	// ctx key
	traceLogKey struct{}
	traceIDKey  struct{}
)

var (
	// Log instance
	Log *Logger

	// TraceIDLogField TraceIDLogField
	TraceIDLogField = "traceid"
)

func Init() {
	logConfig := config.Config.LogConfig
	ec := zap.NewProductionEncoderConfig()
	ec.EncodeTime = zapcore.ISO8601TimeEncoder
	ec.EncodeLevel = zapcore.CapitalLevelEncoder
	core := zapcore.NewCore(
		zapcore.NewJSONEncoder(ec),
		zapcore.NewMultiWriteSyncer(zapcore.AddSync(os.Stdout), zapcore.AddSync(&lumberjack.Logger{
			Filename:   logConfig.FileName,
			MaxSize:    logConfig.MaxSize,
			MaxBackups: logConfig.MaxBackup,
			MaxAge:     logConfig.MaxAge,
			Compress:   true,
			LocalTime:  false,
		})),
		zapcore.Level(logConfig.Level),
	)
	Log = New(zap.New(core, zap.AddCaller(), zap.AddCallerSkip(2)))
}

// New create logger instance
func New(log *zap.Logger) *Logger {
	return &Logger{
		l: log,
	}
}

// With wrap
func (l *Logger) With(fields ...Field) *Logger {
	l.Lock()
	defer l.Unlock()

	return &Logger{
		l: l.l.With(fields...),
	}
}

// WithOptions wrap
func (l *Logger) WithOptions(opts ...Option) *Logger {
	l.Lock()
	defer l.Unlock()

	return &Logger{
		l: l.l.WithOptions(opts...),
	}
}

// Debug wrap
func (l *Logger) Debug(msg string, fields ...Field) {
	l.l.Debug(msg, fields...)
}

// Info wrap
func (l *Logger) Info(msg string, fields ...Field) {
	l.l.Info(msg, fields...)
}

// Warn wrap
func (l *Logger) Warn(msg string, fields ...Field) {
	l.l.Warn(msg, fields...)
}

// Error wrap
func (l *Logger) Error(msg string, fields ...Field) {
	l.l.Error(msg, fields...)
}

// TraceID get traceID
func TraceID(ctx context.Context) string {
	v := ctx.Value(traceIDKey{})
	if v != nil {
		if s, ok := v.(string); ok {
			return s
		}
	}
	return uuid.New().String()
}

// NewTraceLog NewTraceLog
func NewTraceLog(ctx context.Context, log *Logger, args ...string) context.Context {
	var traceID string
	if len(args) > 0 {
		traceID = args[0]
	}

	if traceID == "" {
		traceID = TraceID(ctx)
	}

	l := log.With(zap.String(TraceIDLogField, traceID))
	ctx = context.WithValue(ctx, traceIDKey{}, traceID)
	return context.WithValue(ctx, traceLogKey{}, l)
}

// Info Info
func Info(ctx context.Context, msg string, fields ...Field) {
	FromTrace(ctx).Info(msg, fields...)
}

// Warn Warn
func Warn(ctx context.Context, msg string, fields ...Field) {
	FromTrace(ctx).Warn(msg, fields...)
}

// Error Error
func Error(ctx context.Context, msg string, fields ...Field) {
	FromTrace(ctx).Error(msg, fields...)
}

// FromTrace get logger from context
func FromTrace(ctx context.Context) *Logger {
	if c, ok := ctx.(*gin.Context); ok {
		ctx = c.Request.Context()
	}

	v := ctx.Value(traceLogKey{})
	if v != nil {
		if l, ok := v.(*Logger); ok {
			return l
		}
	}
	return Log
}
