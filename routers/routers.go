package routers

import (
	"fmt"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

// Router controlls the routing
func Router() {
	e := echo.New()

	// Middleware
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	// Routes
	e.GET("/users", createUser)

	e.Logger.Fatal(e.Start(":8000"))
}

func createUser(c echo.Context) error {
	// u := &user{
	// 	ID: seq,
	// }
	// if err := c.Bind(u); err != nil {
	// 	return err
	// }
	// users[u.ID] = u
	// seq++
	// return c.JSON(http.StatusCreated, u)

	return fmt.Errorf("error ghgjghjghjghjghjhgjghjhgjghjghjhgjhgjhgjgh")
}
