package httphandlers

import (
	"gudangku/modules/systems/repositories"
	"net/http"
	"strconv"

	"github.com/labstack/echo"
)

func GetDictionaryByType(c echo.Context) error {
	page, _ := strconv.Atoi(c.QueryParam("page"))
	dctType := c.Param("type")
	result, err := repositories.GetDictionaryByType(page, 10, "api/v1/dct/"+dctType, dctType)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"message": err.Error()})
	}

	return c.JSON(http.StatusOK, result)
}

func GetAllHistory(c echo.Context) error {
	page, _ := strconv.Atoi(c.QueryParam("page"))
	result, err := repositories.GetAllHistory(page, 10, "api/v1/history")
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"message": err.Error()})
	}

	return c.JSON(http.StatusOK, result)
}

func HardDelHistoryById(c echo.Context) error {
	id := c.Param("id")
	result, err := repositories.HardDelHistoryById(id)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"message": err.Error()})
	}

	return c.JSON(http.StatusOK, result)
}
