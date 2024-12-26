package repositories

import (
	"fmt"
	"gudangku/modules/systems/models"
	"gudangku/packages/builders"
	"gudangku/packages/database"
	"gudangku/packages/helpers/generator"
	"gudangku/packages/helpers/response"
	"net/http"
	"strings"
)

func PostDictionaryRepo(d models.PostDictionary, token string) (response.Response, error) {
	// Declaration
	var res response.Response
	var baseTable = "dictionary"
	var sqlStatement string
	token = strings.Replace(token, "Bearer ", "", -1)
	con := database.CreateCon()

	userId, err := builders.GetUserIdFromToken(con, token)
	if err != nil {
		return res, err
	}

	if userId != "" {
		// Check inventory
		available, err := builders.GetInventoryAvailability(con, d.DctName, d.DctType)
		if err != nil {
			return res, err
		}
		if available {
			// Command builder
			sqlStatement = "INSERT INTO " + baseTable + " (id,dictionary_type, dictionary_name) " +
				"VALUES (null,?,?)"

			// Exec
			stmt, err := con.Prepare(sqlStatement)
			if err != nil {
				return res, err
			}

			result, err := stmt.Exec(d.DctType, d.DctName)
			if err != nil {
				return res, err
			}

			rowsAffected, err := result.RowsAffected()
			if err != nil {
				return res, err
			}

			// Response
			res.Status = http.StatusOK
			res.Message = generator.GenerateCommandMsg(baseTable, "create", int(rowsAffected))
			lastInsertID, err := result.LastInsertId()
			if err != nil {
				return res, fmt.Errorf("failed to get LastInsertId: %v", err)
			}

			res.Data = map[string]interface{}{
				"id":            lastInsertID,
				"data":          d,
				"rows_affected": rowsAffected,
			}
		} else {
			res.Status = http.StatusNotFound
			res.Message = "dictionary is already exist"
			res.Data = nil
		}
	} else {
		// Response
		res.Status = http.StatusUnprocessableEntity
		res.Message = "Valid token but user not found"
		res.Data = nil
	}

	return res, nil
}

func DeleteDictionaryByIdRepo(token, id string) (response.Response, error) {
	// Declaration
	var res response.Response
	var baseTable = "dictionary"
	var sqlStatement string
	token = strings.Replace(token, "Bearer ", "", -1)
	con := database.CreateCon()

	userId, err := builders.GetUserIdFromToken(con, token)
	if err != nil {
		return res, err
	}

	if userId != "" {
		// Command builder
		sqlStatement = builders.GetTemplateCommand("hard_delete", baseTable, "id")

		// Exec
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
		res.Message = generator.GenerateCommandMsg(baseTable, "deleted", int(rowsAffected))
	} else {
		// Response
		res.Status = http.StatusUnprocessableEntity
		res.Message = "Valid token but user not found"
	}

	return res, nil
}
