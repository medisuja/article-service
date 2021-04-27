package config

import (
	"github.com/jinzhu/configor"
)

var Config = struct {
	App struct {
		ENV      string
		HttpAddr string
		HttpPort string
		Url      string
	}

	DB struct {
		Driver       string
		Host         string
		Port         string `default:"3306"`
		Name         string
		User         string `default:"root"`
		Password     string `required:"true" env:"DBPassword"`
		Locale       string `default:"Asia/Jakarta"`
		MaxOpenConns int
	}

	REDIS struct {
		RedisHost     string
		RedisPassword string
		RedisPort     string
		ProductID     string
		ExpiredMinute int
	}

	Algolia struct {
		APPKEY         string
		APPKEYCREATE   string
		APPID          string
		HitsPerPage    int
		IndicesArticle string
	}
}{}

func init() {
	configor.Load(&Config, "config.yaml")
}
