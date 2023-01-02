package main

import (
	"github.com/hadziqm/go-svelte/config"
	"github.com/hadziqm/go-svelte/db/wp"
	"github.com/hadziqm/go-svelte/logger"
)



func main() {
  config := config.LoadConf()
  categories := wp.GetCategories(config.Wordpress)
  logger.Print(categories.Data.Categories.Nodes[0].Name)
}
