package httphandlers

import (
	"gudangku/modules/systems/models"
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

func PostDictionary(c echo.Context) error {
	var obj models.PostDictionary
	token := c.Request().Header.Get("Authorization")

	obj.DctName = c.FormValue("dictionary_name")
	obj.DctType = c.FormValue("dictionary_type")

	result, err := repositories.PostDictionaryRepo(obj, token)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"message": err.Error()})
	}

	return c.JSON(http.StatusOK, result)
}

func HardDeleteDictionaryById(c echo.Context) error {
	token := c.Request().Header.Get("Authorization")
	id := c.Param("id")

	result, err := repositories.DeleteDictionaryByIdRepo(token, id)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"message": err.Error()})
	}

	return c.JSON(http.StatusOK, result)
}
