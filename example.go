package main

import (
	"github.com/gofiber/fiber/v2"
	request_fp "github.com/k3ru-sailco/request-fp/lib"
)

func main() {
	/* start listening on the wifi network interface (i am using a macbook in this example)*/
	go request_fp.StartListening("en0")
	router := fiber.New()
	router.Get("/fingerprint-me", func(c *fiber.Ctx) error {
		/* returns the fingerprint of the incoming request */
		data, _ := request_fp.RequestFP(c.IP())
		return c.JSON(data)
	})
	router.Listen(":8080")
}
