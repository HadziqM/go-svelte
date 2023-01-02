package config

import (
	"encoding/json"
	"os"

	"github.com/hadziqm/go-svelte/logger"
)


type Conf struct{
  Init bool `json:"init"`
  Frontend string `json:"frontend"`
  Wordpress string `json:"wordpress"`
}

func LoadConf() Conf {
  config,err := os.Open("./config.json")
  logger.Fatal(err,"error on config opening file")
  defer config.Close()
  jsonParser := json.NewDecoder(config)
  var conf Conf
  jsonParser.Decode(&conf)
  return conf
}
