package containers

import (
	"bzppx-agent-codepub/utils"
	"github.com/snail007/mini-logger"
	"github.com/snail007/mini-logger/writers/console"
	"github.com/snail007/mini-logger/writers/files"
)

var Log logger.MiniLogger
var AccessLog logger.MiniLogger

//initLog
func initLog() {
	var level uint8
	switch Cfg.GetString("log.console-level") {
	case "debug":
		level = logger.AllLevels
	case "info":
		level = logger.InfoLevel | logger.WarnLevel | logger.ErrorLevel | logger.FatalLevel
	case "warn":
		level = logger.WarnLevel | logger.ErrorLevel | logger.FatalLevel
	case "error":
		level = logger.ErrorLevel | logger.FatalLevel
	case "fatal":
		level = logger.FatalLevel
	default:
		level = 0
	}
	Log = logger.New(false, nil)
	Log.AddWriter(console.New(console.ConsoleWriterConfig{
		Format: "{date} {time}.{mili} {level} {fields} {text}",
		Type:   console.T_TEXT,
	}), level)
	CfgF := files.GetDefaultFileConfig()
	CfgF.LogPath = Cfg.GetString("log.dir")
	CfgF.MaxBytes = Cfg.GetInt64("log.FileMaxSize")
	CfgF.MaxCount = Cfg.GetInt("log.MaxCount")
	CfgLevels := Cfg.GetStringSlice("log.level")
	if ok, _ := utils.InArray("debug", CfgLevels); ok {
		CfgF.FileNameSet["debug"] = logger.AllLevels
	}
	if ok, _ := utils.InArray("info", CfgLevels); ok {
		CfgF.FileNameSet["info"] = logger.InfoLevel
	}
	if ok, _ := utils.InArray("error", CfgLevels); ok {
		CfgF.FileNameSet["error"] = logger.WarnLevel | logger.ErrorLevel | logger.FatalLevel
	}
	Log.AddWriter(files.New(CfgF), logger.AllLevels)
}