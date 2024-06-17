package config

import (
	"fmt"
	"log"
	"sync"
	"time"

	"github.com/spf13/viper"
)

var cfg Config
var doOnce sync.Once

type Config struct {
	Application struct {
		Port int `mapstructure:"port"`
		Log  struct {
			Path      string `mapstructure:"path"`
			MaxSize   int    `mapstructure:"maxSize"`
			MaxBackup int    `mapstructure:"maxBackup"`
			MaxAge    int    `mapstructure:"maxAge"`
			Compress  bool   `mapstructure:"compress"`
		} `mapstructure:"log"`
		Key struct {
			Default string `mapstructure:"default"`
			RSA     struct {
				Public  string `mapstructure:"public"`
				Private string `mapstructure:"private"`
			} `mapstructure:"rsa"`
		} `mapstructure:"key"`
		Graceful struct {
			MaxSecond time.Duration `mapstructure:"maxSecond"`
		} `mapstructure:"graceful"`
	} `mapstructure:"application"`

	Auth struct {
		JWTToken struct {
			Type           string `mapstructure:"type"`
			Expired        string `mapstructure:"expired"`
			RefreshExpired string `mapstructure:"refreshExpired"`
		} `mapstructure:"jwtToken"`
	} `mapstructure:"auth"`

	Database struct {
		MySQL struct {
			User            string `mapstructure:"user"`
			Password        string `mapstructure:"password"`
			Host            string `mapstructure:"host"`
			Port            int    `mapstructure:"port"`
			Name            string `mapstructure:"name"`
			MaxIdleConns    int    `mapstructure:"maxIdleConns"`
			MaxOpenConns    int    `mapstructure:"maxOpenConns"`
			ConnMaxLifeTime int    `mapstructure:"connMaxLifeTime"`
			ConnMaxIdleTime int    `mapstructure:"connMaxIdleTime"`
		} `mapstructure:"mysql"`
		Postgres struct {
			User     string `mapstructure:"user"`
			Password string `mapstructure:"password"`
			Host     string `mapstructure:"host"`
			Port     int    `mapstructure:"port"`
			Name     string `mapstructure:"name"`
		} `mapstructure:"postgres"`
	} `mapstructure:"database"`

	Cache struct {
		Redis struct {
			Host     string `mapstructure:"host"`
			Port     int    `mapstructure:"port"`
			Password string `mapstructure:"password"`
		} `mapstructure:"redis"`
	} `mapstructure:"cache"`

	Queue struct {
		Rabbit struct {
			Host     string `mapstructure:"host"`
			User     string `mapstructure:"user"`
			Password string `mapstructure:"password"`
			Port     int    `mapstructure:"port"`
		} `mapstructure:"rabbit"`
	} `mapstructure:"queue"`

	Module struct {
		News struct {
			Scheduler struct {
				GetNewsRSSEnable     bool `mapstructure:"getNewsRSSEnable"`
				GetNewsRSSDuration   int  `mapstructure:"getNewsRSSDuration"`
				GetNewsIndexEnable   bool `mapstructure:"getNewsIndexEnable"`
				GetNewsIndexDuration int  `mapstructure:"getNewsIndexDuration"`
			} `mapstructure:"scheduler"`
			Service struct {
				UrlNewsRSS   string `mapstructure:"urlNewsRSS"`
				UrlNewsIndex string `mapstructure:"urlNewsIndex"`
			} `mapstructure:"service"`
		} `mapstructure:"news"`
		User struct {
			Pubsub struct {
				UserCreatedEnable bool `mapstructure:"userCreatedEnable"`
			} `mapstructure:"pubsub"`
		} `mapstructure:"user"`
	} `mapstructure:"module"`
}

func Get() Config {
	viper.SetConfigName("config")
	viper.AddConfigPath(".")
	viper.AutomaticEnv()
	viper.SetConfigType("yaml")

	err := viper.ReadInConfig()
	if err != nil {
		log.Fatalln(fmt.Sprintf("Cannot read config file: %v", err))
	}

	doOnce.Do(func() {
		err := viper.Unmarshal(&cfg)
		if err != nil {
			log.Fatalln("Cannot unmarshaling config")
		}
	})

	return cfg
}
