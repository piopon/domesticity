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
	configInitialize()
	configServerDefaults()
	configMongoDefaults()
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

// configInitialize is used to initialize Viper configs framework
func configInitialize() {
	viper.SetConfigType("yaml")
	viper.SetConfigName("settings")
	viper.AddConfigPath("scripts")
	readError := viper.ReadInConfig()
	if readError != nil {
		panic("Cannot read configuration file: " + readError.Error())
	}
}

// configServerDefaults is used to setup defaults for server configuration
func configServerDefaults() {
	viper.SetDefault("server.ip", "")
	viper.SetDefault("server.port", "9999")
	viper.SetDefault("server.db-type", "mongo")
	viper.SetDefault("server.timeout.idle", 600)
	viper.SetDefault("server.timeout.write", 10)
	viper.SetDefault("server.timeout.read", 10)
}

// configMongoDefaults is used to setup defaults for MongoDB configuration
func configMongoDefaults() {
	viper.SetDefault("mongo.scheme", "mongodb://")
	viper.SetDefault("mongo.ip", "127.0.0.1")
	viper.SetDefault("mongo.port", "27017")
	viper.SetDefault("mongo.db.name", "event-service")
	viper.SetDefault("mongo.db.collection", "events")
	viper.SetDefault("mongo.timeout.connection", 10)
	viper.SetDefault("mongo.timeout.get", 10)
	viper.SetDefault("mongo.timeout.post", 5)
	viper.SetDefault("mongo.timeout.put", 5)
	viper.SetDefault("mongo.timeout.delete", 3)
}

