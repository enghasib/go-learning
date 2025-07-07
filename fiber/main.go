package main

import (
	"fmt"
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
)

func main() {
	//create a app using fiber
	app := fiber.New()
	app.Use(logger.New())

	app.Get("/", func(c *fiber.Ctx) error {
		helloInfo := c.ClientHelloInfo()
		fmt.Println("Client hello info:", helloInfo)

		return c.SendString("The router are working perfectly")
	}).Name("get stack")

	log.Fatal(app.Listen(":4000"))
}
