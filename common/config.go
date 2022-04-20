package common

import (
	"fmt"
	"gonitor/monitors"
	yaml2 "gopkg.in/yaml.v2"
	"io/ioutil"
)

type Config struct {
	Monitors []struct {
		Name       string
		Type       monitors.MonitorType
		Properties map[string]string
	}
}

func LoadConfig(file string) Config {
	conf := Config{}
	yaml, err := ioutil.ReadFile(file)
	if err != nil {
		fmt.Println("Failed to read file")
		panic(err)
	}
	err = yaml2.Unmarshal(yaml, &conf)
	if err != nil {
		fmt.Println("Failed to unmarshal yaml content")
		panic(err)
	}
	return conf
}
