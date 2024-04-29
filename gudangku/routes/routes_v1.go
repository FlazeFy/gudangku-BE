package routes

import (
	// middlewares "gudangku/middlewares/jwt"
	authhandlers "gudangku/modules/auth/http_handlers"
	stshandlers "gudangku/modules/stats/http_handlers"
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

	// Auth
	e.POST("api/v1/login", authhandlers.PostLoginUser)
	e.POST("api/v1/register", authhandlers.PostRegister)
	e.POST("api/v1/logout", authhandlers.SignOut)

	// Stats
	e.GET("api/v1/stats/total_inventory_by_category", stshandlers.GetTotalInventoryByCategory)
	e.GET("api/v1/stats/total_inventory_by_favorite", stshandlers.GetTotalInventoryByFavorite)
	e.GET("api/v1/stats/total_inventory_by_room", stshandlers.GetTotalInventoryByRoom)

	// =============== Private routes ===============

	return e
}
