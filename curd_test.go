package curd

import (
	"testing"

	_ "github.com/go-sql-driver/mysql"
	"github.com/stretchr/testify/assert"
)

var (
	db = NewDB(
		"test",
		"test",
		"",
		"3306",
		"test",
	)
)

type Order struct {
	ID     int64 `db:"id"`
	Status int   `db:"status"`
}

func TestSelect(t *testing.T) {
	ins := make([]*Order, 0)
	err := db.R(
		"order",
		[]string{"id", "status"},
		map[string]interface{}{"id": 1},
		&ins,
	)
	assert.Nil(t, err)
	assert.Equal(t,
		[]*Order{
			&Order{ID: 1, Status: 3},
		}, ins)
}

func TestSelectPage(t *testing.T) {
	ins := make([]*Order, 0)
	err := db.RPage(
		"order",
		1, 5,
		[]string{"id", "status"},
		map[string]interface{}{"status": 2},
		&ins,
	)
	assert.Nil(t, err)
	assert.Equal(t, 5, len(ins))

	ins = make([]*Order, 0)
	err = db.RPage(
		"order",
		1, 10,
		[]string{"id", "status"},
		map[string]interface{}{"status": 2},
		&ins,
	)
	assert.Nil(t, err)
	assert.Equal(t, 10, len(ins))
}

func TestInsert(t *testing.T) {
	oldCount := 0
	err := db.Get(&oldCount, "select COUNT(*) FROM `user`")
	assert.Nil(t, err)

	_, err = db.C(
		"user",
		map[string]interface{}{
			"name": "314156",
		},
	)
	assert.Nil(t, err)

	newCount := 0
	_ = db.Get(&newCount, "select COUNT(*) FROM `user`")
	assert.Equal(t, newCount, oldCount+1)

	_, err = db.D(
		"user",
		map[string]interface{}{
			"name": "314156",
		},
	)
	assert.Nil(t, err)
	_ = db.Get(&newCount, "select COUNT(*) FROM `user`")
	assert.Equal(t, newCount, oldCount)
}
