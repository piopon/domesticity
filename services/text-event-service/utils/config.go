package utils

import "github.com/spf13/viper"

// Config is a structure holding all configuration data
type Config struct {
	Service ConfigService
	MongoDB ConfigMongo
}

// ConfigService is a structure holding configuration for service
type ConfigService struct {
	IP     string
	Port   string
	TypeDB string
}

// ConfigMongo is a structure holding configuration for MongoDB
type ConfigMongo struct {
	Scheme string
	IP     string
	Port   string
}

// NewConfig is a factory method to create configuration objects
func NewConfig() *Config {
	configDefaults()
	configInitialize()
	return &Config{
		Service: ConfigService{
			IP:     viper.GetString("service.ip"),
			Port:   viper.GetString("service.port"),
			TypeDB: viper.GetString("service.db-type"),
		},
		MongoDB: ConfigMongo{
			Scheme: viper.GetString("mongo.scheme"),
			IP:     viper.GetString("mongo.ip"),
			Port:   viper.GetString("mongo.port"),
		},
	}
}

func configDefaults() {
	viper.SetDefault("service.ip", "")
	viper.SetDefault("service.port", "9999")
	viper.SetDefault("service.db-type", "mongo")
}

func configInitialize() {
	viper.SetConfigType("yaml")
	viper.SetConfigName("settings")
	viper.AddConfigPath("scripts")
	readError := viper.ReadInConfig()
	if readError != nil {
		panic("Cannot read configuration file: " + readError.Error())
	}
}
