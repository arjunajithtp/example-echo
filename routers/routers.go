package routers

import (
	"github.com/arjunajithtp/example-echo/handlers"
	"github.com/arjunajithtp/example-echo/middlewares"
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
	g := e.Group("/id")

	// Authentication
	g.Use(middlewares.CustomAuth)

	g.GET("/:id", handlers.GetID)
	g.POST("/:id", handlers.GetAll)

	e.Logger.Fatal(e.Start(":1234"))
}
