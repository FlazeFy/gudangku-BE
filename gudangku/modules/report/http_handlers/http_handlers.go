package httphandlers

import (
	"gudangku/modules/report/models"
	"gudangku/modules/report/repositories"
	"gudangku/packages/helpers/converter"
	"net/http"
	"path/filepath"
	"strconv"
	"strings"

	"github.com/labstack/echo"
)

func PostReport(c echo.Context) error {
	var obj models.PostReportModel
	token := c.Request().Header.Get("Authorization")

	obj.ReportTitle = c.FormValue("report_title")
	obj.ReportCategory = c.FormValue("report_category")
	obj.ReportDesc = converter.NullableString(c.FormValue("report_desc"))
	obj.ReportItem = converter.NullableString(c.FormValue("report_item"))
	obj.IsReminder, _ = strconv.Atoi(c.FormValue("is_reminder"))
	obj.RemindAt = converter.NullableString(c.FormValue("remind_at"))

	file, err := c.FormFile("file")
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"message": "Failed to retrieve the file"})
	}
	fileExt := strings.ToLower(strings.TrimPrefix(filepath.Ext(file.Filename), "."))
	fileSize := file.Size

	result, err := repositories.PostReportRepo(obj, token, file, fileExt, fileSize)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"message": err.Error()})
	}

	return c.JSON(http.StatusOK, result)
}

func DeleteReportById(c echo.Context) error {
	token := c.Request().Header.Get("Authorization")
	id := c.Param("id")

	result, err := repositories.DeleteReportByIdRepo(token, id)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"message": err.Error()})
	}

	return c.JSON(http.StatusOK, result)
}

func DeleteReportItemById(c echo.Context) error {
	token := c.Request().Header.Get("Authorization")
	id := c.Param("id")

	result, err := repositories.DeleteReportItemByIdRepo(token, id)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"message": err.Error()})
	}

	return c.JSON(http.StatusOK, result)
}

func UpdateReportById(c echo.Context) error {
	var obj models.PostReportModel
	token := c.Request().Header.Get("Authorization")

	id := c.Param("id")
	obj.ReportTitle = c.FormValue("report_title")
	obj.ReportCategory = c.FormValue("report_category")
	obj.ReportDesc = converter.NullableString(c.FormValue("report_desc"))
	obj.IsReminder, _ = strconv.Atoi(c.FormValue("is_reminder"))
	obj.RemindAt = converter.NullableString(c.FormValue("remind_at"))

	result, err := repositories.PutReportByIdRepo(obj, token, id)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"message": err.Error()})
	}

	return c.JSON(http.StatusOK, result)
}

func UpdateReportItemById(c echo.Context) error {
	var obj models.ReportItemModel
	token := c.Request().Header.Get("Authorization")

	id := c.Param("id")
	obj.ItemName = c.FormValue("item_name")
	obj.ItemDesc = converter.NullableString(c.FormValue("item_desc"))
	obj.ItemQty, _ = strconv.Atoi(c.FormValue("item_qty"))
	itemPriceStr := c.FormValue("item_price")
	if itemPriceStr == "" {
		obj.ItemPrice = nil
	} else {
		itemPrice, err := strconv.Atoi(itemPriceStr)
		if err != nil {
			return c.JSON(http.StatusBadRequest, map[string]string{"message": "Invalid item_price value"})
		}
		obj.ItemPrice = &itemPrice
	}

	result, err := repositories.PutReportItemByIdRepo(obj, token, id)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"message": err.Error()})
	}

	return c.JSON(http.StatusOK, result)
}
