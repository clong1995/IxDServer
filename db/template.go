package db

import (
	. "IxDServer/common"
	"fmt"
	"log"
)

func SelectTemplateListByUser(user string) ([]map[string]interface{}, error) {
	rows, err := Db.Query(`SELECT id,name FROM template where user = ? ORDER BY create_time DESC`, user)
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

func SelectTemplateById(id string) (map[string]interface{}, error) {
	rows, err := Db.Query(`SELECT id,name,data FROM template where id = ? limit 1`, id)
	if err != nil {
		log.Println(err)
		return nil, fmt.Errorf(SQL_STR)
	}
	defer rows.Close()
	records, err := ParseRows(rows)
	if err != nil {
		return nil, err
	}
	if len(records) != 1 {
		return nil, fmt.Errorf(MULTIPLE_STR)
	}
	return records[0], nil
}

func InsertTemplate(id, user, name, data string) error {
	_, err := Db.Exec("INSERT INTO template (id,name,data,user) values (?,?,?,?)", id, name, data, user)
	if err != nil {
		log.Println(err)
		return fmt.Errorf(SQL_STR)
	}
	return nil
}
