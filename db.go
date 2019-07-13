package curd

import (
	"fmt"
	"strings"

	"github.com/jmoiron/sqlx"
)

// DB 数据库实例
type DB struct {
	*sqlx.DB
}

const (
	dsn       = "%s:%s@tcp(%s:%s)/%s?%s"
	dsnParams = "charset=utf8mb4&parseTime=true&timeout=5s"
)

// NewDB 新建数据库实例
func NewDB(user, pass, host, port, dbname string) *DB {
	db, err := sqlx.Open("mysql", fmt.Sprintf(dsn, user, pass, host, port, dbname, dsnParams))
	if err != nil {
		panic(fmt.Sprintf("open db err: %v", err))
	}
	if err = db.Ping(); err != nil {
		panic(fmt.Sprintf("ping db err: %v", err))
	}
	return &DB{DB: db}
}

// MakeInsertParams 组装 sql 参数
func MakeInsertParams(data map[string]interface{}) (string, string, []interface{}) {
	length := len(data)
	if length == 0 {
		return "", "", make([]interface{}, 0)
	}
	cols := make([]string, 0, length)
	placeholder := make([]string, 0, length)
	params := make([]interface{}, 0, length)
	for k, v := range data {
		k := strings.Replace(k, "`", "", -1)
		cols = append(cols, fmt.Sprintf("`%s`", k))
		placeholder = append(placeholder, "?")
		params = append(params, v)
	}
	return strings.Join(cols, ","), strings.Join(placeholder, ","), params
}

// MakeWhereParams 构造 where 部分 sql 语句
func MakeWhereParams(data map[string]interface{}) (string, []interface{}) {
	length := len(data)
	if length == 0 {
		return "", make([]interface{}, 0)
	}
	placeholder := make([]string, 0, length)
	params := make([]interface{}, 0, length)
	for k, v := range data {
		k := strings.Replace(k, "`", "", -1)
		placeholder = append(placeholder, fmt.Sprintf("`%s`= ?", k))
		params = append(params, v)
	}
	return strings.Join(placeholder, ","), params
}

// MakeColsParams 拼接字段部分 sql 语句
func MakeColsParams(cols []string) string {
	c := make([]string, 0, len(cols))
	for _, v := range cols {
		v := strings.Replace(v, "`", "", -1)
		c = append(c, fmt.Sprintf("`%s`", v))
	}
	return strings.Join(c, ",")
}

// SpecialField 处理表名，列名的反引号
func SpecialField(s string) string {
	if strings.Contains(s, "`") {
		return s
	}
	return fmt.Sprintf("`%s`", s)
}
