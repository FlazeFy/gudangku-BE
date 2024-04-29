package repositories

import (
	"gudangku/modules/systems/models"
	"gudangku/packages/database"
	"gudangku/packages/helpers/generator"
	"gudangku/packages/helpers/response"
	"gudangku/packages/utils/pagination"
	"math"
	"net/http"
)

func GetAllHistory(page, pageSize int, path string) (response.Response, error) {
	// Declaration
	var obj models.GetHistory
	var arrobj []models.GetHistory
	var res response.Response
	var baseTable = "history"
	var sqlStatement string

	sqlStatement = "SELECT id, history_type, history_context, created_at  " +
		"FROM " + baseTable + " " +
		"ORDER BY created_at DESC " +
		"LIMIT ? OFFSET ?"

	// Exec
	con := database.CreateCon()
	offset := (page - 1) * pageSize
	rows, err := con.Query(sqlStatement, pageSize, offset)
	defer rows.Close()

	if err != nil {
		return res, err
	}

	// Map
	for rows.Next() {
		err = rows.Scan(
			&obj.ID,
			&obj.HistoryType,
			&obj.HistoryContext,
			&obj.CreatedAt,
		)

		if err != nil {
			return res, err
		}

		arrobj = append(arrobj, obj)
	}

	// Page
	total := len(arrobj)

	totalPages := int(math.Ceil(float64(total) / float64(pageSize)))
	pagination := pagination.BuildPaginationResponse(page, pageSize, total, totalPages, path)

	// Response
	res.Status = http.StatusOK
	res.Message = generator.GenerateQueryMsg(baseTable, total)
	if total == 0 {
		res.Data = nil
	} else {
		res.Data = map[string]interface{}{
			"current_page":   page,
			"data":           arrobj,
			"first_page_url": pagination.FirstPageURL,
			"from":           pagination.From,
			"last_page":      pagination.LastPage,
			"last_page_url":  pagination.LastPageURL,
			"links":          pagination.Links,
			"next_page_url":  pagination.NextPageURL,
			"path":           pagination.Path,
			"per_page":       pageSize,
			"prev_page_url":  pagination.PrevPageURL,
			"to":             pagination.To,
			"total":          total,
		}
	}

	return res, nil
}
