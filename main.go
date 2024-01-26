package main

import "github.com/gofiber/fiber/v2"

func main() {
	app := fiber.New()
	app.Static("/","./html")
	app.Get("/*",func(c *fiber.Ctx) error{
		return c.SendFile("./html/index.html")
	})

	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("hello, World")
	})
	app.Listen((":9090"))

}
