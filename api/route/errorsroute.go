package route

import (
	"github.com/Drakkar-Software/Errors-Server/api/controller"

	"github.com/labstack/echo/v4"
)

// Init registers the server routes
func Init(e *echo.Echo) {
	e.POST("/error", controller.PostError)
}
