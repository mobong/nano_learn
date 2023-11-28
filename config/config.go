package config

import (
	"github.com/spf13/viper"
	"os"
	"os/signal"
	"syscall"
)

func init() {
	initConfig()
	go reloadConfig()
}

func initConfig() {
	viper.SetConfigType("toml")
	viper.SetConfigFile("config.toml")
	viper.ReadInConfig()
}

func reloadConfig() {
	sg := make(chan os.Signal, 1)
	signal.Notify(sg, syscall.SIGUSR1)
	for {
		s := <-sg
		switch s {
		case syscall.SIGUSR1:
			initConfig()
		}
	}

}
