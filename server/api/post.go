package api

import (
	"github.com/gofiber/fiber/v2"
	"github.com/hadziqm/go-svelte/db"
)




func GetAll(c *fiber.Ctx) error {
  dbase := db.DbConn()
  defer dbase.Close()
  return nil
}
