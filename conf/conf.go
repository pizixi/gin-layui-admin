package conf

import (
	"log"
	"os"

	"gopkg.in/yaml.v2"
)

// AppConf 代表应用程序的配置部分
type AppConf struct {
	Debug  bool   `yaml:"debug"`
	DDpush string `yaml:"ddpush"`
	Port   string `yaml:"port"`
}

// CenterConf 代表整个配置文件的结构
type CenterConf struct {
	Version string  `yaml:"version"`
	DBLink  string  `yaml:"dblink"`
	App     AppConf `yaml:"app"`
}

var YamlConf CenterConf

// 默认的配置内容
var defaultConf = CenterConf{
	Version: "1.0.1",
	DBLink:  "sqlite:gin-layui-admin.sqlite3",
	App: AppConf{
		Debug:  true,
		DDpush: "https://oapi.dingtalk.com/robot/send?access_token=xxx",
		Port:   "8800",
	},
}

func init() {
	yamlFile, err := os.ReadFile("conf.yaml")
	if err != nil {
		// 如果配置文件不存在，则创建并写入默认配置
		if os.IsNotExist(err) {
			defaultYaml, err := yaml.Marshal(&defaultConf)
			if err != nil {
				log.Fatalf("Error marshaling default config: %v", err)
			}

			err = os.WriteFile("conf.yaml", defaultYaml, 0644)
			if err != nil {
				log.Fatalf("Error writing default config to file: %v", err)
			}

			yamlFile = defaultYaml // 使用默认配置进行Unmarshal
			log.Println("Default configuration file created")
		} else {
			log.Fatalf("yamlFile.Get err #%v ", err)
		}
	}

	err = yaml.Unmarshal(yamlFile, &YamlConf)
	if err != nil {
		log.Fatalf("Unmarshal: %v", err)
	}

	log.Println("conf", YamlConf)
}
