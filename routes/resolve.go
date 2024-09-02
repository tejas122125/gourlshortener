package routes

import (
	"urlshortener/database"

	"github.com/go-redis/redis/v8"
	"github.com/gofiber/fiber/v2"
)



func ResolveUrl(c *fiber.Ctx) error {

url := c.Params("url")

 r := database.CreateClient(0)

 defer r.Close()

 r.Get(database.ctx)


}