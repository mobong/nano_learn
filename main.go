package main

import (
	"github.com/spf13/viper"
	"github.com/urfave/cli"
	"nano_learn/src"
	"os"
)

func main() {
	app := cli.NewApp()
	app.Name = "nano learn"
	app.Version = "2023.0.1"
	app.Action = server
	app.Run(os.Args)
}

func server(*cli.Context) error {
	viper.SetConfigType("toml")
	viper.SetConfigFile("config/config.toml")
	viper.ReadInConfig()
	src.StartUp()
	return nil
}
