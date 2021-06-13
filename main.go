package main

import (
	"fmt"
	"github.com/spf13/viper"
	"github.com/zchary-ma/url-shortener/config"
	"github.com/zchary-ma/url-shortener/storage"
)

func main() {
	var configurations config.Configurations
	loadConfig(&configurations)
	storage := storage.CreateClient(configurations)
	item, _ := storage.Get("demo:version")
	fmt.Println(item)
}

func loadConfig(c *config.Configurations) {
	viper.SetConfigFile("./config.yml")
	if err := viper.ReadInConfig(); err != nil {
		fmt.Printf("Error occurred while reading config file: %s", err)
	}
	err := viper.Unmarshal(&c)
	if err != nil {
		fmt.Printf("Error occurred while unmarshalling config file: %s", err)
	}
}
