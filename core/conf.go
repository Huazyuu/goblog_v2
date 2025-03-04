package core

import (
	"fmt"
	"gopkg.in/yaml.v3"
	"gvb_server/config"
	"log"
	"os"
)

const ConfigFile = "settings.yaml"

func initConf() *config.Config {
	c := &config.Config{}
	yamlConf, err := os.ReadFile(ConfigFile)
	if err != nil {
		panic(fmt.Errorf("get yamlConf error: %s", err))
	}
	err = yaml.Unmarshal(yamlConf, c)
	if err != nil {
		log.Fatalf("config Init Unmarshal: %v", err)
	}
	log.Println("config yamlFile load Init success.")
	return c
}
