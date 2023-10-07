package conf

import (
	"log"
	"os"

	"gopkg.in/yaml.v2"
)

// https://www.jianshu.com/p/84499381a7da
type Mysql struct {
	User     string `yaml:"user"`
	Password string `yaml:"password"`
	Host     string `yaml:"host"`
	Port     string `yaml:"port"`
	Name     string `yaml:"name"`
}

type AppConf struct {
	Debug  bool   `yaml:"debug"`
	DDpush string `yaml:"ddpush"`
	Port   string `yaml:"port"`
}

type CenterConf struct {
	Version string  `yaml:"version"`
	SqlCnf  Mysql   `yaml:"mysql"`
	App     AppConf `yaml:"app"`
}

var YamlConf CenterConf

func init() {
	yamlFile, err := os.ReadFile("conf.yaml")
	if err != nil {
		log.Printf("yamlFile.Get err #%v ", err)
	}

	err = yaml.Unmarshal(yamlFile, &YamlConf)
	if err != nil {
		log.Fatalf("Unmarshal: %v", err)
	}

	log.Println("conf", YamlConf)
}
