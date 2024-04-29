package repositories

import (
	"gudangku/modules/stats/models"
	"gudangku/packages/builders"
	"gudangku/packages/database"
	"gudangku/packages/helpers/generator"
	"gudangku/packages/helpers/response"
	"net/http"
	"strconv"
)

func GetTotalStats(view string, table string, typeStats string, extraTotal *string) (response.Response, error) {
	// Declaration
	var obj models.GetMostAppear
	var arrobj []models.GetMostAppear
	var res response.Response
	var baseTable = table
	var mainCol = view
	var sqlStatement string
	var ord = "desc"

	// Converted column
	var totalStr string

	// Query builder
	sqlStatement = builders.GetTemplateStats(mainCol, baseTable, typeStats, ord, extraTotal)

	// Exec
	con := database.CreateCon()
	rows, err := con.Query(sqlStatement)
	defer rows.Close()

	if err != nil {
		return res, err
	}

	// Map
	for rows.Next() {
		err = rows.Scan(
			&obj.Context,
			&totalStr)

		if err != nil {
			return res, err
		}

		// Converted
		totalInt, err := strconv.Atoi(totalStr)
		if err != nil {
			return res, err
		}

		obj.Total = totalInt
		arrobj = append(arrobj, obj)
	}

	// Response
	res.Status = http.StatusOK
	res.Message = generator.GenerateQueryMsg("Stats", 1)
	res.Data = arrobj

	return res, nil
}
