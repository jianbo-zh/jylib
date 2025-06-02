package bootstrap

import (
	"fmt"
	"os"

	configV1 "github.com/jianbo-zh/jypb/config/v1"

	"github.com/go-kratos/kratos/v2/log"
	"github.com/natefinch/lumberjack"
)

func NewLoggerProvider(serviceInfo *ServiceInfo, bc *configV1.Log) log.Logger {

	saveDir := "./"
	if bc == nil {
		panic(fmt.Errorf("log config is nil"))

	} else if bc.SaveDir != "" {
		saveDir = bc.SaveDir
		fi, err := os.Stat(saveDir)
		if err != nil {
			if os.IsNotExist(err) {
				panic(fmt.Errorf("log save dir %s not exist", saveDir))

			} else {
				panic(fmt.Errorf("os.State %s error: %v", saveDir, err))
			}
		} else if !fi.IsDir() {
			panic(fmt.Errorf("log save dir %s is not a dir", saveDir))
		}
	}

	ll := &lumberjack.Logger{
		Filename:   fmt.Sprintf("%s/%s.log", saveDir, serviceInfo.GetInstanceId()),
		MaxSize:    int(bc.MaxSizeMb),
		MaxAge:     int(bc.MaxAgeDay),
		MaxBackups: int(bc.MaxBackups),
		LocalTime:  bc.LocalTime,
		Compress:   bc.Compress,
	}

	level := log.LevelInfo
	switch bc.Level {
	case configV1.Log_DEBUG:
		level = log.LevelDebug
	case configV1.Log_INFO:
		level = log.LevelInfo
	case configV1.Log_WARN:
		level = log.LevelWarn
	case configV1.Log_ERROR:
		level = log.LevelError
	case configV1.Log_FATAL:
		level = log.LevelFatal
	}

	l := log.NewFilter(log.NewStdLogger(ll), log.FilterLevel(level))

	return log.With(
		l,
		// "service.id", serviceInfo.Id,
		// "service.name", serviceInfo.Name,
		// "service.version", serviceInfo.Version,
		"ts", log.DefaultTimestamp,
		"caller", log.DefaultCaller,
		// "trace_id", tracing.TraceID(),
		// "span_id", tracing.SpanID(),
	)
}
