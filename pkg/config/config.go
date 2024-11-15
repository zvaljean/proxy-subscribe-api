package config

import (
	"fmt"

	"github.com/spf13/viper"
)

type logCnf struct {
	Path  string `mapstructure:"path"`
	Debug bool   `mapstructure:"debug"`
}

type server struct {
	Port   int    `mapstructure:"port"`
	DbPath string `mapstructure:"dbPath"`
}

type bizCnf struct {
	UserPath string `mapstructure:"userPath"`
	Busy     int    `mapstructure:"busy"`
}

type Config struct {
	Server server `mapstructure:"server"`
	Log    logCnf `mapstructure:"log"`
	Biz    bizCnf `mapstructure:"biz"`
}

func LoadConfig(path string) (config Config, err error) {

	viper.AddConfigPath(path)
	viper.SetConfigType("yaml")
	viper.SetConfigName("app")

	err = viper.ReadInConfig()
	if err != nil {
		fmt.Printf("Cnf parse error! %s \n", err.Error())
	}

	viper.SetDefault("server.port", 8081)

	err = viper.Unmarshal(&config)
	return
}
