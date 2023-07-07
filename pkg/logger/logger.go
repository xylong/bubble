package logger

import (
	"bubble/pkg/app"
	"fmt"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"gopkg.in/natefinch/lumberjack.v2"
	"os"
	"strings"
	"time"
)

// Logger 全局 Logger 对象
var Logger *zap.Logger

func InitLogger(filename string, maxSize, maxBackup, maxAge int, compress bool, logType string, level string) {

	// 获取日志写入介质
	writer := getLogWriter(filename, logType, maxSize, maxBackup, maxAge, compress)

	// 设置日志等级，具体请见 config/log.go 文件
	logLevel := getLevel(level)

	// 获取日志存储格式
	encoder := getEncoder()

	// 初始化core
	core := zapcore.NewCore(encoder, writer, logLevel)
	// 初始化 Logger
	Logger = zap.New(core,
		zap.AddCaller(),                   // 调用文件和行号，内部使用 runtime.Caller
		zap.AddCallerSkip(1),              // 封装了一层，调用文件去除一层(runtime.Caller(1))
		zap.AddStacktrace(zap.ErrorLevel), // Error 时才会显示 stacktrace
	)

	// 将自定义的 logger 替换为全局的 logger
	// zap.L().Fatal() 调用时，就会使用我们自定的 Logger
	zap.ReplaceGlobals(Logger)
}

func getLogWriter(filename, logType string, maxSize, maxBackup, maxAge int, compress bool) zapcore.WriteSyncer {

	// 如果配置了按照日期记录日志文件
	if logType == "daily" {
		name := time.Now().Format("2006-01-02") + ".log"
		filename = strings.ReplaceAll(filename, "logs.log", name)
	}

	// 滚动日志，详见 config/log.go
	lumberJackLogger := &lumberjack.Logger{
		Filename:   filename,
		MaxSize:    maxSize,
		MaxAge:     maxAge,
		MaxBackups: maxBackup,
		Compress:   compress,
	}

	// 配置输出介质
	if app.IsLocal() {
		// 本地开发终端打印和记录文件
		return zapcore.NewMultiWriteSyncer(zapcore.AddSync(os.Stdout), zapcore.AddSync(lumberJackLogger))
	} else {
		// 生产环境只记录日志
		return zapcore.AddSync(lumberJackLogger)
	}
}

func getLevel(level string) *zapcore.Level {
	logLevel := new(zapcore.Level)
	if err := logLevel.UnmarshalText([]byte(level)); err != nil {
		fmt.Println("日志初始化错误，日志级别设置有误。请修改 config/log.go 文件中的 log.level 配置项")
	}

	return logLevel
}

func getEncoder() zapcore.Encoder {

	// 日志格式规则
	ec := zapcore.EncoderConfig{
		TimeKey:       "time",
		LevelKey:      "level",
		NameKey:       "logger",
		CallerKey:     "caller", // 代码调用，如 paginator/paginator.go:148
		FunctionKey:   zapcore.OmitKey,
		MessageKey:    "message",
		StacktraceKey: "stacktrace",
		LineEnding:    zapcore.DefaultLineEnding,   // 每行日志的结尾添加 "\n"
		EncodeLevel:   zapcore.CapitalLevelEncoder, // 日志级别名称大写，如 ERROR、INFO
		EncodeTime: func(time time.Time, encoder zapcore.PrimitiveArrayEncoder) {
			encoder.AppendString(time.Format("2006-01-02 15:04:05.000"))
		}, // 时间格式，我们自定义为 2006-01-02 15:04:05
		EncodeDuration: zapcore.SecondsDurationEncoder, // 执行时间，以秒为单位
		EncodeCaller:   zapcore.ShortCallerEncoder,     // Caller 短格式，如：types/converter.go:17，长格式为绝对路径
	}

	if app.IsLocal() {
		// 终端输出的关键词高亮
		ec.EncodeLevel = zapcore.CapitalColorLevelEncoder
		// 控制台输出（支持 stacktrace 换行）
		return zapcore.NewConsoleEncoder(ec)
	}

	// 线上环境使用 JSON 编码器
	return zapcore.NewJSONEncoder(ec)
}
