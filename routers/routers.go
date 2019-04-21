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
	g.Use(middleware.KeyAuth(middlewares.ValidateToken))
	g.GET("/:id", handlers.GetID)
	g.POST("/:id", handlers.GetAll)

	e.Logger.Fatal(e.Start(":12345"))
}
