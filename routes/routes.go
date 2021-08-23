package routes

import (
	"github.com/Drakkar-Software/Errors-Server/api/route"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

// Init initializes the echo routes and registers the middleware
func Init(e *echo.Echo) {
	e.Pre(middleware.RemoveTrailingSlash())
	// allow a max request size of 17000 characters (allows 100 errors with ~100 depth stack traces)
	e.Use(middleware.BodyLimitWithConfig(middleware.BodyLimitConfig{Limit: "1700K"}))
	route.Init(e)
}
