package repositories

import (
	"fmt"
	"gudangku/modules/inventories/models"
	"gudangku/packages/builders"
	"gudangku/packages/database"
	"gudangku/packages/helpers/generator"
	"gudangku/packages/helpers/response"
	"net/http"
	"strconv"
)

func GetListInventoryRepo() (response.Response, error) {
	// Declaration
	var obj models.GetListInventoryModel
	var arrobj []models.GetListInventoryModel
	var res response.Response
	var baseTable = "inventory"
	var sqlStatement string

	// Converted Column
	var InventoryVol string

	colFirstTemplate := builders.GetTemplateSelect("inventory_list", nil, nil)
	sqlStatement = "SELECT " + colFirstTemplate + " " +
		"FROM " + baseTable + " " +
		"WHERE deleted_at is null " +
		"ORDER BY inventory_name ASC "

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
			&obj.ID,
			&obj.InventoryName,
			&InventoryVol,
			&obj.InventoryUnit,
		)

		if err != nil {
			return res, err
		}

		// Converted
		intInventoryVol, err := strconv.Atoi(InventoryVol)
		if err != nil {
			return res, err
		}

		obj.InventoryVol = intInventoryVol

		arrobj = append(arrobj, obj)
	}

	res.Status = http.StatusOK
	res.Message = generator.GenerateQueryMsg(baseTable, len(arrobj))
	res.Data = arrobj

	return res, nil
}

func GetListCalendarRepo() (response.Response, error) {
	// Declaration
	var obj models.GetListCalendarModel
	var arrobj []models.GetListCalendarModel
	var res response.Response
	var baseTable = "inventory"
	var sqlStatement string

	// Converted Column
	var InventoryVol string

	colFirstTemplate := builders.GetTemplateSelect("inventory_calendar", nil, nil)
	sqlStatement = "SELECT " + colFirstTemplate + " " +
		"FROM " + baseTable + " " +
		"WHERE deleted_at is null " +
		"ORDER BY created_at DESC "

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
			&obj.ID,
			&obj.InventoryName,
			&InventoryVol,
			&obj.InventoryUnit,
		)

		if err != nil {
			return res, err
		}

		// Converted
		intInventoryVol, err := strconv.Atoi(InventoryVol)
		if err != nil {
			return res, err
		}

		obj.InventoryVol = intInventoryVol

		arrobj = append(arrobj, obj)
	}

	res.Status = http.StatusOK
	res.Message = generator.GenerateQueryMsg(baseTable, len(arrobj))
	res.Data = arrobj

	return res, nil
}

func GetListContextTotalRepo(ctx string) (response.Response, error) {
	// Declaration
	var obj models.GetListRoomModel
	var arrobj []models.GetListRoomModel
	var res response.Response
	var baseTable = "inventory"
	var sqlStatement string

	sqlStatement = "SELECT " + ctx + " " +
		"FROM " + baseTable + " " +
		"WHERE deleted_at is null " +
		"ORDER BY created_at DESC " +
		"GROUP BY " + ctx

	fmt.Println(sqlStatement)

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
			&obj.InventoryRoom,
		)

		if err != nil {
			return res, err
		}

		arrobj = append(arrobj, obj)
	}

	res.Status = http.StatusOK
	res.Message = generator.GenerateQueryMsg(baseTable, len(arrobj))
	res.Data = arrobj

	return res, nil
}
