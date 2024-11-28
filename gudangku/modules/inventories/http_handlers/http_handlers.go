package httphandlers

import (
	"gudangku/modules/inventories/models"
	"gudangku/modules/inventories/repositories"
	"gudangku/packages/helpers/converter"
	"mime/multipart"
	"net/http"
	"path/filepath"
	"strconv"
	"strings"

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

func PostInventory(c echo.Context) error {
	var obj models.InventoryDetailModel
	token := c.Request().Header.Get("Authorization")

	obj.InventoryName = c.FormValue("inventory_name")
	obj.InventoryCategory = converter.NullableString(c.FormValue("inventory_category"))
	obj.InventoryDesc = converter.NullableString(c.FormValue("inventory_desc"))
	obj.InventoryMerk = converter.NullableString(c.FormValue("inventory_merk"))
	obj.InventoryRoom = converter.NullableString(c.FormValue("inventory_room"))
	obj.InventoryStorage = converter.NullableString(c.FormValue("inventory_storage"))
	obj.InventoryRack = converter.NullableString(c.FormValue("inventory_rack"))
	obj.InventoryPrice, _ = strconv.Atoi(c.FormValue("inventory_price"))
	obj.InventoryImage = converter.NullableString(c.FormValue("inventory_image"))
	obj.InventoryUnit = c.FormValue("inventory_unit")
	obj.InventoryVol, _ = strconv.Atoi(c.FormValue("inventory_vol"))
	obj.InventoryCapacityUnit = converter.NullableString(c.FormValue("inventory_capacity_unit"))
	obj.InventoryCapacityVol, _ = strconv.Atoi(c.FormValue("inventory_capacity_vol"))
	obj.InventoryColor = converter.NullableString(c.FormValue("inventory_color"))
	obj.IsFavorite, _ = strconv.Atoi(c.FormValue("is_favorite"))
	obj.IsReminder, _ = strconv.Atoi(c.FormValue("is_reminder"))

	file, err := c.FormFile("file")
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"message": "Failed to retrieve the file"})
	}
	fileExt := strings.ToLower(strings.TrimPrefix(filepath.Ext(file.Filename), "."))
	fileSize := file.Size

	result, err := repositories.PostInventoryRepo(obj, token, file, fileExt, fileSize)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"message": err.Error()})
	}

	return c.JSON(http.StatusOK, result)
}

func PutInventoryImageById(c echo.Context) error {
	token := c.Request().Header.Get("Authorization")
	id := c.Param("id")

	var fileExt string
	var fileSize int64
	var file *multipart.FileHeader

	file, _ = c.FormFile("file")
	if file != nil {
		fileExt = strings.ToLower(strings.TrimPrefix(filepath.Ext(file.Filename), "."))
		fileSize = file.Size
	}

	result, err := repositories.PutInventoryImageRepo(id, token, file, fileExt, fileSize)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"message": err.Error()})
	}

	return c.JSON(http.StatusOK, result)
}

func DeleteReminderById(c echo.Context) error {
	token := c.Request().Header.Get("Authorization")
	id := c.Param("id")

	result, err := repositories.DeleteReminderByIdRepo(token, id)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"message": err.Error()})
	}

	return c.JSON(http.StatusOK, result)
}

func UpdateReminderById(c echo.Context) error {
	var obj models.PostReminderModel
	token := c.Request().Header.Get("Authorization")

	id := c.Param("id")
	obj.InventoryId = c.FormValue("inventory_id")
	obj.ReminderType = c.FormValue("reminder_type")
	obj.ReminderDesc = c.FormValue("reminder_desc")
	obj.ReminderContext = c.FormValue("reminder_context")

	result, err := repositories.PutReminderByIdRepo(obj, token, id)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"message": err.Error()})
	}

	return c.JSON(http.StatusOK, result)
}
