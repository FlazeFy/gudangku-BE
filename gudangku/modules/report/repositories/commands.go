package repositories

import (
	"encoding/json"
	"fmt"
	"gudangku/middlewares/firebase"
	"gudangku/modules/report/models"
	"gudangku/packages/builders"
	"gudangku/packages/database"
	"gudangku/packages/helpers/generator"
	"gudangku/packages/helpers/response"
	"mime/multipart"
	"net/http"
	"strings"
	"time"

	"github.com/google/uuid"
)

type Config struct {
	MaxSizeFile     int64
	AllowedFileType []string
}

var config = Config{
	MaxSizeFile:     10000000, // 10 MB
	AllowedFileType: []string{"jpg", "jpeg", "gif", "png"},
}

func isValidFileType(fileExt string) bool {
	for _, ext := range config.AllowedFileType {
		if ext == fileExt {
			return true
		}
	}
	return false
}

func PostReportRepo(d models.PostReportModel, token string, file *multipart.FileHeader, fileExt string, fileSize int64) (response.Response, error) {
	// Declaration
	var res response.Response
	var baseTable = "report"
	var itemTable = "report_item"
	var sqlStatement string
	dt := time.Now().Format("2006-01-02 15:04:05")
	token = strings.Replace(token, "Bearer ", "", -1)
	con := database.CreateCon()

	userId, err := builders.GetUserIdFromToken(con, token)
	if err != nil {
		return res, err
	}

	if userId != "" {
		// Data
		id := uuid.Must(uuid.NewRandom())
		username, _, _, err := builders.GetUserSocial(con, userId)
		if err != nil {
			return res, err
		}

		// Validate file type
		if !isValidFileType(fileExt) {
			res.Status = http.StatusInternalServerError
			res.Message = fmt.Sprintf("The file must be a %s file type.", strings.Join(config.AllowedFileType, ", "))
			res.Data = nil
			return res, nil
		}

		// Validate file size
		if fileSize > config.MaxSizeFile {
			res.Status = http.StatusInternalServerError
			res.Message = fmt.Sprintf("The file size must be under %.2f MB", float64(config.MaxSizeFile)/1000000)
			res.Data = nil
			return res, nil
		}

		// Check open file
		fileReader, err := file.Open()
		if err != nil {
			res.Status = http.StatusInternalServerError
			res.Message = "Failed to open the file"
			res.Data = nil
			return res, nil
		}
		defer fileReader.Close()

		// Helper : Firebase Upload
		reportImage, err := firebase.UploadFile(baseTable, userId, username, file, fileExt)
		if err != nil {
			res.Status = http.StatusInternalServerError
			res.Message = "Failed to upload the file"
			res.Data = nil
			return res, nil
		}

		// Command builder
		colFirstTemplate := builders.GetTemplateSelect("report_list", nil, nil)
		colPropsTemplate := builders.GetTemplateSelect("properties_full", nil, nil)
		sqlStatement = `INSERT INTO ` + baseTable + ` (` + colFirstTemplate + `,` + colPropsTemplate + `,report_desc, is_reminder, remind_at, deleted_at)
			VALUES (?,?,?,?,?,?,null,?,?,?,null)`

		// Exec
		stmt, err := con.Prepare(sqlStatement)
		if err != nil {
			return res, err
		}

		result, err := stmt.Exec(id, d.ReportTitle, d.ReportCategory, reportImage, dt, userId, d.ReportDesc, d.IsReminder, d.RemindAt)
		if err != nil {
			return res, err
		}

		rowsAffected, err := result.RowsAffected()
		if err != nil {
			return res, err
		}

		// Insert report items
		var reportItems []models.ReportItemModel
		if err := json.Unmarshal([]byte(*d.ReportItem), &reportItems); err != nil {
			return res, err
		}

		colItemsTemplate := builders.GetTemplateSelect("report_item", nil, nil)
		for _, item := range reportItems {
			itemID := uuid.Must(uuid.NewRandom())
			itemSQL := `INSERT INTO ` + itemTable + ` (id, report_id, inventory_id, ` + colItemsTemplate + `, created_at, created_by)
				VALUES (?, ?, ?, ?, ?, ?, ?, ?, ?)`
			itemStmt, err := con.Prepare(itemSQL)
			if err != nil {
				return res, err
			}
			_, err = itemStmt.Exec(itemID, id, item.InventoryID, item.ItemName, item.ItemDesc, item.ItemQty, item.ItemPrice, dt, userId)
			if err != nil {
				return res, err
			}
		}

		// Response
		d.ReportImage = &reportImage
		res.Status = http.StatusOK
		res.Message = "Report and items added successfully"
		res.Data = map[string]interface{}{
			"id": id,
			"data": map[string]interface{}{
				"report_title":    d.ReportTitle,
				"report_category": d.ReportCategory,
				"report_desc":     d.ReportDesc,
				"report_image":    d.ReportImage,
				"is_reminder":     d.IsReminder,
				"remind_at":       d.RemindAt,
			},
			"items":         reportItems,
			"rows_affected": rowsAffected,
		}
	} else {
		// Response
		res.Status = http.StatusUnprocessableEntity
		res.Message = "Valid token but user not found"
		res.Data = nil
	}

	return res, nil
}

