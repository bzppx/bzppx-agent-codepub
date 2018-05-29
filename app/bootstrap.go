package app

import (
	"flag"
	"fmt"
	"os"
	"strings"

	"github.com/fatih/color"
	"github.com/phachon/go-logger"
	"github.com/spf13/viper"
)

var (
	flagConf = flag.String("conf", "config.toml", "please input conf path")
)

var (
	Version = "v0.8.4"

	Conf = viper.New()

	Log = go_logger.NewLogger()

	AppPath = ""

	RootPath = ""
)

// 启动初始化
func init() {
	initFlag()
	initConfig()
	initLog()
	initPoster()
	initPath()
}

// init flag
func initFlag() {
	flag.Parse()
}

// init config
func initConfig() {

	if *flagConf == "" {
		Log.Error("config file not found!")
		os.Exit(1)
	}

	Conf.SetConfigType("toml")
	Conf.SetConfigFile(*flagConf)
	err := Conf.ReadInConfig()
	if err != nil {
		Log.Error("Fatal error config file: " + err.Error())
		os.Exit(1)
	}

	file := Conf.ConfigFileUsed()
	if file != "" {
		Log.Info("Use config file: " + file)
	}
}

// init log
func initLog() {

	// console adapter
	Log.Detach("console")
	consoleConfig := &go_logger.ConsoleConfig{
		Color: true,
	}
	Log.Attach("console", go_logger.LOGGER_LEVEL_DEBUG, consoleConfig)

	// file adapter config
	fileLevelStr := Conf.GetString("log.level")
	levelFilenameConf := Conf.GetStringMapString("log.levelFilename")
	levelFilename := map[int]string{}
	if len(levelFilenameConf) > 0 {
		for levelStr, levelFile := range levelFilenameConf {
			level := Log.LoggerLevel(levelStr)
			levelFilename[level] = levelFile
		}
	}
	fileConfig := &go_logger.FileConfig{
		Filename:      Conf.GetString("log.filename"),
		LevelFileName: levelFilename,
		MaxSize:       Conf.GetInt64("log.maxSize"),
		MaxLine:       Conf.GetInt64("log.maxLine"),
		DateSlice:     Conf.GetString("log.dateSlice"),
		JsonFormat:    Conf.GetBool("log.jsonFormat"),
		Format:        Conf.GetString("log.format"),
	}
	Log.Attach("file", Log.LoggerLevel(fileLevelStr), fileConfig)
}

// init poster
func initPoster() {
	fg := color.New(color.FgHiYellow).SprintFunc()
	logo := fg(`
                    __                     __                                     __
  _____ ____   ____/ /___   ____   __  __ / /_       ____    ____   ___   ____   / /_
 / ___// __ \ / __  // _ \ / __ \ / / / // __ \     / __ \  / __ \ / _ \ / __ \ / __/
/ /__ / /_/ // /_/ //  __// /_/ // /_/ // /_/ / -- / /_/ / / /_/ //  __// / / // /_
\___/ \____/ \__,_/ \___ / .___/ \__,_//_.___/     \__/\_\ \___,/ \___//_/ /_/ \__/
                        /_/                               /____/
` +
		"Author: bzppx\r\n" +
		"Version: " + Version + "\r\n" +
		"Link: https://github.com/bzppx/bzppx-agent-codepub")
	fmt.Println(logo)
}

// init dir and path
func initPath() {
	AppPath, _ = os.Getwd()
	RootPath = strings.Replace(AppPath, "app", "", 1)
}
