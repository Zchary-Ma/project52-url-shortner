package pkg

import (
	"fmt"
	"github.com/spf13/viper"
)

// Configurations exported
type Configurations struct {
	Redis RedisConfig
}

// RedisConfig  exported
type RedisConfig struct {
	Address  string
	Password string
	DB       int
}

func LoadConfig(c *Configurations) {
	viper.SetConfigFile("./config.yml")
	if err := viper.ReadInConfig(); err != nil {
		fmt.Printf("Error occurred while reading config file: %s", err)
	}
	err := viper.Unmarshal(&c)
	if err != nil {
		fmt.Printf("Error occurred while unmarshalling config file: %s", err)
	}
}
