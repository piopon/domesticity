package utils

//Config is a structure holding all configuration data
type Config struct {
	ServerIP   string
	ServerPort string
}

//NewConfig is a factory method to create configuration objects
func NewConfig() *Config {
	return &Config{ServerIP: "", ServerPort: "9999"}
}
