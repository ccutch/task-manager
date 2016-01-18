package taskmanager

import (
	"fmt"
	"io/ioutil"
	"os"
	"runtime"

	"gopkg.in/yaml.v2"
)

var globalConfig struct {
	User     string `yaml:"user"`
	Server   string `yaml:"server"`
	FileName string `yaml:"filename,omitempty"`
}

var ConfigComplete = false

func init() {
	LoadLastConfig()
}

func SetUser(u string) {
	globalConfig.User = u
	fmt.Println("getting yaml", u, globalConfig.User, globalConfig.Server)
	if globalConfig.Server != "" {
		ConfigComplete = true
	}
	SaveConfig()
}

func SetServer(s string) {
	globalConfig.Server = s
	if globalConfig.User != "" {
		ConfigComplete = true
	}
	SaveConfig()
}

func GetConfigYaml() string {
	d, _ := yaml.Marshal(globalConfig)
	return string(d)
}

func UserHomeDir() string {
	if runtime.GOOS == "windows" {
		home := os.Getenv("HOMEDRIVE") + os.Getenv("HOMEPATH")
		if home == "" {
			home = os.Getenv("USERPROFILE")
		}
		return home
	}
	return os.Getenv("HOME")
}

func LoadLastConfig() {
	fn := UserHomeDir() + "/.taskmanager.config"
	d, err := ioutil.ReadFile(fn)
	if err != nil {
		fmt.Println("Error loading config", err.Error())
	}
	yaml.Unmarshal(d, &globalConfig)
	if globalConfig.Server != "" && globalConfig.User != "" {
		ConfigComplete = true
	}
	if globalConfig.FileName == "" {
		globalConfig.FileName = ".tasks.yml"
	}
}

func SaveConfig() {
	c := GetConfigYaml()
	fn := UserHomeDir() + "/.taskmanager.config"
	err := ioutil.WriteFile(fn, []byte(c), 0777)
	if err != nil {
		fmt.Println("Error saving config", err.Error())
	}
}
