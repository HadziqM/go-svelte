package api

import(
"github.com/gofiber/fiber/v2"
)

func Ada() string{
  return "tested"
}
func Index(App *fiber.App){
  App.Get("/api", func (c *fiber.Ctx) error {
        return c.SendString("Hello, World!")
    })
}
