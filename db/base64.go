package db

import (
	. "IxDServer/common"
	"fmt"
	"log"
)

func SelectBase64(id string) (map[string]interface{}, error) {
	rows, err := Db.Query(`SELECT value FROM base64 where id = ? limit 1`, id)
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

func InsertBase64(id, value string) error {
	_, err := Db.Exec("INSERT INTO base64 (id,value) values (?,?)", id, value)
	if err != nil {
		log.Println(err)
		return fmt.Errorf(SQL_STR)
	}
	return nil
}

func UpdateBase64(id, value string) error {
	_, err := Db.Exec("UPDATE base64 SET value=? WHERE id = ?", value, id)
	if err != nil {
		log.Println(err)
		return fmt.Errorf(SQL_STR)
	}
	return nil
}
