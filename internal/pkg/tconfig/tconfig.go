package tconfig

import (
	"fmt"
	"os"
	"path"

	"github.com/fsnotify/fsnotify"
	"github.com/spf13/viper"
)

const (
	EnvProduction  = "production"
	EnvTest        = "test"
	EnvDevelopment = "development"
)

type TConfig struct {
	_config *viper.Viper
}

func New() *TConfig {
	fmt.Println("init config file.")
	appenv := os.Getenv("APPENV")
	if appenv == "" {
		appenv = EnvDevelopment
	}
	path := path.Join(os.Getenv("ROOT_PATH"), "etc/", appenv)

	_config := viper.New()
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
	_config.Set("APPENV", appenv)

	return &TConfig{
		_config: _config,
	}
}

func (config *TConfig) Get(key string) interface{} {
	return config._config.Get(key)
}

func (config *TConfig) GetStringDefault(key, def string) string {
	val := config._config.GetString(key)
	if val == "" {
		return def
	}
	return val
}

func (config *TConfig) GetString(key string) string {
	return config._config.GetString(key)
}

func (config *TConfig) GetStringSlice(key string) []string {
	return config._config.GetStringSlice(key)
}
