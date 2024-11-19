package httphandlers

import (
	"gudangku/modules/inventories/repositories"
	"net/http"

	"github.com/labstack/echo"
)

func GetListInventory(c echo.Context) error {
	result, err := repositories.GetListInventoryRepo()
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"message": err.Error()})
	}

	return c.JSON(http.StatusOK, result)
}

func GetListCalendar(c echo.Context) error {
	result, err := repositories.GetListCalendarRepo()
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"message": err.Error()})
	}

	return c.JSON(http.StatusOK, result)
}

func GetListRoom(c echo.Context) error {
	result, err := repositories.GetListContextTotalRepo("inventory_room")
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"message": err.Error()})
	}

	return c.JSON(http.StatusOK, result)
}
