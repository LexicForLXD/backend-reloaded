package config

import (
	"fmt"
	"io/ioutil"
	"log"

	"github.com/fsnotify/fsnotify"
	"github.com/spf13/viper"
)

func InitViper() {
	viper.SetConfigName("config")
	viper.AddConfigPath("./config")
	if err := viper.ReadInConfig(); err != nil {
		panic(err)
	}

	viper.WatchConfig()
	viper.OnConfigChange(func(e fsnotify.Event) {
		fmt.Println("Config file changed:", e.Name)
		readCerts()
	})

	readCerts()
	viper.SetDefault("rollbar.codeVersion", "v1")
	viper.SetDefault("rollbar.serverRoot", "github.com/lexicforlxd/backend-reloaded")
	viper.SetDefault("rollbar.environment", "development")
	viper.SetDefault("database.host", "localhost")
	viper.SetDefault("timeout", 30)
	viper.SetDefault("port", 8080)
}

func readCerts() {
	cert, err := ioutil.ReadFile(viper.GetString("tls.certFile"))
	key, err := ioutil.ReadFile(viper.GetString("tls.keyFile"))
	if err != nil {
		log.Fatal(err)
	}
	viper.Set("tls.cert", string(cert))
	viper.Set("tls.key", string(key))
}
