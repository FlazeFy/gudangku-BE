package repositories

import (
	"fmt"
	"gudangku/modules/inventories/models"
	"gudangku/packages/builders"
	"gudangku/packages/database"
	"gudangku/packages/helpers/generator"
	"gudangku/packages/helpers/messager"
	"gudangku/packages/helpers/response"
	"log"
	"net/http"
	"strings"
	"time"

	"github.com/google/uuid"
)

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
		allow, err := builders.GetReminderAvailability(con, d.InventoryId, d.ReminderType, d.ReminderContext)
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
					telegramUserId, isValid, err := builders.GetUserSocial(con, userId)
					if err != nil {
						return res, err
					}
					if isValid != false {
						if telegramUserId != 0 {
							msg := fmt.Sprintf("You have create a reminder. Here's the reminder description for [DEMO]. " + d.ReminderDesc)
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
