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
