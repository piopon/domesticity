package utils

import "github.com/spf13/viper"

//Config is a structure holding all configuration data
type Config struct {
	ServerIP     string
	ServerPort   string
	DataBaseType string
}

//NewConfig is a factory method to create configuration objects
func NewConfig() *Config {
	configDefaults()
	configInitialize()
	return &Config{
		ServerIP:     viper.GetString("service.ip"),
		ServerPort:   viper.GetString("service.port"),
		DataBaseType: viper.GetString("service.db-type"),
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
