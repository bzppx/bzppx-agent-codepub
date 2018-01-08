package main

import (
	"bzppx-agent-codepub/utils"
	"github.com/snail007/mini-logger"
	"github.com/snail007/mini-logger/writers/console"
	"github.com/snail007/mini-logger/writers/files"
)

var log logger.MiniLogger
var accessLog logger.MiniLogger

//initLog
func initLog() {
	var level uint8
	switch cfg.GetString("log.console-level") {
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
	log = logger.New(false, nil)
	log.AddWriter(console.New(console.ConsoleWriterConfig{
		Format: "{date} {time}.{mili} {level} {fields} {text}",
		Type:   console.T_TEXT,
	}), level)
	cfgF := files.GetDefaultFileConfig()
	cfgF.LogPath = cfg.GetString("log.dir")
	cfgF.MaxBytes = cfg.GetInt64("log.FileMaxSize")
	cfgF.MaxCount = cfg.GetInt("log.MaxCount")
	cfgLevels := cfg.GetStringSlice("log.level")
	if ok, _ := utils.InArray("debug", cfgLevels); ok {
		cfgF.FileNameSet["debug"] = logger.AllLevels
	}
	if ok, _ := utils.InArray("info", cfgLevels); ok {
		cfgF.FileNameSet["info"] = logger.InfoLevel
	}
	if ok, _ := utils.InArray("error", cfgLevels); ok {
		cfgF.FileNameSet["error"] = logger.WarnLevel | logger.ErrorLevel | logger.FatalLevel
	}
	log.AddWriter(files.New(cfgF), logger.AllLevels)
}
