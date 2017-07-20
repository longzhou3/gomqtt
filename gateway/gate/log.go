package gate

import (
	"encoding/json"
	"fmt"
	"log"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

// Logger ...
var Logger *zap.Logger

// InitLogger ...
func InitLogger(lp string, lv string, isDebug bool) {

	var js string
	if isDebug {
		js = fmt.Sprintf(`{
		"level": "%s",
		"encoding": "json",
		"outputPaths": ["stdout"],
		"errorOutputPaths": ["stdout"]
		}`, lv)
	} else {
		js = fmt.Sprintf(`{
		"level": "%s",
		"encoding": "json",
		"outputPaths": ["%s"],
		"errorOutputPaths": ["%s"]
		}`, lv, lp, lp)
	}

	var cfg zap.Config
	if err := json.Unmarshal([]byte(js), &cfg); err != nil {
		panic(err)
	}

	cfg.EncoderConfig = zap.NewProductionEncoderConfig()
	cfg.EncoderConfig.EncodeTime = zapcore.ISO8601TimeEncoder
	// cfg.EncoderConfig.EncodeCaller = zapcore.FullCallerEncoder
	var err error
	Logger, err = cfg.Build()
	if err != nil {
		log.Fatal("init logger error: ", err)
	}

	// var level zap.Level

	// switch strings.ToLower(lv) {
	// case "debug":
	// 	level = zap.DebugLevel
	// case "info":
	// 	level = zap.InfoLevel
	// case "warn":
	// 	level = zap.WarnLevel
	// case "error":
	// 	level = zap.ErrorLevel
	// case "fatal":
	// 	level = zap.FatalLevel
	// default:
	// 	level = zap.DebugLevel
	// }

	// if isDebug {
	// 	Logger = zap.New(
	// 		zap.NewJSONEncoder(
	// 			zap.RFC3339NanoFormatter("@timestamp"), // human-readable timestamps
	// 			zap.MessageKey("@message"),             // customize the message key
	// 			zap.LevelString("@level"),              // stringify the log level
	// 		),
	// 		zap.AddCaller(),
	// 		level,
	// 	)
	// } else {
	// 	f, err := os.OpenFile(lp, os.O_APPEND|os.O_RDWR|os.O_CREATE, 0666)
	// 	if err != nil {
	// 		log.Panic(err)
	// 	}

	// 	Logger = zap.New(
	// 		zap.NewJSONEncoder(
	// 			zap.RFC3339Formatter("@timestamp"), // human-readable timestamps
	// 			zap.MessageKey("@message"),         // customize the message key
	// 			zap.LevelString("@level"),          // stringify the log level
	// 		),
	// 		zap.Output(f),
	// 		zap.AddCaller(),
	// 		level,
	// 	)
	// }
}
