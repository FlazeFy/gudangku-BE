package converter

import (
	"database/sql"
	"encoding/json"
	"strings"
)

func CheckNullString(data sql.NullString) string {
	var res string
	if data.Valid {
		res = data.String
	} else {
		res = ""
	}

	return res
}

func TotalChar(val string) int {
	trimed := strings.TrimSpace(val)
	return len(trimed)
}

func ConvertStringBool(val string) bool {
	if val == "0" {
		return false
	} else {
		return true
	}
}

func MapToString(val map[string]string) string {
	result, _ := json.Marshal(val)
	return string(result)
}

func StringToMap(val string) (map[string]string, error) {
	var result map[string]string
	err := json.Unmarshal([]byte(val), &result)
	return result, err
}
