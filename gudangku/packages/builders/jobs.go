package builders

import (
	"database/sql"
	"fmt"
	"strconv"
)

func GetTotalCount(con *sql.DB, table string, view *string) (int, error) {
	var count int
	var sqlStatement string

	// Fix this. if table empty, there will be an error
	if view != nil {
		sqlStatement = "SELECT COUNT(*) FROM " + table + " WHERE " + *view
	} else {
		sqlStatement = "SELECT COUNT(*) FROM " + table
	}

	err := con.QueryRow(sqlStatement).Scan(&count)
	if err != nil {
		return 0, err
	}
	return count, nil
}

func GetUserIdFromToken(con *sql.DB, token string) (string, error) {
	var id string

	sqlStatement := "SELECT users.id FROM users JOIN personal_access_tokens ON personal_access_tokens.tokenable_id = users.id WHERE token = ? LIMIT 1"
	err := con.QueryRow(sqlStatement, token).Scan(&id)
	if err != nil {
		if err == sql.ErrNoRows {
			return "", nil
		}
		return "", err
	}

	return id, nil
}

func GetUserRegisterAvailability(con *sql.DB, username, email string) (bool, error) {
	checkStatement := "SELECT 1 FROM users WHERE username = ? OR email = ? LIMIT 1"
	row := con.QueryRow(checkStatement, username, email)

	var dummy int
	err := row.Scan(&dummy)
	if err != nil {
		if err == sql.ErrNoRows {
			return true, nil
		}
		return false, err
	}

	return false, nil
}

func GetReminderAvailability(con *sql.DB, inventory_id, reminder_type, reminder_context, created_by string) (bool, error) {
	checkStatement := "SELECT 1 FROM reminder WHERE inventory_id = ? AND reminder_type = ? AND reminder_context = ? AND created_by = ? LIMIT 1"
	row := con.QueryRow(checkStatement, inventory_id, reminder_type, reminder_context, created_by)

	var dummy int
	err := row.Scan(&dummy)
	if err != nil {
		if err == sql.ErrNoRows {
			return true, nil
		}
		return false, err
	}

	return false, nil
}

func GetInventoryAvailability(con *sql.DB, inventory_name, created_by string) (bool, error) {
	checkStatement := "SELECT 1 FROM inventory WHERE inventory_name = ? AND created_by = ? LIMIT 1"
	row := con.QueryRow(checkStatement, inventory_name, created_by)

	var dummy int
	err := row.Scan(&dummy)
	if err != nil {
		if err == sql.ErrNoRows {
			return true, nil
		}
		return false, err
	}

	return false, nil
}

func GetInventoryName(con *sql.DB, inventory_id, user_id string) (string, error) {
	checkStatement := "SELECT inventory_name FROM inventory WHERE id = ? AND created_by = ? LIMIT 1"
	row := con.QueryRow(checkStatement, inventory_id, user_id)

	var inventoryName string
	err := row.Scan(&inventoryName)
	if err != nil {
		if err == sql.ErrNoRows {
			return "", nil
		}
		return "", err
	}

	return inventoryName, nil
}

func GetInventoryImageById(con *sql.DB, inventory_id, user_id string) (string, *string, error) {
	checkStatement := "SELECT inventory_name, inventory_image FROM inventory WHERE id = ? AND created_by = ? LIMIT 1"
	row := con.QueryRow(checkStatement, inventory_id, user_id)

	var inventoryName string
	var inventoryImage *string
	err := row.Scan(&inventoryName, &inventoryImage)
	if err != nil {
		if err == sql.ErrNoRows {
			return "", nil, nil
		}
		return "", nil, err
	}

	return inventoryName, inventoryImage, nil
}

func GetUserSocial(con *sql.DB, user_id string) (string, int64, bool, error) {
	checkStatement := "SELECT username, telegram_user_id, telegram_is_valid, email FROM users WHERE id = ? LIMIT 1"
	row := con.QueryRow(checkStatement, user_id)

	var telegramUserId *string
	var username *string
	var telegramIsValid int
	var email *string

	err := row.Scan(&username, &telegramUserId, &telegramIsValid, &email)
	if err != nil {
		if err == sql.ErrNoRows {
			return "", 0, false, nil
		}
		return "", 0, false, err
	}

	var telegramUserIdValue int64
	var isValid bool = (telegramIsValid == 1)

	if telegramUserId != nil {
		telegramUserIdValue, err = strconv.ParseInt(*telegramUserId, 10, 64)
		if err != nil {
			return "", 0, false, fmt.Errorf("failed to convert telegram_user_id to int64: %w", err)
		}
	}

	return *username, telegramUserIdValue, isValid, nil
}
