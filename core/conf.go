package core

import (
	"backend/config"
	"fmt"
	"gopkg.in/yaml.v3"

	"log"
	"os"
)

func InitConf() *config.Config {
	c := &config.Config{}
	confPath := "settings.yaml"
	yamlConf, err := os.ReadFile(confPath)
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
