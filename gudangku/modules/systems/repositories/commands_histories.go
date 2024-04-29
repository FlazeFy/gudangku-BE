package repositories

import (
	"gudangku/packages/builders"
	"gudangku/packages/database"
	"gudangku/packages/helpers/generator"
	"gudangku/packages/helpers/response"
	"net/http"
)

func HardDelHistoryById(id string) (response.Response, error) {
	// Declaration
	var res response.Response
	var baseTable = "history"
	var sqlStatement string

	// Command builder
	sqlStatement = builders.GetTemplateCommand("hard_delete", baseTable, "id")

	// Exec
	con := database.CreateCon()
	stmt, err := con.Prepare(sqlStatement)
	if err != nil {
		return res, err
	}

	result, err := stmt.Exec(id)
	if err != nil {
		return res, err
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return res, err
	}

	// Response
	res.Status = http.StatusOK
	res.Message = generator.GenerateCommandMsg(baseTable, "permanently delete", int(rowsAffected))
	res.Data = map[string]int64{
		"rows_affected": rowsAffected,
	}

	return res, nil
}
