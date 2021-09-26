package main

import (
	"net"
	"os"
	"strings"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/utils"
)

type Response struct {
	Version     string `json:"version"`
	RequestID   string `json:"request_id"`
	HostName    string `json:"host_name"`
	HostAddress string `json:"host_address"`
}

func main() {
	app := fiber.New()

	app.Get("/", func(ctx *fiber.Ctx) error {
		hostName, err := os.Hostname()
		if err != nil {
			panic(err)
		}

		addrs, err := net.LookupHost(hostName)
		if err != nil {
			panic(err)
		}

		hostAddres := strings.Join(addrs, " ")

		return ctx.JSON(Response{
			Version:     "1.0",
			RequestID:   utils.UUIDv4(),
			HostName:    hostName,
			HostAddress: hostAddres,
		})
	})

	panic(app.Listen(":5000"))
}
