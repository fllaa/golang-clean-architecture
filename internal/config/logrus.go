package config

import (
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
)

func NewLogger(viper *viper.Viper) *logrus.Logger {
	log := logrus.New()
	format := viper.GetString("log.format")

	log.SetLevel(logrus.Level(viper.GetInt32("log.level")))
	switch format {
	case "json":
		log.SetFormatter(&logrus.JSONFormatter{})
	case "text":
		log.SetFormatter(&logrus.TextFormatter{
			FullTimestamp: true,
		})
	}

	return log
}
