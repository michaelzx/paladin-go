package zaplog

import (
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

var (
	zapLogger *zap.Logger
	logDir    = "./app.log"
)

func GetZap() *zap.Logger {
	return zapLogger
}

// // Init 初始化，指定是否是生产环境
// func Init(cfg *config.LoggerConfig) {
// 	zapCfg := getZapConfig(cfg.Dev, cfg.FileLog)
// 	// 因为我们做了一层包装，所以需要跳过一层caller
// 	// 否则，日志的caller位置，始终显示的是当前logger包中的位置
// 	callerOption := zap.AddCallerSkip(1)
// 	var err error
// 	zapLogger, err = zapCfg.Build(callerOption)
// 	if err != nil {
// 		panic(err)
// 	}
// 	defer func() {
// 		zapLogger.Sync()
// 	}()
// }

func getZapConfig(devMode bool, logFilePath string) zap.Config {
	var loggingLevel zapcore.Level
	var OutputPaths []string
	var ErrorOutputPaths []string
	if devMode {
		// 开发模式
		loggingLevel = zap.DebugLevel
		OutputPaths = []string{"stdout"}
		ErrorOutputPaths = []string{"stderr"}
	} else {
		// 生产模式
		// TODO 对日志进行分割
		loggingLevel = zap.InfoLevel
		if logFilePath != "" {
			OutputPaths = []string{"stdout"}
			ErrorOutputPaths = []string{"stderr"}
		} else {
			OutputPaths = []string{"stdout"}
			ErrorOutputPaths = []string{"stderr"}
		}
	}
	return zap.Config{
		Level:       zap.NewAtomicLevelAt(loggingLevel),
		Development: devMode,
		Encoding:    "console",
		// 如果需要忽略输出，则赋值：zapcore.OmitKey
		EncoderConfig: zapcore.EncoderConfig{
			TimeKey:        "ts",
			LevelKey:       "level",
			NameKey:        "logger",
			CallerKey:      "callerKey",
			MessageKey:     "msg",
			StacktraceKey:  "stacktrace",
			LineEnding:     zapcore.DefaultLineEnding,
			EncodeLevel:    zapcore.CapitalColorLevelEncoder,
			EncodeTime:     zapcore.ISO8601TimeEncoder,
			EncodeDuration: zapcore.NanosDurationEncoder,
			EncodeCaller:   customCallerEncoder,
		},
		OutputPaths:      OutputPaths,
		ErrorOutputPaths: ErrorOutputPaths,
	}
}
func customCallerEncoder(caller zapcore.EntryCaller, enc zapcore.PrimitiveArrayEncoder) {
	// TODO 考虑优化caller显示方式
	enc.AppendString(caller.TrimmedPath())
}
