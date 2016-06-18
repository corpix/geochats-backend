package config

//Config represents global configuration of application
type Config struct {
	ListenAddr   string
	DatabaseAddr string
	DatabaseName string
}

var c *Config

// Get returns the configuration object
func Get() *Config {
	if c == nil {
		panic("Config is not set yet!")
	}
	return c
}

// Set updates current config with new value
func Set(config *Config) {
	c = config
}
