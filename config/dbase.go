package config

import (
	"fmt"

	"xorm.io/xorm"
)

func DBase() *xorm.Engine {
	db, err := xorm.NewEngine(ENGINE(), DBCONFIG())

	if err != nil {
		fmt.Println(err)
	}
	return db
}
