package configuration

// Config - Struct to hold all configurable bot settings
type Config struct {
	Path           string
	Botconfig      botConfig      `yaml:"beerbot"`
	Databaseconfig databaseConfig `yaml:"database"`
}
type botConfig struct {
	RunAsDaemon bool   `yaml:"daemon"`
	Token       string `yaml:"token"`
}

type databaseConfig struct {
	Path string `yaml:"path"`
}

// SetPath - set the path for where the config file is located
func (c *Config) SetPath(path string) {
	c.Path = path
}

// Parse - Parses config located at config path
func (c *Config) Parse() error {
	return nil
}
