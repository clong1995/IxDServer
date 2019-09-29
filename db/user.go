package db

import (
	. "IxDServer/common"
	"fmt"
	"log"
)

//插入用户
func InsertUser(id, email, password, belong string, rank int) error {
	_, err := Db.Exec(
		"INSERT user (id,email,password,belong,rank) values (?,?,?,?,?)",
		id, email, password, belong, rank)
	if err != nil {
		log.Println(err)
		return fmt.Errorf(SQL_STR)
	}
	return nil
}

//插入用户
func UserDelete(uid string) error {
	_, err := Db.Exec("DELETE FROM user WHERE id = ?", uid)
	if err != nil {
		log.Println(err)
		return fmt.Errorf(SQL_STR)
	}
	return nil
}

//是否存在email
func HasEmail(email string) (bool, error) {
	rows, err := Db.Query("select email from user where email=? limit 1", email)
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

//根据email查询用户信息
func SelectUserByEmail(email string) (map[string]interface{}, error) {
	rows, err := Db.Query(`
		select 
			id,
			email,
			password,
			status,
			create_time
		from user 
		where email=? limit 1`, email)

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
		return nil, fmt.Errorf(EMPTY_STR)
	}
	return records[0], nil
}

//根据id查询用户信息
func SelectUserById(id string) (map[string]interface{}, error) {
	rows, err := Db.Query(`
		select 
			id,
			email,
			status,
			rank,
			company,
			belong,
			create_time
		from user 
		where id= ? limit 1`, id)
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
		log.Println(EMPTY_STR)
		return nil, fmt.Errorf(EMPTY_STR)
	}
	return records[0], nil
}

//获取用户诶列表
func SelectUserList(company string, rank int) ([]map[string]interface{}, error) {
	rows, err := Db.Query(
		`SELECT id,email,status,company,create_time FROM user where belong = ? AND rank = ? ORDER BY update_time DESC`,
		company, rank)
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
