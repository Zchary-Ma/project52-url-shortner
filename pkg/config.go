package pkg

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
