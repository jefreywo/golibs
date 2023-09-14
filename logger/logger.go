package logger

import (
	"io"
	"time"

	rotatelogs "github.com/lestrrat-go/file-rotatelogs"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"gopkg.in/natefinch/lumberjack.v2"
)

var publicLogger *zap.Logger

type LogLevel string

const (
	DebugLevel  LogLevel = "debug"
	InfoLevel   LogLevel = "info"
	WarnLevel   LogLevel = "warn"
	ErrorLevel  LogLevel = "error"
	DPanicLevel LogLevel = "dpanic"
	PanicLevel  LogLevel = "panic"
	FatalLevel  LogLevel = "fatal"
)

var LogLevelMatchMap = map[LogLevel]zapcore.Level{
	DebugLevel:  zapcore.DebugLevel,
	InfoLevel:   zapcore.InfoLevel,
	WarnLevel:   zapcore.WarnLevel,
	ErrorLevel:  zapcore.ErrorLevel,
	DPanicLevel: zapcore.DPanicLevel,
	PanicLevel:  zapcore.PanicLevel,
	FatalLevel:  zapcore.FatalLevel,
}

var DefaultEncoderConfig = zapcore.EncoderConfig{
	LevelKey:   "level",
	TimeKey:    "time",
	CallerKey:  "",
	MessageKey: "content",
}

type LoggerOption struct {
	Level           LogLevel
	Encoder         zapcore.EncoderConfig
	DivisionWriter  io.Writer
	ErrWriterSyncer zapcore.WriteSyncer
	ExtendCore      []zapcore.Core
}

type Option func(*LoggerOption)

func InitLoger(fileName string, opts ...Option) {
	opt := &LoggerOption{
		Level:          InfoLevel,
		DivisionWriter: sizeDivisionWriter(fileName),
	}

	for _, o := range opts {
		o(opt)
	}

	corelevel, ok := LogLevelMatchMap[opt.Level]
	if !ok {
		corelevel = zapcore.InfoLevel
	}

	var core []zapcore.Core
	core = append(core, zapcore.NewCore(
		opt.Encoder,
		zapcore.NewMultiWriteSyncer(zapcore.AddSync(opt.DivisionWriter)),
		zap.NewAtomicLevelAt(corelevel),
	))

	if opt.ErrWriterSyncer != nil {
		errLevel := zap.NewAtomicLevel()
		errLevel.SetLevel(zapcore.ErrorLevel)
		core = append(core, zapcore.NewCore(
			opt.Encoder,
			zapcore.NewMultiWriteSyncer(opt.ErrWriterSyncer),
			errLevel,
		))
	}

	if len(opt.ExtendCore) > 0 {
		core = append(core, opt.ExtendCore...)
	}

	publicLogger = zap.New(
		zapcore.NewTee(core...),
		zap.AddCaller(),
		zap.AddCallerSkip(1),
	)
}

func SetLogLevel(level LogLevel) Option {
	return func(lo *LoggerOption) {
		lo.Level = level
	}
}
func SetDivisionWriter(hook io.Writer) Option {
	return func(lo *LoggerOption) {
		lo.DivisionWriter = hook
	}
}

// 按大小切分日志文件
func sizeDivisionWriter(fileName string) io.Writer {
	return &lumberjack.Logger{
		Filename:   fileName,
		MaxSize:    128,
		MaxBackups: 16,
		MaxAge:     30,
		Compress:   true,
		LocalTime:  true,
	}
}

// 按时间切分日志文件
func timeDivisionWriter(filename string) io.Writer {
	hook, err := rotatelogs.New(
		filename+".%Y%m%d",                        // %Y%m%d%H 年月日时
		rotatelogs.WithLinkName(filename),         // 建立软链接
		rotatelogs.WithRotationTime(time.Hour*24), // 按天
		rotatelogs.WithRotationCount(30),          // 保留30个, 与WithMaxAge不能同时设定，
		//rotatelogs.WithMaxAge(time.Hour*24),	   // 不设定默认保存7天
	)
	if err != nil {
		panic(err)
	}
	return hook
}
