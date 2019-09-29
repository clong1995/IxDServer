package db

import (
	. "IxDServer/common"
	"fmt"
	"log"
)

func InsertProject(id, user, name string) error {
	_, err := Db.Exec("INSERT INTO project (id,user,name) values (?,?,?)", id, user, name)
	if err != nil {
		log.Println(err)
		return fmt.Errorf(SQL_STR)
	}
	return nil
}

func SelectProject(user string) ([]map[string]interface{}, error) {
	rows, err := Db.Query(`SELECT id,name FROM project where user = ? ORDER BY update_time DESC`, user)
	if err != nil {
		log.Println(err)
		return nil, fmt.Errorf(SQL_STR)
	}
	defer rows.Close()
	records, err := ParseRows(rows)
	if err != nil {
		return nil, err
	}
	return records, nil
}

//删除设备
func DeleteProject(project string) error {
	_, err := Db.Exec("DELETE FROM project WHERE id=?", project)
	if err != nil {
		log.Println(err)
		return fmt.Errorf(SQL_STR)
	}
	return nil
}
