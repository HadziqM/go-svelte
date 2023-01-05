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


func apiRoute(c *fiber.App){
  c.Get("api/post",api.GetIndex)
  c.Get("api/post/:slug",api.GetPost)
}


func main() {
  //load config
  conf := config.LoadConf()

  // fiber init
  app := fiber.New()
  api.Index(app)
  api.SetCors(app,conf.Frontend)
  apiRoute(app)

  //sqlite init
  dbase,err := sql.Open("sqlite3","./sqlite.db")
  logger.Fatal(err,"error on main connect db")
  defer dbase.Close()

  if conf.Init{
    db.Init(dbase)
    db.Update(dbase,conf.Wordpress)
  }

  log.Fatal(app.Listen(":8000"))
}
