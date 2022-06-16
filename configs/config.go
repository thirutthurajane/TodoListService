package configs

import (
	"context"
	"github.com/pkg/errors"
	"github.com/spf13/viper"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"time"
)

type Config struct {
	Server ServerConfiguration
}

type ServerConfiguration struct {
	Env        string
	DbName     string
	Secret     string
	Develop    EnvConfiguration
	Production EnvConfiguration
}

type EnvConfiguration struct {
	Port   string
	Db     string
	Github Github
}

type Github struct {
	Cid    string
	Secret string
}

var Configuration Config

func GetConn() (*mongo.Client, error) {
	var uri string

	switch Configuration.Server.Env {
	case "Develop":
		uri = Configuration.Server.Develop.Db
	case "Production":
		uri = Configuration.Server.Production.Db
	default:
		uri = ""
	}

	client, err := mongo.NewClient(options.Client().ApplyURI(uri))
	if err != nil {
		return nil, errors.Wrap(err, "Apply mongo uri error")
	}

	return client, nil
}

func init() {

	viper.AddConfigPath("./cfg")
	viper.SetConfigName("config")
	viper.SetConfigType("yaml")

	err := viper.ReadInConfig()
	if err != nil {
		errors.Wrap(err, "Config read error")
	}
	err = viper.Unmarshal(&Configuration)
	if err != nil {
		errors.Wrap(err, "Can not unmashal config")
	}
}

func GetContext() (context.Context, context.CancelFunc) {
	return context.WithTimeout(context.Background(), 10*time.Second)
}
