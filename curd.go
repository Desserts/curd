package curd

import (
	"fmt"
)

// R 查询语句
func (d *DB) R(tableName string, cols []string, condition map[string]interface{}, ins interface{}) error {
	query := fmt.Sprintf("SELECT %s FROM %s", MakeColsParams(cols), SpecialField(tableName))
	params := make([]interface{}, 0)
	ph := ""
	if len(condition) > 0 {
		ph, params = MakeWhereParams(condition)
		query += " WHERE " + ph
	}
	err := d.Select(ins, query, params...)
	return err
}

// RPage 分页的查询语句
func (d *DB) RPage(tableName string, page, pageSize int, cols []string, condition map[string]interface{}, ins interface{}) error {
	query := fmt.Sprintf("SELECT %s FROM %s", MakeColsParams(cols), SpecialField(tableName))
	params := make([]interface{}, 0)
	ph := ""
	if len(condition) > 0 {
		ph, params = MakeWhereParams(condition)
		query += " WHERE " + ph
	}
	query += " LIMIT ?, ?"
	params = append(params, page-1, pageSize)
	err := d.Select(ins, query, params...)
	return err
}

// C ...
func (d *DB) C(tableName string, data map[string]interface{}) (int64, error) {
	cols, ph, params := MakeInsertParams(data)
	query := fmt.Sprintf("INSERT INTO %s (%s) VALUES(%s)", SpecialField(tableName), cols, ph)
	res, err := d.Exec(query, params...)
	if err != nil {
		return 0, err
	}
	return res.RowsAffected()
}

// U 更新数据
func (d *DB) U(tableName string, data map[string]interface{}) (int64, error) {
	ph, params := MakeWhereParams(data)
	query := fmt.Sprintf("UPDATE %s SET %s", SpecialField(tableName), ph)
	res, err := d.Exec(query, params...)
	if err != nil {
		return 0, err
	}
	return res.RowsAffected()
}

// UWhere 带 where 的更新
func (d *DB) UWhere(tableName string, data map[string]interface{}, condition map[string]interface{}) (int64, error) {
	ph, params := MakeWhereParams(data)
	wph, wparams := MakeWhereParams(condition)
	query := fmt.Sprintf("UPDATE %s SET %s WHERE %s", SpecialField(tableName), ph, wph)
	params = append(params, wparams)
	res, err := d.Exec(query, params...)
	if err != nil {
		return 0, err
	}
	return res.RowsAffected()
}

// D 删除语句
func (d *DB) D(tableName string, data map[string]interface{}) (int64, error) {
	ph, params := MakeWhereParams(data)
	query := fmt.Sprintf("DELETE FROM %s WHERE %s", SpecialField(tableName), ph)
	res, err := d.Exec(query, params...)
	if err != nil {
		return 0, err
	}
	return res.RowsAffected()
}