func DeleteReportByIdRepo(token, id string) (response.Response, error) {
	// Declaration
	var res response.Response
	var baseTable = "report"
	var itemTable = "report_item"
	var sqlStatement, sqlItemStatement string
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
		if rowsAffected > 0 {
			sqlItemStatement = `DELETE FROM ` + itemTable + ` WHERE report_id = ? AND created_by = ?`
			stmt, err := con.Prepare(sqlItemStatement)
			if err != nil {
				return res, err
			}
			_, err = stmt.Exec(id, userId)
			if err != nil {
				return res, err
			}
			res.Status = http.StatusOK
			res.Message = "Report deleted"
		} else {
			res.Status = http.StatusNotFound
			res.Message = "Report not found"
		}
	} else {
		// Response
		res.Status = http.StatusUnprocessableEntity
		res.Message = "Valid token but user not found"
		res.Data = nil
	}

	return res, nil
}

func DeleteReportItemByIdRepo(token, id string) (response.Response, error) {
	// Declaration
	var res response.Response
	var baseTable = "report_item"
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
		if rowsAffected > 0 {
			res.Status = http.StatusOK
			res.Message = "Report item removed"
		} else {
			res.Status = http.StatusNotFound
			res.Message = "Report item not found"
		}
	} else {
		// Response
		res.Status = http.StatusUnprocessableEntity
		res.Message = "Valid token but user not found"
		res.Data = nil
	}

	return res, nil
}

func PutReportByIdRepo(d models.PostReportModel, token, id string) (response.Response, error) {
	// Declaration
	var res response.Response
	var baseTable = "report"
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
		sqlStatement = `UPDATE ` + baseTable + ` SET report_title = ?, report_desc = ?, report_category = ?, is_reminder = ?, remind_at = ?, updated_at = ? WHERE id = ?`

		// Exec
		stmt, err := con.Prepare(sqlStatement)
		if err != nil {
			return res, err
		}

		result, err := stmt.Exec(d.ReportTitle, d.ReportDesc, d.ReportCategory, d.IsReminder, d.RemindAt, dt, id)
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
			"id": id,
			"data": map[string]interface{}{
				"report_title":    d.ReportTitle,
				"report_category": d.ReportCategory,
				"report_desc":     d.ReportDesc,
				"is_reminder":     d.IsReminder,
				"remind_at":       d.RemindAt,
			},
			"rows_affected": rowsAffected,
		}
	} else {
		// Response
		res.Status = http.StatusUnprocessableEntity
		res.Message = "Valid token but user not found"
		res.Data = nil
	}

	return res, nil
}

func PutReportItemByIdRepo(d models.ReportItemModel, token, id string) (response.Response, error) {
	// Declaration
	var res response.Response
	var baseTable = "report_item"
	var sqlStatement string
	token = strings.Replace(token, "Bearer ", "", -1)
	con := database.CreateCon()

	userId, err := builders.GetUserIdFromToken(con, token)
	if err != nil {
		return res, err
	}

	if userId != "" {
		// Command builder
		sqlStatement = `UPDATE ` + baseTable + ` SET item_name = ?, item_desc = ?, item_qty = ?, item_price = ? WHERE id = ?`

		// Exec
		stmt, err := con.Prepare(sqlStatement)
		if err != nil {
			return res, err
		}

		result, err := stmt.Exec(d.ItemName, d.ItemDesc, d.ItemQty, d.ItemPrice, id)
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
			"id": id,
			"data": map[string]interface{}{
				"item_name":  d.ItemName,
				"item_desc":  d.ItemDesc,
				"item_qty":   d.ItemQty,
				"item_price": d.ItemPrice,
			},
			"rows_affected": rowsAffected,
		}
	} else {
		// Response
		res.Status = http.StatusUnprocessableEntity
		res.Message = "Valid token but user not found"
		res.Data = nil
	}

	return res, nil
}
