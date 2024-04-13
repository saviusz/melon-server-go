package main

import (
	"log"
	database "saviusz/melon-server/src/db"
	"saviusz/melon-server/src/routes"
	"saviusz/melon-server/src/userRepo"

	"github.com/joho/godotenv"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/upper/db/v4"
)

func main() {

	err := godotenv.Load(".env")
	if err != nil {
		log.Fatalf("Error loading .env file: %s", err)
	}

	e := echo.New()
	e.Pre(middleware.AddTrailingSlash())

	sess, err := database.New()
	if err != nil {
		log.Fatal("Error: ", err)
	}

	userRepo := userRepo.New(sess)
	userRouter := routes.UserRouter{UR: userRepo}

	// Uncomment the following line (and the github.com/upper/db import path) to
	// write SQL statements to os.Stdout:
	db.LC().SetLevel(db.LogLevelDebug)

	userRouter.Register(e)

	log.Fatal(e.Start(":3000"))
}
