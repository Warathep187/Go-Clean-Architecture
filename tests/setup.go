package test_setup

import (
	"go-clean-arch/config"
	"strings"
	"sync"

	"github.com/spf13/viper"
)

type (
	Config struct {
		Server *Server
		Db     *Db
	}

	Server struct {
		Port int
	}

	Db struct {
		Host     string
		Port     int
		User     string
		Password string
		DBName   string
		SSLMode  string
		TimeZone string
	}
)

var (
	once           sync.Once
	configInstance *config.Config
)

func GetTestConfigs() *config.Config {
	once.Do(func() {
		viper.SetConfigName("config-test")
		viper.SetConfigType("yaml")
		viper.AddConfigPath("../../")
		viper.AutomaticEnv()
		viper.SetEnvKeyReplacer(strings.NewReplacer(".", "_"))

		if err := viper.ReadInConfig(); err != nil {
			panic(err)
		}

		if err := viper.Unmarshal(&configInstance); err != nil {
			panic(err)
		}
	})

	return configInstance
}
