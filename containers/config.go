package containers

import (
	"flag"
	"github.com/spf13/viper"
	"log"
	"os"
)

var Cfg = viper.New()

var (
	flagConf = flag.String("conf", "config.toml", "please input conf file path")
)

// init flag
func initFlag() {
	flag.Parse()
}

// init config
func initConfig() (err error) {
	if *flagConf == "" {
		log.Println("config file not found!")
		os.Exit(1)
	}

	Cfg.SetConfigType("toml")
	Cfg.SetConfigFile(*flagConf)
	err = Cfg.ReadInConfig()
	if err != nil {
		log.Printf("Fatal error config file: "+err.Error())
		os.Exit(1)
	}

	file := Cfg.ConfigFileUsed()
	if(file != "") {
		log.Printf("Use config file: " + file)
	}
	return nil
}