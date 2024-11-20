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

func GetListMerk(c echo.Context) error {
	result, err := repositories.GetListContextTotalRepo("inventory_merk")
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"message": err.Error()})
	}

	return c.JSON(http.StatusOK, result)
}

func GetInventoryByStorage(c echo.Context) error {
	room := c.Param("room")
	storage := c.Param("storage")
	result, err := repositories.GetInventoryByStorageRepo(room, storage)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"message": err.Error()})
	}

	return c.JSON(http.StatusOK, result)
}
