package httphandlers

import (
	"gudangku/modules/inventories/models"
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

func GetInventoryDetail(c echo.Context) error {
	id := c.Param("id")
	result, err := repositories.GetInventoryDetailRepo(id)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"message": err.Error()})
	}

	return c.JSON(http.StatusOK, result)
}

func PostReminder(c echo.Context) error {
	var obj models.PostReminderModel
	token := c.Request().Header.Get("Authorization")

	obj.InventoryId = c.FormValue("inventory_id")
	obj.ReminderDesc = c.FormValue("reminder_desc")
	obj.ReminderContext = c.FormValue("reminder_context")
	obj.ReminderType = c.FormValue("reminder_type")

	result, err := repositories.PostReminderRepo(obj, token)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"message": err.Error()})
	}

	return c.JSON(http.StatusOK, result)
}
