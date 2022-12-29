package main

import (
	"database/sql"
	"log"

	"github.com/hadziqm/go-svelte/api"
	"github.com/hadziqm/go-svelte/config"
	"github.com/hadziqm/go-svelte/db"
	"github.com/hadziqm/go-svelte/logger"

	"github.com/gofiber/fiber/v2"
	_ "github.com/mattn/go-sqlite3"
)


func main() {
  //load config
  conf := config.LoadConf()

  // fiber init
  App := fiber.New()
  api.Index(App)
  api.SetCors(App,conf.Frontend)

  //sqlite init
  dbase,err := sql.Open("sqlite3","./test.db")
  logger.Fatal(err,"sucessfully connect to database")
  defer dbase.Close()

  if conf.Init{
    db.Init(dbase)
  }

  log.Fatal(App.Listen(":8000"))
}
