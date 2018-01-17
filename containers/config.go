package containers

import (
	"flag"
	"fmt"
	"os"
	"strings"
	"github.com/spf13/pflag"
	"github.com/spf13/viper"
	"log"
)

var Cfg = viper.New()

func initConfig() (err error) {
	Cfg.SetDefault("agent-code.version", "1.0")
	pflag.CommandLine.AddGoFlagSet(flag.CommandLine)
	//cli&default config
	configFile := pflag.String("config", "", "config file path")
	version := pflag.Bool("version", false, "show version")
	if *version {
		fmt.Printf("agent-code v%s - https://github.com/bzppx/bzppx-codepub\n", Cfg.GetString("agent-code.version"))
		os.Exit(0)
	}
	pflag.String("rpc-listen", ":9091", "the api port to listen")
	pflag.String("level", "debug", "log level to show in console")
	pflag.String("log-dir", "log", "the directory which store log files")
	pflag.Int64("log-max-size", 102400000, "log file max size(bytes) for rotate")
	pflag.Int("log-max-count", 3, "log file max count for rotate to remain")
	pflag.StringSlice("log-level", []string{"info", "error", "debug"}, "log to file level,multiple splitted by comma(,)")
	pflag.Parse()

	//bind flag
	Cfg.BindPFlag("rpc.listen", pflag.Lookup("rpc-listen"))
	Cfg.BindPFlag("log.dir", pflag.Lookup("log-dir"))
	Cfg.BindPFlag("log.level", pflag.Lookup("log-level"))
	Cfg.BindPFlag("log.console-level", pflag.Lookup("level"))
	Cfg.BindPFlag("log.fileMaxSize", pflag.Lookup("log-max-size"))
	Cfg.BindPFlag("log.maxCount", pflag.Lookup("log-max-count"))
	if *configFile != "" {
		Cfg.SetConfigFile(*configFile)
	} else {
		Cfg.SetConfigName("config")
		Cfg.AddConfigPath("/etc/agent-code/")
		Cfg.AddConfigPath("$HOME/.agent-code")
		Cfg.AddConfigPath(".agent-code")
		Cfg.AddConfigPath(".")
	}
	err = Cfg.ReadInConfig()
	file := Cfg.ConfigFileUsed()
	if err != nil && !strings.Contains(fmt.Sprintf("%s", err.Error()), "Not") {
		log.Printf("%s", err)
	} else if file != "" {
		log.Printf("use config file : %s\n", file)
	}
	return
}