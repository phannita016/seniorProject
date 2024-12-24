package config

import (
	"log/slog"
	"strings"

	"github.com/spf13/viper"
)

type (
	AppConfig struct {
		Server         Server
		DatabaseServer DatabaseServer
	}

	DatabaseServer struct {
		ServiceName string
		URI         string
		Addr        string
		Username    string
		Password    string
		Database    string
	}

	Server struct {
		Address string
		Port    string
	}
)

func New(file string) (*AppConfig, error) {
	if file != "" {
		viper.SetConfigFile(file)
	} else {
		viper.SetConfigName("config")
		viper.AddConfigPath(".")
	}
	viper.SetConfigType("yaml")

	viper.AutomaticEnv()
	viper.SetEnvKeyReplacer(strings.NewReplacer(".", "_"))

	if err := viper.ReadInConfig(); err != nil {
		slog.Error("ReadInConfig", slog.Any("err", err))
		slog.Info("Load configuration in system environment.")
	}

	return &AppConfig{
		Server:         getServer(),
		DatabaseServer: getDatabaseServer(),
	}, nil

}

func getServer() Server {
	return Server{
		Address: viper.GetString("address"),
		Port:    viper.GetString("port"),
	}
}

func getDatabaseServer() DatabaseServer {
	return DatabaseServer{
		ServiceName: viper.GetString("service"),
		URI:         viper.GetString("mongo.uri"),
		Addr:        viper.GetString("mongo.host"),
		Username:    viper.GetString("mongo.username"),
		Password:    viper.GetString("mongo.password"),
		Database:    viper.GetString("mongo.database"),
	}
}
