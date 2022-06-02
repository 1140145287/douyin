package logger

import (
	"douyin/pkg/setting"

	"github.com/natefinch/lumberjack"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

/***
自定义logger主要有一下三个部分：
1、确定要写入的位置
2、确定编码方式
3、确定等级（弹出提示）
***/
func NewLogger(loggerSetting *setting.LoggerSettingS, runMode string) (*zap.Logger, error) {
	//确定写入位置
	//os.OpenFile(loggerSetting.LogSavePath+"/"+loggerSetting.LogFileName, os.O_CREATE|os.O_APPEND|os.O_RDWR, 0744)
	lumberjackLogger := &lumberjack.Logger{
		Filename:   loggerSetting.LogSavePath + "/" + loggerSetting.LogFileName,
		MaxSize:    loggerSetting.MaxPageSize,
		MaxBackups: loggerSetting.LogMaxBackups,
		MaxAge:     loggerSetting.MaxPageAge,
	}
	writeSyncer := zapcore.AddSync(lumberjackLogger)
	encoder := getEncoder()
	logLevel := new(zapcore.Level)
	err := logLevel.UnmarshalText([]byte(loggerSetting.LogLevel))
	var core zapcore.Core
	core = zapcore.NewCore(encoder, writeSyncer, logLevel)
	logger := zap.New(core)

	return logger, err
}

//获取日志对应的编码方式
func getEncoder() zapcore.Encoder {
	encoderConfig := zap.NewProductionEncoderConfig()
	encoderConfig.EncodeTime = zapcore.ISO8601TimeEncoder
	encoderConfig.EncodeLevel = zapcore.CapitalLevelEncoder
	return zapcore.NewConsoleEncoder(encoderConfig)
}
