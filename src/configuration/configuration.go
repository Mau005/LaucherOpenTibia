package configuration

import (
	"fmt"
	"os"

	"github.com/go-yaml/yaml"
)

var CONFIG *Configuration
var API *Api

type Api struct {
	TitleWindow string `yaml:"TitleWindow"`
	NameApp     string `yaml:"NameApp"`
	Version     string `yaml:"Version"`
	PathCLient  string `yaml:"PathCLient"`
	PathLogo    string `yaml:"PathLogo"`
	NameButton  string `yaml:"NameButton"`
}

type Configuration struct {
	Api Api `yaml:"Configuration"`
}

func Load(fileName string) error {
	config := Configuration{}

	content, err := os.ReadFile(fileName)
	if err != nil {
		return err
	}

	err = yaml.Unmarshal(content, &config)
	if err != nil {
		return err
	}
	fmt.Println(config)
	CONFIG = &config
	API = &config.Api

	return nil
}
