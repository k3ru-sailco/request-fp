package main

import (
	"github.com/gofiber/fiber/v2"
	tcp_data "github.com/k3ru-sailco/request-fp/lib"
)

func main() {
	/* start listening on the wifi network interface (i am using a macbook in this example)*/
	go tcp_data.StartListening("en0")
	router := fiber.New()
	router.Get("/data", func(c *fiber.Ctx) error {
		/* returns all the fingerprint data from incoming packets */
		return c.JSON(tcp_data.FINGERPRINTS)
	})
	router.Listen(":8080")
}
