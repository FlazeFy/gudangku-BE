package routes

import (
	syshandlers "gudangku/modules/systems/http_handlers"
	"net/http"

	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
)

func InitV1() *echo.Echo {
	e := echo.New()
	e.Use(middleware.CORS())

	e.GET("api/", func(c echo.Context) error {
		return c.String(http.StatusOK, "Welcome to Gudangku")
	})

	// =============== Public routes ===============

	// Dictionary
	e.GET("api/v1/dct/:ord", syshandlers.GetDictionaryByType)

	// =============== Private routes ===============

	return e
}
