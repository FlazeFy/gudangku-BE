package routes

import (
	middlewares "gudangku/middlewares/jwt"
	authhandlers "gudangku/modules/auth/http_handlers"
	invhandlers "gudangku/modules/inventories/http_handlers"
	rpthandlers "gudangku/modules/report/http_handlers"
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
	e.POST("api/v1/dct", syshandlers.PostDictionary, middlewares.CustomJWTAuth)
	e.DELETE("api/v1/dct/delete/:id", syshandlers.HardDeleteDictionaryById, middlewares.CustomJWTAuth)

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
	e.POST("api/v1/inventory", invhandlers.PostInventory, middlewares.CustomJWTAuth)
	e.PUT("api/v1/inventory/edit_image/:id", invhandlers.PutInventoryImageById, middlewares.CustomJWTAuth)
	e.PUT("api/v1/inventory/edit_layout/:id", invhandlers.PutInventoryLayoutById, middlewares.CustomJWTAuth)
	e.PUT("api/v1/inventory/recover/:id", invhandlers.PutRecoverInventoryById, middlewares.CustomJWTAuth)
	e.DELETE("api/v1/inventory/delete/:id", invhandlers.SoftDeleteInventoryById, middlewares.CustomJWTAuth)
	e.DELETE("api/v1/inventory/destroy/:id", invhandlers.HardDeleteInventoryById, middlewares.CustomJWTAuth)

	// Reminder
	e.POST("api/v1/reminder", invhandlers.PostReminder, middlewares.CustomJWTAuth)
	e.PUT("api/v1/reminder/:id", invhandlers.UpdateReminderById, middlewares.CustomJWTAuth)
	e.DELETE("api/v1/reminder/:id", invhandlers.DeleteReminderById, middlewares.CustomJWTAuth)

	// Report
	e.POST("api/v1/report", rpthandlers.PostReport, middlewares.CustomJWTAuth)
	e.DELETE("api/v1/report/delete/item/:id", rpthandlers.DeleteReportItemById, middlewares.CustomJWTAuth)
	e.DELETE("api/v1/report/delete/report/:id", rpthandlers.DeleteReportById, middlewares.CustomJWTAuth)
	e.PUT("api/v1/report/update/item/:id", rpthandlers.UpdateReportItemById, middlewares.CustomJWTAuth)
	e.PUT("api/v1/report/update/report/:id", rpthandlers.UpdateReportById, middlewares.CustomJWTAuth)

	// History
	e.GET("api/v1/history", syshandlers.GetAllHistory)
	e.DELETE("api/v1/history/:id", syshandlers.HardDelHistoryById)

	return e
}
