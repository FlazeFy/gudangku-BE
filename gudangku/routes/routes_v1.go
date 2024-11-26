package routes

import (
	middlewares "gudangku/middlewares/jwt"
	authhandlers "gudangku/modules/auth/http_handlers"
	invhandlers "gudangku/modules/inventories/http_handlers"
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
	e.GET("api/v1/dct/:type", syshandlers.GetDictionaryByType)

	// Auth
	e.POST("api/v1/login", authhandlers.PostLoginUser)
	e.POST("api/v1/register", authhandlers.PostRegister)
	e.POST("api/v1/logout", authhandlers.SignOut, middlewares.CustomJWTAuth)
	e.GET("api/v1/user/my_profile", authhandlers.GetMyProfile, middlewares.CustomJWTAuth)

	// Stats
	e.GET("api/v1/stats/total_inventory_by_category", stshandlers.GetTotalInventoryByCategory)
	e.GET("api/v1/stats/total_inventory_by_favorite", stshandlers.GetTotalInventoryByFavorite)
	e.GET("api/v1/stats/total_inventory_by_room", stshandlers.GetTotalInventoryByRoom)
	e.GET("api/v1/stats/total_inventory_by_merk", stshandlers.GetTotalInventoryByMerk)
	e.GET("api/v1/stats/total_reminder_by_type", stshandlers.GetTotalReminderByType)
	e.GET("api/v1/stats/total_report_by_category", stshandlers.GetTotalReportByCategory)

	// Inventory
	e.GET("api/v1/inventory/list", invhandlers.GetListInventory)
	e.GET("api/v1/inventory/calendar", invhandlers.GetListCalendar)
	e.GET("api/v1/inventory/room", invhandlers.GetListRoom)
	e.GET("api/v1/inventory/merk", invhandlers.GetListMerk)
	e.GET("api/v1/inventory/search/by_room_storage/:room/:storage", invhandlers.GetInventoryByStorage)
	e.GET("api/v1/inventory/detail/:id", invhandlers.GetInventoryDetail)
	e.POST("api/v1/inventory", invhandlers.PostInventory)

	// Reminder
	e.POST("api/v1/reminder", invhandlers.PostReminder, middlewares.CustomJWTAuth)

	// History
	e.GET("api/v1/history", syshandlers.GetAllHistory)
	e.DELETE("api/v1/history/:id", syshandlers.HardDelHistoryById)

	return e
}
