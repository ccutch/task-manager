package gtm

import (
	"fmt"
	"io/ioutil"
	"os"
	"runtime"

	"gopkg.in/yaml.v2"
)

type Config struct {
	User     string `yaml:"user"`
	FileName string `yaml:"filename,omitempty"`
}

var GlobalConfig = &Config{FileName: ".tasks.yml"}

// Set User for global config
func (c *Config) SetUser(u string) {
	GlobalConfig.User = u
	SaveConfig()
}

// Return a yaml representation of the global config
func GetConfigYaml() string {
	d, _ := yaml.Marshal(GlobalConfig)
	return string(d)
}

// Get user home directory to save config file
func userHomeDir() string {
	if runtime.GOOS == "windows" {
		home := os.Getenv("HOMEDRIVE") + os.Getenv("HOMEPATH")
		if home == "" {
			home = os.Getenv("USERPROFILE")
		}
		return home
	}
	return os.Getenv("HOME")
}

// Load global config
func LoadLastConfig() {
	fn := userHomeDir() + "/.gtmrc"
	d, _ := ioutil.ReadFile(fn)
	yaml.Unmarshal(d, &GlobalConfig)
}

// Save global config
func SaveConfig() {
	c := GetConfigYaml()
	fn := userHomeDir() + "/.gtmrc"
	err := ioutil.WriteFile(fn, []byte(c), 0777)
	if err != nil {
		fmt.Println("Error saving config", err.Error())
	}
}

// init, called when package is imported
func init() {
	LoadLastConfig()
}
