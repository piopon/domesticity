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
		Server:  getConfigServer(),
		MongoDB: getConfigMongo(),
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

// getConfigServer is used to read current sever settings from config file
func getConfigServer() ConfigServer {
	return ConfigServer{
		IP:     viper.GetString("server.ip"),
		Port:   viper.GetString("server.port"),
		TypeDB: viper.GetString("server.db-type"),
		Timeout: ConfigServerTimeout{
			Idle:  viper.GetInt("server.timeout.idle"),
			Read:  viper.GetInt("server.timeout.write"),
			Write: viper.GetInt("server.timeout.read"),
		},
	}
}

// getConfigServer is used to read current MongoDB settings from config file
func getConfigMongo() ConfigMongo {
	return ConfigMongo{
		Scheme: viper.GetString("mongo.scheme"),
		IP:     viper.GetString("mongo.ip"),
		Port:   viper.GetString("mongo.port"),
		Database: ConfigMongoData{
			Name:       viper.GetString("mongo.db.name"),
			Collection: viper.GetString("mongo.db.collection"),
		},
		Timeout: ConfigMongoTimeout{
			Connection: viper.GetInt("mongo.timeout.connection"),
			Get:        viper.GetInt("mongo.timeout.get"),
			Post:       viper.GetInt("mongo.timeout.post"),
			Put:        viper.GetInt("mongo.timeout.put"),
			Delete:     viper.GetInt("mongo.timeout.delete"),
		},
	}
}
