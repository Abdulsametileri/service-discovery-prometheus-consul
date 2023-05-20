package main

import (
	"flag"
	"fmt"
	"github.com/gofiber/adaptor/v2"
	"github.com/gofiber/fiber/v2"
	"github.com/prometheus/client_golang/prometheus/promhttp"
)

var (
	addr = flag.String("listen", ":8080", "The address to listen on for HTTP requests.")
)

func main() {
	flag.Parse()

	app := fiber.New(fiber.Config{DisableStartupMessage: true})
	app.Get("/metrics", adaptor.HTTPHandler(promhttp.Handler()))

	fmt.Println("Server is starting on port " + *addr)

	err := app.Listen(*addr)
	fmt.Println(err)
}
