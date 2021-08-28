package tconfig

import (
	"fmt"
	"os"

	"github.com/fsnotify/fsnotify"
	"github.com/spf13/viper"
)

var (
	_config *viper.Viper
)

func init() {
	fmt.Println("init config file.")
	appenv := os.Getenv("APPENV")
	if appenv == "" {
		appenv = "production"
	}
	path := os.Getenv("ROOT_PATH") + "etc/" + appenv

	_config = viper.New()
	_config.SetConfigName("config")
	_config.SetConfigType("yaml")
	_config.AddConfigPath(path)
	if err := _config.ReadInConfig(); err != nil {
		panic(fmt.Sprintf("can't load config from path [%v] get error [%v]", path, err))
	}

	_config.WatchConfig()
	_config.OnConfigChange(func(e fsnotify.Event) {
		fmt.Println("config file changed: ", e.Name)
	})
}

func Get(key string) interface{} {
	return _config.Get(key)
}

func GetStringDefault(key, def string) string {
	val := _config.GetString(key)
	if val == "" {
		return def
	}
	return val
}

func GetString(key string) string {
	return _config.GetString(key)
}
