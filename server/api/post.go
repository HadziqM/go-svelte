package api

import (
	"github.com/gofiber/fiber/v2"
	"github.com/hadziqm/go-svelte/db"
)

func GetIndex(c *fiber.Ctx) error {
  dbase := db.DbConn()
  defer dbase.Close()
  data := db.GetIndex(&dbase)
  return c.JSON(data)
}
func GetPost(c *fiber.Ctx)error{
  dbase := db.DbConn()
  defer dbase.Close()
  param := c.Params("slug")
  data := db.GetSpecificPost(&dbase,param)
  return c.JSON(data)
}
