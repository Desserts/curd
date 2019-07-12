package curd

import (
	"fmt"
)

func (d *DB) Select(tableName string, cols []string, ins interface{}, condition map[string]interface{}) error {
	query := fmt.Sprintf("SELECT %s FROM %s")
	params := make([]interface{}, 0)
	ph := ""
	if len(condition) > 0 {
		ph, params = MakeWhereParams(condition)
		query += " WHERE " + ph
	}
	_, err := d.sqlxDB.Select(ins, query, params...)
	if err != nil {
		return err
	}
}

func (d *DB) SelectPage(tableName string, page, pageSize int, cols []string, ins interface{}, condition map[string]interface{}) error {
	query := fmt.Sprintf("SELECT %s FROM %s")
	params := make([]interface{}, 0)
	ph := ""
	if len(condition) > 0 {
		ph, params = MakeWhereParams(condition)
		query += " WHERE " + ph
	}
	query += " LIMIT ?, ?"
	params = append(page-1, pageSize)
	_, err := d.sqlxDB.Select(ins, query, params...)
	if err != nil {
		return err
	}
}

func (d *DB) Insert(tableName string, data map[string]interface{}) (int, error) {
	cols, ph, params := MakeWhereParams(data)
	query := fmt.Sprintf("INSERT INTO %s (%s) VALUES(%s)", tableName, cols, ph)
	res, err := d.sqlxDB.Exec(query, params)
	if err != nil {
		return 0, err
	}
	return res.RowsAffected()
}

func (d *DB) Update(tableName string, data map[string]interface{}) (int, error) {
	ph, params := MakeWhereParams(data)
	query := fmt.Sprintf("UPDATE %s SET %s", tableName, ph)
	res, err := d.sqlxDB.Exec(query, params...)
	if err != nil {
		return 0, err
	}
	return res.RowsAffected()
}

func (d *DB) UpdateWhere(tableName string, data map[string]interface{}, conditon map[string]interface{}) (int, error) {
	ph, params := MakeWhereParams(data)
	wph, wparams := MakeWhereParams(condition)
	query := fmt.Sprintf("UPDATE %s SET %s WHERE %s", tableName, ph, wph)
	params = append(params, wparams)
	res, err := d.sqlxDB.Exec(query, params...)
	if err != nil {
		return 0, err
	}
	return res.RowsAffected()
}

func (d *DB) Delete(tableName string, data map[string]interface{}) (int, error) {
	ph, params := MakeWhereParams(data)
	query := fmt.Sprintf("DELETE FROM %s WHERE %s", tableName, ph)
	res, err := d.sqlxDB.Exec(query, params...)
	if err != nil {
		return 0, err
	}
	return res.RowsAffected()
}
