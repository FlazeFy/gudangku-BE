package repositories

import (
	"gudangku/modules/inventories/models"
	stats_model "gudangku/modules/stats/models"
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
	var InventoryPrice string

	colFirstTemplate := builders.GetTemplateSelect("inventory_calendar", nil, nil)
	sqlStatement = "SELECT id, " + colFirstTemplate + " " +
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
			&InventoryPrice,
			&obj.CreatedAt,
		)

		if err != nil {
			return res, err
		}

		// Converted
		intInventoryPrice, err := strconv.Atoi(InventoryPrice)
		if err != nil {
			return res, err
		}

		obj.InventoryPrice = intInventoryPrice

		arrobj = append(arrobj, obj)
	}

	res.Status = http.StatusOK
	res.Message = generator.GenerateQueryMsg(baseTable, len(arrobj))
	res.Data = arrobj

	return res, nil
}

func GetListContextTotalRepo(ctx string) (response.Response, error) {
	// Declaration
	var obj models.GetListContextModel
	var arrobj []models.GetListContextModel
	var res response.Response
	var baseTable = "inventory"
	var sqlStatement string

	sqlStatement = "SELECT " + ctx + " as context " +
		"FROM " + baseTable + " " +
		"WHERE deleted_at is null " +
		"GROUP BY " + ctx + " " +
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
			&obj.Context,
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

func GetInventoryByStorageRepo(room, storage string) (response.ResponseWithStats, error) {
	// Declaration
	var obj models.GetInventoryByStorageModel
	var arrobj []models.GetInventoryByStorageModel
	var res response.ResponseWithStats
	var baseTable = "inventory"
	var sqlStatement string

	// Converted column
	var InventoryVol string
	var InventoryPrice string

	// Query Data
	colFirstTemplate := builders.GetTemplateSelect("inventory_list", nil, nil)
	sqlStatement = "SELECT " + colFirstTemplate + ", inventory_category, inventory_price " +
		"FROM " + baseTable + " " +
		"WHERE inventory_storage = '" + storage + "' " +
		"AND inventory_room = '" + room + "' " +
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
			&obj.InventoryCategory,
			&InventoryPrice,
		)

		if err != nil {
			return res, err
		}

		// Converted
		intInventoryVol, err := strconv.Atoi(InventoryVol)
		intInventoryPrice, err := strconv.Atoi(InventoryPrice)
		if err != nil {
			return res, err
		}
		obj.InventoryVol = intInventoryVol
		obj.InventoryPrice = intInventoryPrice

		arrobj = append(arrobj, obj)
	}

	if len(arrobj) > 0 {
		var sqlStatementStats string
		var objStats stats_model.GetMostAppear
		var arrobjStats []stats_model.GetMostAppear

		// Converted column
		var Total string

		sqlStatementStats = "SELECT inventory_category as context, COUNT(1) as total " +
			"FROM " + baseTable + " " +
			"WHERE inventory_storage = '" + storage + "' " +
			"AND inventory_room = '" + room + "' " +
			"GROUP BY inventory_category "

		// Exec
		con := database.CreateCon()
		rowsStats, err := con.Query(sqlStatementStats)
		defer rows.Close()

		if err != nil {
			return res, err
		}

		// Map
		for rowsStats.Next() {
			err = rowsStats.Scan(
				&objStats.Context,
				&Total,
			)

			if err != nil {
				return res, err
			}

			// Converted
			intTotal, err := strconv.Atoi(Total)
			if err != nil {
				return res, err
			}
			objStats.Total = intTotal

			arrobjStats = append(arrobjStats, objStats)
		}

		res.Status = http.StatusOK
		res.Message = generator.GenerateQueryMsg(baseTable, len(arrobj))
		res.Data = arrobj
		res.Stats = arrobjStats
	} else {
		res.Status = http.StatusNotFound
		res.Message = generator.GenerateQueryMsg(baseTable, 0)
	}

	return res, nil
}
