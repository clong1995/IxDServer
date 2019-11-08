package db

import (
	. "IxDServer/common"
	"fmt"
	"log"
)

//是否存在email
func HasCompany(name string) (bool, error) {
	rows, err := Db.Query("select id from company where name=? limit 1", name)
	if err != nil {
		log.Println(err)
		return true, fmt.Errorf(SQL_STR)
	}
	defer rows.Close()
	records, err := ParseRows(rows)
	if err != nil {
		return true, err
	}
	if len(records) > 0 {
		return true, nil
	}
	return false, nil
}

//插入用户
func InsertCompany(id, name string) error {
	_, err := Db.Exec(
		"INSERT company (id,name) values (?,?)",
		id, name)
	if err != nil {
		log.Println(err)
		return fmt.Errorf(SQL_STR)
	}
	return nil
}
