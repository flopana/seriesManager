package config

import (
	"log"
)

type Config struct {
	VlcPath string
}

func ParseConfig(logger *log.Logger) Config {
	var config Config
	//b, _ := ioutil.ReadFile("config.json")
	//err := json.Unmarshal(b, &config)
	//if err != nil {
	//	logger.Fatalf("%s", err)
	//}
	return config
}