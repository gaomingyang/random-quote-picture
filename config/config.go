package config

import (
	"fmt"
	"github.com/fsnotify/fsnotify"
	"github.com/spf13/viper"
	"log"
	"path/filepath"
)

func Init(workPath string) {
	log.Println("config init.")
	configPath := filepath.Join(workPath, "conf/app.yaml")
	err := initConfig(configPath)
	if err != nil {
		panic(err)
	}
}

func initConfig(filepath string) error {
	viper.SetConfigType("yaml")
	viper.SetConfigFile(filepath)
	err := viper.ReadInConfig()
	if err != nil {
		if _, ok := err.(viper.ConfigFileNotFoundError); ok {
			// Config file not found; ignore error if desired
			log.Printf("config file %s not exist", filepath)
		} else {
			// Config file was found but another error was produced
			panic(fmt.Errorf("fatal error config file: %w", err))
		}
	}

	viper.OnConfigChange(func(e fsnotify.Event) {
		log.Println("Config file changed:", e.Name)
	})
	viper.WatchConfig()
	return err
}
