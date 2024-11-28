package repositories

import (
	"gudangku/modules/inventories/models"
	"gudangku/packages/builders"
	"gudangku/packages/database"
	"gudangku/packages/helpers/generator"
	"gudangku/packages/helpers/response"
	"net/http"
	"strings"
	"time"
)

func DeleteReminderByIdRepo(token, id string) (response.Response, error) {
	// Declaration
	var res response.Response
	var baseTable = "reminder"
	var sqlStatement string
	token = strings.Replace(token, "Bearer ", "", -1)
	con := database.CreateCon()

	userId, err := builders.GetUserIdFromToken(con, token)
	if err != nil {
		return res, err
	}

	if userId != "" {
		// Command builder
		sqlStatement = `DELETE FROM ` + baseTable + ` WHERE id = ? AND created_by = ?`

		// Exec
		stmt, err := con.Prepare(sqlStatement)
		if err != nil {
			return res, err
		}
		result, err := stmt.Exec(id, userId)
		if err != nil {
			return res, err
		}
		rowsAffected, err := result.RowsAffected()
		if err != nil {
			return res, err
		}

		// Response
		res.Status = http.StatusOK
		res.Message = generator.GenerateCommandMsg(baseTable, "deleted", int(rowsAffected))
	} else {
		// Response
		res.Status = http.StatusUnprocessableEntity
		res.Message = "Valid token but user not found"
	}

	return res, nil
}

func PutReminderByIdRepo(d models.PostReminderModel, token, id string) (response.Response, error) {
	// Declaration
	var res response.Response
	var baseTable = "reminder"
	var sqlStatement string
	dt := time.Now().Format("2006-01-02 15:04:05")
	token = strings.Replace(token, "Bearer ", "", -1)
	con := database.CreateCon()

	userId, err := builders.GetUserIdFromToken(con, token)
	if err != nil {
		return res, err
	}

	if userId != "" {
		// Command builder
		sqlStatement = `UPDATE ` + baseTable + ` SET inventory_id = ?, reminder_type = ?, reminder_desc = ?, reminder_context = ?, updated_at = ? WHERE id = ?`

		// Exec
		stmt, err := con.Prepare(sqlStatement)
		if err != nil {
			return res, err
		}

		result, err := stmt.Exec(d.InventoryId, d.ReminderType, d.ReminderDesc, d.ReminderContext, dt, id)
		if err != nil {
			return res, err
		}

		rowsAffected, err := result.RowsAffected()
		if err != nil {
			return res, err
		}

		// Response
		res.Status = http.StatusOK
		res.Message = generator.GenerateCommandMsg(baseTable, "updated", int(rowsAffected))
		res.Data = map[string]interface{}{
			"id":            id,
			"data":          d,
			"rows_affected": rowsAffected,
		}
	} else {
		// Response
		res.Status = http.StatusUnprocessableEntity
		res.Message = "Valid token but user not found"
	}

	return res, nil
}
