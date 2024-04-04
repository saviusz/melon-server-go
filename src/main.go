package main

import (
	"log"
	"os"
	"saviusz/melon-server/src/logger"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func main() {
	logging := logger.NewLogger(os.Stdout, logger.LoggerOpts{
		TimeFormat: "15:04:05.000",
	})

	e := echo.New()
	e.Use(middleware.RequestLoggerWithConfig(middleware.RequestLoggerConfig{
		HandleError: true,
		LogURI:      true,
		LogMethod:   true,
		LogError:    true,
		LogStatus:   true,
		LogValuesFunc: func(c echo.Context, v middleware.RequestLoggerValues) error {
			if v.Error != nil {
				logging.Log(v.Error.Error(), logger.LogOpts{
					Level:  logger.LevelError,
					Status: uint16(v.Status),
					Method: v.Method,
					Path:   v.URI,
				})
			} else {
				logging.Log("", logger.LogOpts{
					Level:  logger.LevelInfo,
					Status: uint16(v.Status),
					Method: v.Method,
					Path:   v.URI,
				})
			}

			return nil
		},
	}))

	e.GET("/", func(c echo.Context) error {
		return c.String(200, "Hello world")
	})

	log.Fatal(e.Start(":3000"))
}
