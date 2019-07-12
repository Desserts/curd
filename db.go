package curd

import (
	"strings"

	"github.com/jmoiron/sqlx"
)

type DB struct {
	sqlxDB *sqlx.DB
}

// MakeSqlParams 组装 sql 参数
func MakeInsertParams(data map[string]interface{}) (string, string, []interface{}) {
	length := len(data)
	if length == 0 {
		return "", "", make([]interface{}, 0)
	}
	cols := make([]string, 0, length)
	placeholder := make([]string, 0, length)
	params := make([]interface{}, 0, length)
	for k, v := range data {
		cols = append(cols, k)
		placeholder = append(placeholder, "?")
		params = append(params, v)
	}
	return strings.Join(cols, ","), strings.Join(placeholder, ","), params
}

func MakeWhereParams(data map[string]interface{}) (string, []interface{}) {
	length := len(data)
	if length == 0 {
		return "", "", make([]interface{}, 0)
	}
	placeholder := make([]string, 0, length)
	params := make([]interface{}, 0, length)
	for k, v := range data {
		placeholder = append(placeholder, k+"= ?")
		params = append(params, v)
	}
	return strings.Join(cols, ","), params
}
