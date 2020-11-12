package utils

import "github.com/spf13/viper"

// Config is a structure holding all configuration data
type Config struct {
	Server  ConfigServer
	MongoDB ConfigMongo
}

// ConfigServer is a structure holding configuration for server
type ConfigServer struct {
	IP      string
	Port    string
	TypeDB  string
	Timeout ConfigServerTimeout
}

// ConfigServerTimeout is a structure holding configuration for server timeouts
type ConfigServerTimeout struct {
	Idle  int
	Read  int
	Write int
}

// ConfigMongo is a structure holding configuration for MongoDB
type ConfigMongo struct {
	Scheme   string
	IP       string
	Port     string
	Database ConfigMongoData
	Timeout  ConfigMongoTimeout
}

// ConfigMongoData is a structure holding configuration for MongoDB database
type ConfigMongoData struct {
	Name       string
	Collection string
}

// ConfigMongoTimeout is a structure holding configuration for MongoDB timeouts
type ConfigMongoTimeout struct {
	Connection int
	Get        int
	Post       int
	Put        int
	Delete     int
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
	viper.SetDefault("mongo.scheme", "mongodb://")
	viper.SetDefault("mongo.ip", "127.0.0.1")
	viper.SetDefault("mongo.port", "27017")
	viper.SetDefault("mongo.conn-time", 10)
	viper.SetDefault("mongo.oper-time", 5)
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
