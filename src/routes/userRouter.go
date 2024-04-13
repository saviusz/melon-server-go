package routes

import (
	"saviusz/melon-server/src/userRepo"

	"github.com/labstack/echo/v4"
)

type UserRouter struct {
	UR userRepo.UserRepo
}

func (router *UserRouter) Register(e *echo.Echo) {
	g := e.Group("/users")
	g.GET("/", router.getUsers)
	g.POST("/", router.createUser)
}

func (router *UserRouter) getUsers(c echo.Context) error {
	users := []userRepo.UserData{}
	for _, v := range router.UR.GetAll() {
		users = append(users, v.UserData)
	}
	return c.JSON(200, users)
}

func (router *UserRouter) createUser(c echo.Context) error {
	user, err := router.UR.CreateUser(userRepo.UserCreate{
		Name:  "Nowy",
		Email: "email@com.net",
	})

	if err != nil {
		return c.JSON(500, err)
	}

	return c.JSON(200, user)
}
