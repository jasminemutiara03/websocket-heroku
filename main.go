package main

import (
	"log"

	"github.com/jasminemutiara03/websocket-heroku/module"
	"github.com/jasminemutiara03/websocket-heroku/url"

	"github.com/gofiber/fiber/v2"
)

func main() {
	go module.RunHub()
	site := fiber.New()
	url.Web(site)
	log.Fatal(site.Listen(":3000"))
}
