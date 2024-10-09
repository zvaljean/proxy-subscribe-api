package log

import (
	"fmt"
	"os"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"gopkg.in/natefinch/lumberjack.v2"
)

const LogPre = "#----#"

func InitLog(logName string) {
	writeSyncer := zapcore.AddSync(
		&lumberjack.Logger{
			Filename: logName,
			MaxSize:  512,
			// MaxAge:     0,
			MaxBackups: 0,
			LocalTime:  true,
			Compress:   false,
		})

	encoder := zap.NewProductionEncoderConfig()
	encoder.EncodeTime = zapcore.ISO8601TimeEncoder

	// zapcore.Lock(os.Stdout)
	// https://stackoverflow.com/questions/68472667/how-to-log-to-stdout-or-stderr-based-on-log-level-using-uber-go-zap
	core := zapcore.NewTee(
		zapcore.NewCore(
			zapcore.NewJSONEncoder(encoder),
			writeSyncer,
			zap.NewAtomicLevelAt(zap.DebugLevel)),

		zapcore.NewCore(
			zapcore.NewConsoleEncoder(encoder),
			os.Stdout,
			zap.NewAtomicLevelAt(zap.DebugLevel)))

	log := zap.New(core,
		zap.AddCaller(),
		zap.AddCallerSkip(1))

	zap.ReplaceGlobals(log)
}

func tip(temp string) string {
	return fmt.Sprintf("%s-> %s ", LogPre, temp)
}

// FatalCheck, fail will call os.Exit
func FatalCheck(err error, msg string) {
	if err != nil {
		Fatal(msg, err.Error())
	}
}

func ErrorCheck(err error, msg string) bool {
	if err != nil {
		Error(msg, err.Error())
		return true
	}
	return false
}

func ZapL() *zap.Logger {
	return zap.L()
}
func Zap() *zap.SugaredLogger {
	return zap.S()
}

// Info : tip + newline
func Info(format string, args ...interface{}) {
	zap.S().Infof(tip(format), args...)
}

// Error : tip + newline
func Error(format string, args ...interface{}) {
	zap.S().Errorf(tip(format), args...)
}

// Warn : tip + newline
func Warn(format string, args ...interface{}) {
	zap.S().Warnf(tip(format), args...)
}

// Debug : tip + newline
func Debug(format string, args ...interface{}) {
	zap.S().Debugf(tip(format), args...)
}

// Fatal : tip + newline
func Fatal(format string, args ...interface{}) {
	zap.S().Fatalf(tip(format), args...)
}
