package httphandlers

import (
	"gudangku/modules/stats/repositories"
	"net/http"

	"github.com/labstack/echo"
)

func GetTotalInventoryByCategory(c echo.Context) error {
	view := "inventory_category"
	table := "inventory"

	result, err := repositories.GetTotalStats(view, table, "most_appear", nil)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"message": err.Error()})
	}

	return c.JSON(http.StatusOK, result)
}

func GetTotalInventoryByFavorite(c echo.Context) error {
	view := `
		CASE 
			WHEN is_favorite = 1 THEN 'Favorite' 
			ELSE 'Normal Item' 
		END AS context, 
		COUNT(1) as total
	`
	table := "inventory"
	colext := "is_favorite"

	result, err := repositories.GetTotalStats(view, table, "raw_select", &colext)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"message": err.Error()})
	}

	return c.JSON(http.StatusOK, result)
}

func GetTotalInventoryByRoom(c echo.Context) error {
	view := "inventory_room"
	table := "inventory"

	result, err := repositories.GetTotalStats(view, table, "most_appear", nil)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"message": err.Error()})
	}

	return c.JSON(http.StatusOK, result)
}

func GetTotalReminderByType(c echo.Context) error {
	view := "reminder_type"
	table := "reminder"

	result, err := repositories.GetTotalStats(view, table, "most_appear", nil)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"message": err.Error()})
	}

	return c.JSON(http.StatusOK, result)
}

func GetTotalReportByCategory(c echo.Context) error {
	view := "report_category"
	table := "report"

	result, err := repositories.GetTotalStats(view, table, "most_appear", nil)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"message": err.Error()})
	}

	return c.JSON(http.StatusOK, result)
}

func GetTotalInventoryByMerk(c echo.Context) error {
	view := "inventory_merk"
	table := "inventory"

	result, err := repositories.GetTotalStats(view, table, "most_appear", nil)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"message": err.Error()})
	}

	return c.JSON(http.StatusOK, result)
}
