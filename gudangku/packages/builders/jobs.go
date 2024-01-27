package builders

import "database/sql"

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
