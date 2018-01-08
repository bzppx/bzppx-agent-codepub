package main

import (
	"flag"
	"fmt"
	"os"
	"strings"
	"github.com/fatih/color"
	"github.com/spf13/pflag"
	"github.com/spf13/viper"
)

var (
	cfg = viper.New()
)

func initConfig() (err error) {
	cfg.SetDefault("agent-code.version", "1.0")
	pflag.CommandLine.AddGoFlagSet(flag.CommandLine)
	//cli&default config
	configFile := pflag.String("config", "", "config file path")
	version := pflag.Bool("version", false, "show version")
	if *version {
		fmt.Printf("agent-code v%s - https://github.com/bzppx/bzppx-codepub\n", cfg.GetString("agent-code.version"))
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
	cfg.BindPFlag("rpc.listen", pflag.Lookup("rpc-listen"))
	cfg.BindPFlag("log.dir", pflag.Lookup("log-dir"))
	cfg.BindPFlag("log.level", pflag.Lookup("log-level"))
	cfg.BindPFlag("log.console-level", pflag.Lookup("level"))
	cfg.BindPFlag("log.fileMaxSize", pflag.Lookup("log-max-size"))
	cfg.BindPFlag("log.maxCount", pflag.Lookup("log-max-count"))
	if *configFile != "" {
		cfg.SetConfigFile(*configFile)
	} else {
		cfg.SetConfigName("config")
		cfg.AddConfigPath("/etc/agent-code/")
		cfg.AddConfigPath("$HOME/.agent-code")
		cfg.AddConfigPath(".agent-code")
		cfg.AddConfigPath(".")
	}
	err = cfg.ReadInConfig()
	file := cfg.ConfigFileUsed()
	if err != nil && !strings.Contains(fmt.Sprintf("%s", err.Error()), "Not") {
		fmt.Printf("%s", err)
	} else if file != "" {
		fmt.Printf("use config file : %s\n", file)
	}
	return
}

func poster() string {
	fg := color.New(color.FgHiYellow).SprintFunc()
	return fg(`
 █████╗      ██████╗     ███████╗    ███╗   ██╗    ████████╗    ██╗  ██╗
██╔══██╗    ██╔════╝     ██╔════╝    ████╗  ██║    ╚══██╔══╝    ╚██╗██╔╝
███████║    ██║  ███╗    █████╗      ██╔██╗ ██║       ██║        ╚███╔╝ 
██╔══██║    ██║   ██║    ██╔══╝      ██║╚██╗██║       ██║        ██╔██╗ 
██║  ██║    ╚██████╔╝    ███████╗    ██║ ╚████║       ██║       ██╔╝ ██╗
╚═╝  ╚═╝     ╚═════╝     ╚══════╝    ╚═╝  ╚═══╝       ╚═╝       ╚═╝  ╚═╝
Author: bzppx
Link  : https://github.com/bzppx/bzppx-codepub
`)
}
