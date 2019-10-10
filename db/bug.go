package db

import (
	. "IxDServer/common"
	"fmt"
	"log"
)

//插入用户
func InsertBug(uid, id, title, position, reappear, expect string, typee, severity, priority int) error {
	_, err := Db.Exec(
		`INSERT 
    				bug (user,id, title, position, reappear, expect, typee, severity, priority,state) 
    			values (?,?,?,?,?,?,?,?,?,1)`,
		uid, id, title, position, reappear, expect, typee, severity, priority)
	if err != nil {
		log.Println(err)
		return fmt.Errorf(SQL_STR)
	}
	return nil
}

//获取用户诶列表
func SelectBugList(company string, rank int) ([]map[string]interface{}, error) {
	rows, err := Db.Query(
		`SELECT 
       				u.email as email,
       				b.id as id, b.title as title, b.typee as type, b.severity as severity, b.priority as priority, 
       				b.state as state, b.create_time as create_time
				FROM bug b
				LEFT JOIN user u on b.user = u.id
				ORDER BY b.update_time DESC`)
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

//获取bug详细信息
func SelectBugById(id string) (map[string]interface{}, error) {
	rows, err := Db.Query(
		`SELECT user,id, title, position, reappear, expect, typee as type, severity, priority, state 
				FROM bug 
				where id = ? LIMIT 1`,
		id)
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
