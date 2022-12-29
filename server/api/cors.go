package api

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
)


func SetCors(app *fiber.App,adrress string){
  app.Use(cors.New(cors.Config{
    AllowOrigins: adrress,
    AllowHeaders: "Origin, Content-Type, Accept",
  }))
}
