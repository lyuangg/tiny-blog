package configs

import (
	"fmt"

	"github.com/fsnotify/fsnotify"
	"github.com/spf13/viper"
)

// Config global
type Config struct {
	Host string   `mapstructure:"host"`
	Port int      `mapstructure:"port"`
	Mode string   `mapstructure:"mode"`
	Blog blog     `mapstructure:"blog"`
	Db   database `mapstructure:"db"`
	Log  log      `mapstructure:"log"`
}

//blog config
type blog struct {
	Domain   string `mapstructure:"domain"`
	Name     string `mapstructure:"name"`
	Author   string `mapstructure:"author"`
	Desc     string `mapstructure:"desc"`
	Keywords string `mapstructure:"keywords"`
	PageSize int    `mapstructure:"pagesize"`
}

// database config
type database struct {
	Host     string `mapstructure:"host"`
	Port     int    `mapstructure:"port"`
	Name     string `mapstructure:"name"`
	User     string `mapstructure:"user"`
	Password string `mapstructure:"password"`
	Maxconns int    `mapstructure:"maxconns"`
}

// log config
type log struct {
	Path  string `mapstructure:"path"`
	Level string `mapstructure:"level"`
	Days  int64  `mapstructure:"days"`
}

// Conf is config instance
var Conf = new(Config)

// LoadConfig init the config instance
func LoadConfig() {
	viper.SetConfigFile("./config.yaml")
	err := viper.ReadInConfig()
	if err != nil {
		panic(fmt.Errorf("Fatal error config file: %s ", err))
	}

	if err := viper.Unmarshal(Conf); err != nil {
		panic(fmt.Errorf("Unmarshal conf failed, err:%s ", err))
	}

	viper.WatchConfig()
	viper.OnConfigChange(func(e fsnotify.Event) {
		fmt.Println("Config file changed:", e.Name)
		if err := viper.Unmarshal(Conf); err != nil {
			panic(fmt.Errorf("Unmarshal conf failed, err:%s ", err))
		}
	})
}
