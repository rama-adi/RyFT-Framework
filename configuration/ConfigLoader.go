package configuration

import (
	"github.com/BurntSushi/toml"
)

// Configuration ---
//
// This struct is used to store the configuration for the framework
type Configuration struct {
	Application struct {
		Name string `toml:"name"`
		Port string `toml:"port"`
	} `toml:"application"`

	Security struct {
		Key        string `toml:"key"`
		DebugMode  bool   `toml:"debug_mode"`
		Production bool   `toml:"production"`
	} `toml:"security"`

	Database struct {
		Enabled  bool   `toml:"enabled"`
		Host     string `toml:"host"`
		Port     string `toml:"port"`
		Username string `toml:"username"`
		Password string `toml:"password"`
		Name     string `toml:"name"`
	} `toml:"database"`

	Authentication struct {
		Enabled           bool   `toml:"enabled"`
		AuthenticationUrl string `toml:"authentication_url"`
	} `toml:"authentication"`

	Caching struct {
		Enabled bool `toml:"enabled"`
	} `toml:"caching"`
}

// LoadConfigFile ---
//
// This function is used to load the configuration file
// Ryft by default uses TOML as the configuration file format
// It looks for the config file in the root directory
func LoadConfigFile() (Configuration, error) {

	var config Configuration
	_, err := toml.DecodeFile("./configuration/config.toml", &config)

	return config, err

}
