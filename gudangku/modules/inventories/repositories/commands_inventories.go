package repositories

import (
	"fmt"
	"gudangku/middlewares/firebase"
	"gudangku/modules/inventories/models"
	"gudangku/packages/builders"
	"gudangku/packages/database"
	"gudangku/packages/helpers/generator"
	"gudangku/packages/helpers/messager"
	"gudangku/packages/helpers/response"
	"log"
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
	MaxSizeFile:     5000000, // 5 MB
	AllowedFileType: []string{"jpg", "jpeg", "gif"},
}

func PostReminderRepo(d models.PostReminderModel, token string) (response.Response, error) {
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
		// Check email or username
		allow, err := builders.GetReminderAvailability(con, d.InventoryId, d.ReminderType, d.ReminderContext, userId)
		if err != nil {
			return res, err
		}
		if allow {
			// Check inventory
			inventory_name, err := builders.GetInventoryName(con, d.InventoryId, userId)
			if err != nil {
				return res, err
			}
			if inventory_name != "" {
				// Data
				id := uuid.Must(uuid.NewRandom())

				// Command builder
				sqlStatement = "INSERT INTO " + baseTable + " (id, inventory_id, reminder_desc, reminder_type, reminder_context, created_at, created_by, updated_at) " +
					"VALUES (?,?,?,?,?,?,?,null)"

				// Exec
				stmt, err := con.Prepare(sqlStatement)
				if err != nil {
					return res, err
				}

				result, err := stmt.Exec(id, d.InventoryId, d.ReminderDesc, d.ReminderType, d.ReminderContext, dt, userId)
				if err != nil {
					return res, err
				}

				rowsAffected, err := result.RowsAffected()
				if err != nil {
					return res, err
				}

				if rowsAffected > 0 {
					username, telegramUserId, isValid, err := builders.GetUserSocial(con, userId)
					if err != nil {
						return res, err
					}
					if isValid != false {
						if telegramUserId != 0 {
							msg := fmt.Sprintf("%s", "Hello, "+username+". You have create a reminder. Here's the reminder description for [DEMO]. "+d.ReminderDesc)
							messager.SendTelegramMessage(telegramUserId, msg)
						} else {
							log.Println("Telegram user ID not available for user:", userId)
						}
					}
				}

				// Response
				res.Status = http.StatusOK
				res.Message = generator.GenerateCommandMsg(baseTable, "create", int(rowsAffected))
				res.Data = map[string]interface{}{
					"id":            id,
					"data":          d,
					"rows_affected": rowsAffected,
				}
			} else {
				res.Status = http.StatusNotFound
				res.Message = "inventory not found"
				res.Data = nil
			}
		} else {
			res.Status = http.StatusConflict
			res.Message = "reminder with same type and context has been used"
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

func isValidFileType(fileExt string) bool {
	for _, ext := range config.AllowedFileType {
		if ext == fileExt {
			return true
		}
	}
	return false
}

func PostInventoryRepo(d models.InventoryDetailModel, token string, file *multipart.FileHeader, fileExt string, fileSize int64) (response.Response, error) {
	// Declaration
	var res response.Response
	var baseTable = "inventory"
	var sqlStatement string
	dt := time.Now().Format("2006-01-02 15:04:05")
	token = strings.Replace(token, "Bearer ", "", -1)
	con := database.CreateCon()

	userId, err := builders.GetUserIdFromToken(con, token)
	if err != nil {
		return res, err
	}

	if userId != "" {
		// Check inventory
		available, err := builders.GetInventoryAvailability(con, d.InventoryName, userId)
		if err != nil {
			return res, err
		}
		if available {
			// Data
			id := uuid.Must(uuid.NewRandom())
			username, telegramUserId, isValid, err := builders.GetUserSocial(con, userId)
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

			// Helper : Firebase Uplaod
			inventoryImage, err := firebase.UploadFile(baseTable, userId, username, file, fileExt)
			if err != nil {
				res.Status = http.StatusInternalServerError
				res.Message = "Failed to upload the file"
				res.Data = nil
				return res, nil
			}

			// Command builder
			colFirstTemplate := builders.GetTemplateSelect("inventory_list", nil, nil)
			colSecondTemplate := builders.GetTemplateSelect("inventory_placement", nil, nil)
			colPropsTemplate := builders.GetTemplateSelect("properties_full", nil, nil)
			sqlStatement = "INSERT INTO " + baseTable + " (" + colFirstTemplate + "," + colPropsTemplate + "," + colSecondTemplate + ",inventory_category, inventory_desc, inventory_merk, inventory_price, inventory_image, inventory_capacity_unit, inventory_capacity_vol, inventory_color, is_favorite, is_reminder, deleted_at) " +
				"VALUES (?,?,?,?,?,?,null,?,?,?,?,?,?,?,?,?,?,?,?,?,null)"

			// Exec
			stmt, err := con.Prepare(sqlStatement)
			if err != nil {
				return res, err
			}

			result, err := stmt.Exec(id, d.InventoryName, d.InventoryVol, d.InventoryUnit, dt, userId, d.InventoryRoom, d.InventoryStorage, d.InventoryRack, d.InventoryCategory, d.InventoryDesc, d.InventoryMerk, d.InventoryPrice, inventoryImage, d.InventoryCapacityUnit, d.InventoryCapacityVol, d.InventoryColor, d.IsFavorite, d.IsReminder)
			if err != nil {
				return res, err
			}

			rowsAffected, err := result.RowsAffected()
			if err != nil {
				return res, err
			}

			if rowsAffected > 0 {
				if isValid != false {
					if telegramUserId != 0 {
						msg := fmt.Sprintf("%s", "inventory created, its called "+d.InventoryName)
						messager.SendTelegramMessage(telegramUserId, msg)
					} else {
						log.Println("Telegram user ID not available for user:", userId)
					}
				}
			}

			// Response
			d.InventoryImage = &inventoryImage
			res.Status = http.StatusOK
			res.Message = generator.GenerateCommandMsg(baseTable, "create", int(rowsAffected))
			res.Data = map[string]interface{}{
				"id":            id,
				"data":          d,
				"rows_affected": rowsAffected,
			}
		} else {
			res.Status = http.StatusNotFound
			res.Message = "inventory is already exist"
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

func PutInventoryImageRepo(id, token string, file *multipart.FileHeader, fileExt string, fileSize int64) (response.Response, error) {
	// Declaration
	var res response.Response
	var baseTable = "inventory"
	var sqlStatement string
	dt := time.Now().Format("2006-01-02 15:04:05")
	token = strings.Replace(token, "Bearer ", "", -1)
	con := database.CreateCon()
	var newInventoryImage *string

	userId, err := builders.GetUserIdFromToken(con, token)
	if err != nil {
		return res, err
	}

	if userId != "" {
		// Get old inventory image
		_, oldInventoryImage, err := builders.GetInventoryImageById(con, id, userId)
		if err != nil {
			return res, err
		}
		if oldInventoryImage != nil {
			// Delete image
			err := firebase.DeleteFile(*oldInventoryImage)
			if err != nil {
				return res, err
			}
			newInventoryImage = nil
		}
		if file != nil {
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

			// Helper : Firebase Upload image
			inventoryImage, err := firebase.UploadFile(baseTable, userId, username, file, fileExt)
			if err != nil {
				return res, err
			}
			newInventoryImage = &inventoryImage
		} else {
			newInventoryImage = nil
		}

		// Command builder
		sqlStatement = "UPDATE " + baseTable + " SET inventory_image = ?, updated_at = ? WHERE id = ? AND created_by = ?"

		// Exec
		stmt, err := con.Prepare(sqlStatement)
		if err != nil {
			return res, err
		}

		result, err := stmt.Exec(newInventoryImage, dt, id, userId)
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
			res.Message = generator.GenerateCommandMsg(baseTable, "update", 1)
		} else {
			res.Status = http.StatusNotFound
			res.Message = generator.GenerateCommandMsg(baseTable, "update", 0)
		}
	} else {
		// Response
		res.Status = http.StatusUnprocessableEntity
		res.Message = "Valid token but user not found"
	}

	return res, nil
}
