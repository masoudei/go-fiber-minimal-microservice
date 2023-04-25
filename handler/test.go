package handler

import (
	"github.com/gofiber/fiber/v2"
)

// Test example
//
//	@Summary		Hello World
//	@Description	Hello, World!
//	@ID				Test
//	@Accept			json
//	@Produce		json
//	@Success		200		{string}	string			"ok"
//	@Failure		400		{object}	string	"Hello, World!"
//	@Failure		404		{object}	string	"404"
//	@Router			/	[get]
func Test(c *fiber.Ctx) error {
	return c.SendString("Hello, World!")
}
