package repositories

import (
	"database/sql"
	"gudangku/modules/auth/models"
	"gudangku/packages/builders"
	"gudangku/packages/database"
	"gudangku/packages/helpers/converter"
	"gudangku/packages/helpers/generator"
	"gudangku/packages/helpers/response"
	"net/http"
	"strconv"
)

func GetMyProfileRepo(token string) (response.Response, error) {
	// Declaration
	var obj models.UserProfileModel
	var res response.Response
	var baseTable = "users"
	var sqlStatement string

	// Nullable Column
	var TelegramUserId, LineUserId, FirebaseFCMToken, Phone, Timezone, UpdatedAt sql.NullString

	// Converted Column
	var TelegramIsValid string

	colTemplate := builders.GetTemplateSelect("social", &baseTable, nil)
	sqlStatement = "SELECT username, " + colTemplate + ", users.created_at, users.updated_at " +
		"FROM " + baseTable + " " +
		"JOIN personal_access_tokens ON personal_access_tokens.tokenable_id = users.id " +
		"WHERE token = ? LIMIT 1"

	// Exec
	con := database.CreateCon()
	row := con.QueryRow(sqlStatement, token)

	err := row.Scan(
		&obj.Username,
		&obj.Email,
		&TelegramUserId,
		&TelegramIsValid,
		&FirebaseFCMToken,
		&LineUserId,
		&Phone,
		&Timezone,
		&obj.CreatedAt,
		&UpdatedAt,
	)

	if err != nil {
		if err == sql.ErrNoRows {
			res.Status = http.StatusNotFound
			res.Message = "User not found"
			res.Data = nil
			return res, nil
		}
		return res, err
	}

	intTelegramIsValid, err := strconv.Atoi(TelegramIsValid)
	if err != nil {
		intTelegramIsValid = 0
	}
	obj.TelegramIsValid = intTelegramIsValid

	obj.TelegramUserId = converter.CheckNullString(TelegramUserId)
	obj.FirebaseFCMToken = converter.CheckNullString(FirebaseFCMToken)
	obj.LineUserId = converter.CheckNullString(LineUserId)
	obj.Phone = converter.CheckNullString(Phone)
	obj.Timezone = converter.CheckNullString(Timezone)
	obj.UpdatedAt = converter.CheckNullString(UpdatedAt)

	// Success response
	res.Status = http.StatusOK
	res.Message = generator.GenerateQueryMsg(baseTable, 1)
	res.Data = obj

	return res, nil
}
