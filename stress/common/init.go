package common

import (
	"github.com/spf13/viper"
	"github.com/mitchellh/go-homedir"
	"fmt"
	"os"
	"steve/servicelauncher/loggerwin"
	"flag"
	"github.com/Sirupsen/logrus"
)

var (
	config  = flag.String("config", "./config.yml", "config.yml")
	Waitc chan struct{}
)

func Init() {
	Waitc = make(chan struct{})
	flag.Parse()
	configFile := initConfig()
	initLogger()
	if configFile != "" {
		logrus.WithField("config", configFile).Info("using config file")
		return
	}
}

func initConfig() string {
	if *config != "" {
		// Use config file from the flag.
		viper.SetConfigFile(*config)
	} else {
		// Find home directory.
		home, err := homedir.Dir()
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}

		// Search config in home directory with name ".serviceloader" (without extension).
		viper.AddConfigPath(home)
		viper.SetConfigName(".serviceloader")
	}

	viper.AutomaticEnv() // read in environment variables that match

	// If a config file is found, read it in.
	if err := viper.ReadInConfig(); err == nil {
		return viper.ConfigFileUsed()
	}
	return ""
}

func initLogger() {
	loggerwin.SetupLog(viper.GetString("log_prefix"), viper.GetString("log_dir"),
		viper.GetString("log_level"), viper.GetBool("log_stderr"))
	//loggerwin.SetupLog("stress", "./log", "debug", true)
}