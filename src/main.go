package main

import (
	"log"
	"os"
	"saviusz/melon-server/src/logger"

	"github.com/labstack/echo/v4"
)

func main() {
	logging := logger.NewLogger(os.Stdout, logger.LoggerOpts{
		TimeFormat: "15:04:05.000",
	})
	e := echo.New()

	e.GET("/", func(c echo.Context) error {
		logging.Log("Elo", logger.LogOpts{
			Level:  logger.LevelDebug,
			Status: 100,
		})
		logging.Log("Elo\nDruga linijka", logger.LogOpts{})
		logging.Log("Elo", logger.LogOpts{
			Level:  logger.LevelWarn,
			Status: 300,
		})
		logging.Log("Elo", logger.LogOpts{
			Level:  logger.LevelError,
			Status: 400,
		})
		return c.String(200, "Hello world")
	})

	log.Fatal(e.Start(":3000"))
}
