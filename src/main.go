package main

import (
	"log"

	"github.com/labstack/echo/v4"
)

func main() {
	e := echo.New()

	e.GET("/", func(c echo.Context) error {
		return c.String(200, "Hello world")
	})

	log.Fatal(e.Start(":3000"))
}