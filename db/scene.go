package db

import (
	. "IxDServer/common"
	"fmt"
	"log"
)

func SelectSceneListByProject(project string) ([]map[string]interface{}, error) {
	rows, err := Db.Query(`SELECT id,name FROM scene where project = ? ORDER BY update_time DESC`, project)
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

func SelectSceneById(id string) (map[string]interface{}, error) {
	rows, err := Db.Query(`SELECT id,name,data FROM scene where id = ? limit 1`, id)
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

func InsertScene(id, project, name, data string) error {
	_, err := Db.Exec("INSERT INTO scene (id,name,data,project) values (?,?,?,?)", id, name, data, project)
	if err != nil {
		log.Println(err)
		return fmt.Errorf(SQL_STR)
	}
	return nil
}

func UpdateScene(id, name, data string) error {
	_, err := Db.Exec("UPDATE scene SET name=?,data=? WHERE id = ?", name, data, id)
	if err != nil {
		log.Println(err)
		return fmt.Errorf(SQL_STR)
	}
	return nil
}
